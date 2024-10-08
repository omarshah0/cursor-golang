package api

import (
	"project/handlers"
	"project/routes"
	"project/storage"

	"github.com/gin-gonic/gin"
)

// StartServer initializes the Gin router and starts the server
func StartServer(store storage.Storage) {
	router := gin.Default()

	// Initialize handlers with the storage
	userHandler := handlers.NewUserHandler(store)
	bookHandler := handlers.NewBookHandler(store)

	// Set up routes
	routes.SetupUserRoutes(router, userHandler)
	routes.SetupBookRoutes(router, bookHandler)

	// Start the server on port 8080
	router.Run(":8080")
}
