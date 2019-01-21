package models

import (
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
)

type CycleWithTimeAsString struct {
	StartTime string `json:"startTime" bson:"startTime" binding:"required"`
	EndTime   string `json:"endTime" bson:"endTime" binding:"required"`
}

// Cycle contains a list your daily nutrient intakes
type Cycle struct {
	StartTime    time.Time     `json:"startTime" bson:"startTime"`
	EndTime      time.Time     `json:"endTime" bson:"endTime"`
	DailyIntakes []DailyIntake `json:"dailyIntakes" bson:"dailyIntakes"`
}

// AddCycle adds a new embedded document in user document
func AddCycle(uuid string, startTime time.Time, endTime time.Time) {
	collection := getUserCollection()

	c := Cycle{
		StartTime:    startTime,
		EndTime:      endTime,
		DailyIntakes: []DailyIntake{}}

	_, err := collection.UpdateOne(nil,
		bson.D{bson.E{Key: "uuid", Value: "123123123"}},
		bson.M{"$push": bson.D{bson.E{Key: "cycles", Value: c}}})
	if err != nil {
		log.Println(err)
	}
}
