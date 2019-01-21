package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type CycleWithTimeAsString struct {
	StartTime string `json:"startTime" bson:"startTime" binding:"required"`
	EndTime   string `json:"endTime" bson:"endTime" binding:"required"`
}

// Cycle contains a list your daily nutrient intakes
type Cycle struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id" binding:"required"`
	StartTime    time.Time          `json:"startTime" bson:"startTime"`
	EndTime      time.Time          `json:"endTime" bson:"endTime"`
	DailyIntakes []DailyIntake      `json:"dailyIntakes" bson:"dailyIntakes"`
}

// AddCycle adds a new embedded document in user document
func AddCycle(uuid string, startTime time.Time, endTime time.Time, user *User) error {
	collection := getUserCollection()

	c := Cycle{
		ID:           primitive.NewObjectID(),
		StartTime:    startTime,
		EndTime:      endTime,
		DailyIntakes: []DailyIntake{}}

	_, err := collection.UpdateOne(nil,
		bson.D{bson.E{Key: "uuid", Value: uuid}},
		bson.M{"$push": bson.D{bson.E{Key: "cycles", Value: c}}})
	if err != nil {
		return err
	}
	result := collection.FindOne(nil, bson.D{bson.E{Key: "uuid", Value: uuid}})

	result.Decode(&user)
	return nil
}