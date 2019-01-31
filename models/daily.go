package models

import (
	"time"

	parse "github.com/hyjaz/fitness-goal-tracker-server/utility"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//DailyIntakeDateAsString is simply used so that a unix timestamp can be binded to the struct
type DailyIntakeDateAsString struct {
	ID             string           `json:"_id" bson:"_id" binding:"required"`
	Date           string           `json:"date" bson:"date,string" binding:"required"`
	MacroNutrients []MacroNutrients `json:"macroNutrients" bson:"macroNutrients" binding:"required"`
}

//DailyIntake takes
type DailyIntake struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Date           time.Time          `json:"date,string" bson:"date,string" binding:"required"`
	MacroNutrients []MacroNutrients   `json:"macroNutrients" bson:"macroNutrients" binding:"required"`
}

// AddDailyIntake add a new DailyIntake in Cycle
func AddDailyIntake(dailyIntakeDateAsString DailyIntakeDateAsString, user *User) error {
	collection := getUserCollection()

	cycleObjectID, err := primitive.ObjectIDFromHex(dailyIntakeDateAsString.ID)
	for index := range dailyIntakeDateAsString.MacroNutrients {
		dailyIntakeDateAsString.MacroNutrients[index].ID = primitive.NewObjectID()
	}
	d := DailyIntake{
		ID:             primitive.NewObjectID(),
		Date:           parse.ConvertUnixTimestampToTime(dailyIntakeDateAsString.Date),
		MacroNutrients: dailyIntakeDateAsString.MacroNutrients}

	filter := bson.M{"uuid": user.UUID, "cycles": bson.M{"$elemMatch": bson.M{"_id": cycleObjectID}}}
	update := bson.M{"$push": bson.M{"cycles.$.dailyIntakes": d}}

	_, err = collection.UpdateOne(nil, filter, update)

	if err != nil {
		return err
	}
	collection.FindOne(nil, bson.M{"uuid": user.UUID}).Decode(&user)
	return nil
}
