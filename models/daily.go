package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

//DailyIntakeTimeAsString is simply used so that a unix timestamp can be binded to the struct
type DailyIntakeTimeAsString struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Date           string             `json:"date" bson:"date,string" binding:"required"`
	MacroNutrients []MacroNutrients   `json:"macroNutrients" bson:"macroNutrients" binding:"required"`
}

//DailyIntake takes
type DailyIntake struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	Date           time.Time          `json:"date,string" bson:"date,string" binding:"required"`
	MacroNutrients []MacroNutrients   `json:"macroNutrients" bson:"macroNutrients" binding:"required"`
}

// MacroNutrients macros for a meal
type MacroNutrients struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	MealNumber    int                `json:"mealNumber" bson:"mealNumber" binding:"required"`
	Proteins      string             `json:"proteins" bson:"proteins" binding:"required"`
	Carbohydrates string             `json:"carbohydrates" bson:"carbohydrates" binding:"required"`
	Fat           string             `json:"fat" bson:"fat" binding:"required"`
	Status        bool               `json:"status" bson:"status"`
}

// AddDailyIntake add a new DailyIntake in Cycle
func AddDailyIntake(id primitive.ObjectID, date time.Time, macroNutrients []MacroNutrients, user *User) error {
	collection := getUserCollection()

	for index := range macroNutrients {
		macroNutrients[index].ID = primitive.NewObjectID()
	}
	d := DailyIntake{
		ID:             primitive.NewObjectID(),
		Date:           date,
		MacroNutrients: macroNutrients}

	_, err := collection.UpdateOne(nil,
		bson.D{bson.E{Key: "uuid", Value: user.UUID}, bson.E{Key: "cycles", Value: bson.M{"$elemMatch": bson.M{"_id": id}}}},
		bson.M{"$push": bson.M{"cycles.$.dailyIntakes": d}})
	if err != nil {
		return err
	}
	collection.FindOne(nil, bson.D{bson.E{Key: "uuid", Value: user.UUID}}).Decode(&user)
	return nil
}
