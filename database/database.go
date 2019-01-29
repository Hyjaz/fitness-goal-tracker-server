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

// Init creates a new mongo client and returns a mongo.Database pointer
func Init(host string, port string, database string) {
	if client == nil && db == nil {
		client = createMongoClient(host, port)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := client.Connect(ctx)
		for err != nil {
			log.Fatal(err)
			time.Sleep(1000 * time.Millisecond)
			err = client.Connect(ctx)
		}

		db = client.Database(database)
		log.Println("database started")
	} else {
		log.Fatal("A mongodb client already exists.")
	}
}

// Disconnect closes the connection to the connected mongo db server
func Disconnect() {
	log.Println("Disconnecting database...")
	if err := client.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Println("database disconnected.")
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
		log.Fatal(err)
	}
	return client
}
