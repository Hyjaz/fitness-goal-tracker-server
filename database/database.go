package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
)

var client *mongo.Client
var db *mongo.Database

// New creates a new mongo client and returns a mongo.Database pointer
func New(host string, port string, database string) {
	if client == nil && db == nil {
		log.Println("starting database...")
		client = createMongoClient(host, port)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := client.Connect(ctx)
		log.Println(err)
		for err != nil {
			time.Sleep(1000 * time.Millisecond)
			err = client.Connect(ctx)
			log.Println(err)
		}

		if err != nil {
			log.Fatal(err)
		}

		db = client.Database(database)
		log.Println("database started")
	} else {
		log.Fatal("A mongodb client already exists.")
	}
}

// Disconnect closes the connection to the connected mongo db server
func Disconnect() error {
	return client.Disconnect(context.Background())
}

// GetDb returns a pointer to an database
func GetDb() *mongo.Database {
	return db
}

func createMongoClient(host string, port string) *mongo.Client {
	uri := fmt.Sprintf("mongodb://%s:%s", host, port)
	log.Println(uri)
	client, err := mongo.NewClient(uri)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
