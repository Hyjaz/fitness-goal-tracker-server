package models

import (
	"context"

	"github.com/hyjaz/fitness-goal-tracker-server/database"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// User form tag is required so that we can bind the query string to this struct
type User struct {
	UUID   string  `form:"uuid" json:"uuid" bson:"uuid" binding:"required"`
	Cycles []Cycle `json:"cycles" bson:"cycles"`
}

func getUserCollection() *mongo.Collection {
	db := database.GetDb()
	return db.Collection("users")
}

// GetUser gets a user if it exists, otherwise inserts a new user and returns the document
func GetUser(user *User) error {
	collection := getUserCollection()
	filter := bson.M{"uuid": user.UUID}
	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		err = addUser(user.UUID)
		if err != nil {
			return err
		}
		collection.FindOne(context.Background(), filter).Decode(&user)
	}

	return nil
}

func addUser(uuid string) error {
	collection := getUserCollection()
	_, err := collection.InsertOne(context.Background(), User{UUID: uuid, Cycles: []Cycle{}})
	if err != nil {
		return err
	}
	return nil
}
