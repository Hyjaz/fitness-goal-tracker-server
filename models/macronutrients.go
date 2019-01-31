package models

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

// MacroNutrientsWithCycleAndDailyIntakeID struct for incoming message
type MacroNutrientsWithCycleAndDailyIntakeID struct {
	CycleID        string         `json:"_cycleId" binding:"required"`
	DailyID        string         `json:"_dailyId" bson:"_id" binding:"required"`
	MacroNutrients MacroNutrients `json:"macroNutrients" bson:"macroNutrients" binding:"required"`
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

// AddMacroNutrient adds a macronutrient in the correct cycle and dailyintake
func AddMacroNutrient(macroNutrientsWithCycleAndDailyIntakeID MacroNutrientsWithCycleAndDailyIntakeID, user *User) error {

	collection := getUserCollection()
	cycleID, err := primitive.ObjectIDFromHex(macroNutrientsWithCycleAndDailyIntakeID.CycleID)
	dailyID, err := primitive.ObjectIDFromHex(macroNutrientsWithCycleAndDailyIntakeID.DailyID)
	macroNutrientsWithCycleAndDailyIntakeID.MacroNutrients.ID = primitive.NewObjectID()

	filter := bson.M{"uuid": user.UUID}
	update := bson.M{"$push": bson.M{"cycles.$[i].dailyIntakes.$[j].macroNutrients": macroNutrientsWithCycleAndDailyIntakeID.MacroNutrients}}
	optionsArrayFilters := options.ArrayFilters{}
	cycleFilter := bson.M{"i._id": cycleID}
	dailyIntakeFilter := bson.M{"j._id": dailyID}
	optionsArrayFilters.Filters = append(optionsArrayFilters.Filters, cycleFilter)
	optionsArrayFilters.Filters = append(optionsArrayFilters.Filters, dailyIntakeFilter)
	options := options.UpdateOptions{ArrayFilters: &optionsArrayFilters}

	_, err = collection.UpdateOne(nil, filter, update, &options)
	if err != nil {
		return err
	}

	collection.FindOne(nil, bson.M{"uuid": user.UUID}).Decode(&user)
	return nil
}
