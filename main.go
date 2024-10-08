package main

import (
	"log"
	"os"
	"project/api"
	"project/storage"
)

func main() {
	// Retrieve PostgreSQL DSN from environment variable
	dsn := os.Getenv("POSTGRES_DSN")
	if dsn == "" {
		log.Default().Println("POSTGRES_DSN environment variable is required for PostgreSQL storage")
	}

	// Initialize PostgreSQL storage
	store, err := storage.NewMemoryStorage(dsn)
	if err != nil {
		log.Fatalf("Failed to initialize PostgreSQL storage: %v", err)
	}

	log.Println("Using PostgreSQL Storage")

	// Initialize and start the API server
	api.StartServer(store)
}
