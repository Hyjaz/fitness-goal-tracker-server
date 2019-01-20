package models

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/hyjaz/fitness-goal-tracker-server/database"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type User struct {
	UUID   string  `json:"uuid" bson:"uuid"`
	Cycles []Cycle `json:"cycles" bson:"cycles"`
}

type Cycle struct {
	StartTime   time.Time
	EndTime     time.Time
	CycleIntake []DailyIntakes
}

type DailyIntakes struct {
	Date           time.Time
	MacroNutrients []MacroNutrients
}

type MacroNutrients struct {
	MealNumber    int
	Proteins      string
	Carbohydrates string
	Fat           string
}

func getCollection() *mongo.Collection {
	db := database.GetDb()
	return db.Collection("users")
}

// GetUser gets a user if it exists, otherwise inserts a new user and returns the document
func GetUser(uuid string) (User, error) {
	var user User
	collection := getCollection()

	filter := User{UUID: uuid}
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		err = addUser(uuid)
		if err != nil {
			return User{}, err
		} else {
			collection.FindOne(context.Background(), filter).Decode(&user)
		}

	}

	return user, nil
}

func addUser(uuid string) error {
	collection := getCollection()
	insertOneResult, err := collection.InsertOne(context.Background(), User{UUID: uuid})
	if err != nil {
		return errors.New("Could not add user")
	} else {
		log.Printf("Inserted a single document: %+v\n", insertOneResult)
		return nil
	}
}
