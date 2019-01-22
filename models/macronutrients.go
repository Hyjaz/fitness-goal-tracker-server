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

func AddMacroNutrient(macroNutrientsWithIDAsString MacroNutrientsWithIDAsString, user *User) error {
	collection := getUserCollection()
	ID, err := primitive.ObjectIDFromHex(macroNutrientsWithIDAsString.ID)
	macroNutrientsWithIDAsString.MacroNutrients.ID = primitive.NewObjectID()
	filter := bson.M{"uuid": user.UUID, "cycles.dailyIntakes": bson.M{"$elemMatch": bson.M{"_id": ID}}}
	// Since only one document will be returned, we can safely index
	update := bson.M{"$push": bson.M{"cycles.$.dailyIntakes.$[].macroNutrients": macroNutrientsWithIDAsString.MacroNutrients}}
	optionsArrayFilters := options.ArrayFilters{Filters: [bson.M{"i._id": ID}]}
	// options := options.UpdateOptions{ArrayFilters: }
	arrayFilter := bson.M{"$arrayFilter": bson.M{"i._id": ID}}
	_, err = collection.UpdateOne(nil, filter, update, arrayFilter)

	if err != nil {
		return err
	}
	collection.FindOne(nil, bson.M{"uuid": user.UUID}).Decode(&user)
	return nil
}
