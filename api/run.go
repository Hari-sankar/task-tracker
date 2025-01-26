package api

import (
	"log"
	"os"
	"task-tracker/api/server"
	"task-tracker/pkg/database/postgres"
	"task-tracker/pkg/logger"
	"task-tracker/pkg/swagger"

	"github.com/joho/godotenv"
)

// Run initializes the required components and starts the server
func Run() {

	logger.InitLogger()
	logger.Info("Starting server")

	//Loading env variables
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println(" Error : Error loading environment variables")

	}

	//Initialising db
	connectionURL := os.Getenv("DB_SOURCE")

	psqlDB, err := postgres.NewPsqlDB(connectionURL)
	if err != nil {
		log.Fatalf("Postgresql init: %s", err)
	} else {
		log.Default().Printf("Postgres connected, Status: %#v", psqlDB.Stat())
	}
	defer psqlDB.Close()

	// Generate fresh docs on startup
	swagger.GenerateSwaggerDocs()

	//initialise server
	server := server.NewServer(psqlDB)

	server.Run()

}
