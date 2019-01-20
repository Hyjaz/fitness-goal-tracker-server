package models

import (
	"context"
	"errors"
	"log"

	"github.com/hyjaz/fitness-goal-tracker-server/database"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type User struct {
	UUID   string  `json:"uuid" bson:"uuid"`
	Cycles []Cycle `json:"cycles" bson:"cycles"`
}

func getUserCollection() *mongo.Collection {
	db := database.GetDb()
	return db.Collection("users")
}

// GetUser gets a user if it exists, otherwise inserts a new user and returns the document
func GetUser(uuid string) (User, error) {
	var user User
	collection := getUserCollection()

	err := collection.FindOne(context.Background(), bson.M{"uuid": uuid}).Decode(&user)

	if err != nil {
		err = addUser(uuid)
		if err != nil {
			return User{}, err
		}
		collection.FindOne(context.Background(), bson.M{"uuid": uuid}).Decode(&user)
	}

	return user, nil
}

func addUser(uuid string) error {
	collection := getUserCollection()
	insertOneResult, err := collection.InsertOne(context.Background(), User{UUID: uuid, Cycles: []Cycle{}})
	if err != nil {
		return errors.New("Could not add user")
	}
	log.Printf("Inserted a single document: %+v\n", insertOneResult)
	return nil
}
