package main

import (
	"os"
	"os/signal"

	"github.com/hyjaz/fitness-goal-tracker-server/database"
	"github.com/hyjaz/fitness-goal-tracker-server/server"
)

func main() {

	database.Init(os.Getenv("DBHOSTNAME"), os.Getenv("DBPORTNUMBER"), os.Getenv("DBNAME"))
	server.Init()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	database.Disconnect()
	server.Shutdown()
}
