package models

import "time"

type DailyIntake struct {
	Date           time.Time        `json:"date" bson:"date"`
	MacroNutrients []MacroNutrients `json:"macroNutrients" bson:"macroNutrients"`
}

type MacroNutrients struct {
	MealNumber    int    `json:"mealNumber" bson:"mealNumber"`
	Proteins      string `json:"proteins" bson:"proteins"`
	Carbohydrates string `json:"carbohydrates" bson:"carbohydrates"`
	Fat           string `json:"fat" bson:"fat"`
}
