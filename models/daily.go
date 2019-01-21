package models

import (
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

type DailyIntakeTimeAsString struct {
	Date           string           `json:"date" bson:"date,string" binding:"required"`
	MacroNutrients []MacroNutrients `json:"macroNutrients" bson:"macroNutrients" binding:"required"`
}

//DailyIntake takes
type DailyIntake struct {
	Date           time.Time        `json:"date,string" bson:"date,string" binding:"required"`
	MacroNutrients []MacroNutrients `json:"macroNutrients" bson:"macroNutrients" binding:"required"`
}

type MacroNutrients struct {
	MealNumber    int    `json:"mealNumber" bson:"mealNumber" binding:"required"`
	Proteins      string `json:"proteins" bson:"proteins" binding:"required"`
	Carbohydrates string `json:"carbohydrates" bson:"carbohydrates" binding:"required"`
	Fat           string `json:"fat" bson:"fat" binding:"required"`
}

func AddDailyIntake(uuid string, date time.Time, macroNutrients []MacroNutrients) {
	collection := getUserCollection()
	d := DailyIntake{
		Date:           date,
		MacroNutrients: macroNutrients}

	_, err := collection.UpdateOne(nil,
		bson.D{bson.E{Key: "uuid", Value: "123123123"}},
		bson.M{"$push": bson.M{"cycles.0.dailyIntakes": d}})
	if err != nil {
		log.Println(err)
	}
}
