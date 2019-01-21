package main

import (
	"log"
	"os"

	"github.com/hyjaz/fitness-goal-tracker-server/database"
	"github.com/hyjaz/fitness-goal-tracker-server/server"
)

func main() {
	// TODO: if docker is not used to run this web server, the param needs to be changed to localhost
	database.New(os.Getenv("DBHOSTNAME"), "27017", "fitness-goal-tracker")

	server.Init()

	log.Println("Disconnecting database...")
	if err := database.Disconnect(); err != nil {
		log.Fatal(err)
	}
	log.Println("database disconnected.")
}
