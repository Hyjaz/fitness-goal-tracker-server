package models

import (
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

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

type MacroNutrients struct {
	ID            primitive.ObjectID `json:"_id" bson:"_id"`
	MealNumber    int                `json:"mealNumber" bson:"mealNumber" binding:"required"`
	Proteins      string             `json:"proteins" bson:"proteins" binding:"required"`
	Carbohydrates string             `json:"carbohydrates" bson:"carbohydrates" binding:"required"`
	Fat           string             `json:"fat" bson:"fat" binding:"required"`
}

func AddDailyIntake(uuid string, id primitive.ObjectID, date time.Time, macroNutrients []MacroNutrients) User {
	collection := getUserCollection()

	for index := range macroNutrients {
		macroNutrients[index].ID = primitive.NewObjectID()
	}
	d := DailyIntake{
		ID:             primitive.NewObjectID(),
		Date:           date,
		MacroNutrients: macroNutrients}

	_, err := collection.UpdateOne(nil,
		bson.D{bson.E{Key: "uuid", Value: "123123123"}, bson.E{Key: "cycles", Value: bson.M{"$elemMatch": bson.M{"_id": id}}}},
		bson.M{"$push": bson.M{"cycles.0.dailyIntakes": d}})
	if err != nil {
		log.Fatal(err)
	}
	result := collection.FindOne(nil, bson.D{bson.E{Key: "uuid", Value: "123123123"}})

	var user User
	result.Decode(&user)
	return user
}
