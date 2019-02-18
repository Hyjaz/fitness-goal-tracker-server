package models

import (
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// CycleWithTimeAsString is simply used so that a unix timestamp can be binded to the struct
type CycleWithTimeAsString struct {
	Name      string `json:"name" bson:"name" binding:"required"`
	StartTime string `json:"startTime" bson:"startTime" binding:"required"`
	EndTime   string `json:"endTime" bson:"endTime" binding:"required"`
}

// Cycle contains a list your daily nutrient intakes
type Cycle struct {
	ID           primitive.ObjectID `json:"_id" bson:"_id"`
	Name         string             `json:"name" bson:"name"`
	StartTime    time.Time          `json:"startTime" bson:"startTime"`
	EndTime      time.Time          `json:"endTime" bson:"endTime"`
	DailyIntakes []DailyIntake      `json:"dailyIntakes" bson:"dailyIntakes"`
}

// AddCycle adds a new embedded document in user document
func AddCycle(name string, startTime time.Time, endTime time.Time, user *User) error {
	collection := getUserCollection()
	c := Cycle{
		ID:           primitive.NewObjectID(),
		Name:         name,
		StartTime:    startTime,
		EndTime:      endTime,
		DailyIntakes: []DailyIntake{}}
	filter := bson.M{"uuid": user.UUID}
	update := bson.M{"$push": bson.M{"cycles": c}}
	_, err := collection.UpdateOne(nil, filter, update)
	if err != nil {
		return err
	}
	result := collection.FindOne(nil, filter)

	result.Decode(&user)
	return nil
}
