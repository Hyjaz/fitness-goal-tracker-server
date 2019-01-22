package models

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo/options"
)

type MacroNutrientsWithIDAsString struct {
	ID             string         `json:"_id" bson:"_id" binding:"required"`
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
type Something struct {
	S []interface{}
}

func AddMacroNutrient(macroNutrientsWithIDAsString MacroNutrientsWithIDAsString, user *User) error {
	collection := getUserCollection()
	ID, err := primitive.ObjectIDFromHex(macroNutrientsWithIDAsString.ID)
	macroNutrientsWithIDAsString.MacroNutrients.ID = primitive.NewObjectID()

	filter := bson.M{"uuid": user.UUID, "cycles.dailyIntakes": bson.M{"$elemMatch": bson.M{"_id": ID}}}
	update := bson.M{"$push": bson.M{"cycles.$.dailyIntakes.$[i].macroNutrients": macroNutrientsWithIDAsString.MacroNutrients}}

	//options
	//Can I convert a []T to an []interface{}?
	//Not directly. It is disallowed by the language specification because the
	//two types do not have the same representation in memory. It is necessary to copy
	//the elements individually to the destination slice. This example converts a slice of int to a slice of interface{}:
	optionsArrayFilters := options.ArrayFilters{}
	dailyIntakeFilter := bson.M{"i._id": ID}
	optionsArrayFilters.Filters = append(optionsArrayFilters.Filters, dailyIntakeFilter)
	options := options.UpdateOptions{ArrayFilters: &optionsArrayFilters}
	_, err = collection.UpdateOne(nil, filter, update, &options)

	if err != nil {
		return err
	}
	collection.FindOne(nil, bson.M{"uuid": user.UUID}).Decode(&user)
	return nil
}
