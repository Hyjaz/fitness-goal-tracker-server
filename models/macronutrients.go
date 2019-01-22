package models

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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
	update := bson.M{"$push": bson.M{"cycles.0.dailyIntakes.0.macroNutrients": macroNutrientsWithIDAsString.MacroNutrients}}
	_, err = collection.UpdateOne(nil, filter, update)

	if err != nil {
		return err
	}
	collection.FindOne(nil, bson.M{"uuid": user.UUID}).Decode(&user)
	return nil
}
