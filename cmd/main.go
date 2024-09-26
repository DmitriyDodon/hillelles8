package main

import (
	"les8/config"
	"les8/db"
	"les8/server"
	"les8/server/controller"

	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := config.InnitConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	dbConnection, err := db.NewConnection(config)

	if err != nil {
		log.Fatal(err.Error())
	}

	defer dbConnection.Close()

	_, err = dbConnection.RunQueryFromFile(config.GetMigrationPath())

	if err != nil {
		log.Fatalf("Migration failed to up: %s", err.Error())
	}

	controller := controller.NewController(dbConnection)

	server := server.NewServer(config, controller)

	server.Start()
}
