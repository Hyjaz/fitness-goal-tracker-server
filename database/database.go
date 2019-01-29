package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
)

var client *mongo.Client
var db *mongo.Database

// Init creates a new mongo client and returns a mongo.Database pointer
func Init(host string, port string, database string) {
	if client == nil && db == nil {
		log.Println("starting database...")
		client = createMongoClient(host, port)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := client.Ping(ctx, readpref.Primary())
		log.Println(err)
		for err != nil {
			time.Sleep(1000 * time.Millisecond)
			err = client.Connect(ctx)
			if err != nil {
				log.Fatal(err)
			} else {
				err = client.Ping(ctx, readpref.Primary())
				log.Println(err)
			}

		}
		client.Connect(ctx)
		if err != nil {

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
