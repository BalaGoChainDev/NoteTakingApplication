package main

import (
	// Standard libs
	"log"
	"os"

	// ThirdParty libs
	"github.com/joho/godotenv"

	// Custom Libs
	"github.com/BalaGoChainDev/NoteTakingApplication/repositories"
	"github.com/BalaGoChainDev/NoteTakingApplication/server"
)

func init() {
	// Loads environment variables from the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %s", err)
	}

	// Initializes the database connection
	err = repositories.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB: %s", err)
	}
}

func main() {
	server := server.NewServer()
	listenAddr := os.Getenv("LISTEN_ADDR")

	log.Printf("Server listening on %v", listenAddr)
	err := server.Run(listenAddr)
	if err != nil {
		log.Fatalf("Error while serving: %s", err)
	}
}
