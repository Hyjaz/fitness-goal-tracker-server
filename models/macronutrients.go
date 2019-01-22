package models

import (
	"log"

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
	dailyIntakeObjectID, err := primitive.ObjectIDFromHex(macroNutrientsWithIDAsString.ID)
	macroNutrientsWithIDAsString.MacroNutrients.ID = primitive.NewObjectID()
	result := collection.FindOne(nil, bson.D{bson.E{Key: "uuid", Value: user.UUID}, bson.E{Key: "dailyIntakes", Value: bson.M{"$elemMatch": bson.M{"_id": dailyIntakeObjectID}}}})
	result.Decode(&user)
	log.Println(user)
	_, err = collection.UpdateOne(nil, bson.D{bson.E{Key: "uuid", Value: user.UUID}, bson.E{Key: "dailyIntakes", Value: bson.M{"$elemMatch": bson.M{"_id": dailyIntakeObjectID}}}},
		bson.M{"$push": bson.M{"dailyIntakes.$.macroNutrients": macroNutrientsWithIDAsString.MacroNutrients}})

	if err != nil {
		return err
	}
	collection.FindOne(nil, bson.D{bson.E{Key: "uuid", Value: user.UUID}}).Decode(&user)
	return nil
}
