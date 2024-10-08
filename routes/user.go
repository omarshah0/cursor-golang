package routes

import (
	"project/handlers"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes sets up the user-related routes
func SetupUserRoutes(router *gin.Engine, handler *handlers.UserHandler) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", handler.GetUsers)
		userGroup.GET("/:id", handler.GetUser)
		userGroup.POST("/", handler.CreateUser)
		userGroup.PUT("/:id", handler.UpdateUser)    // Added route for updating a user
		userGroup.DELETE("/:id", handler.DeleteUser) // Added route for deleting a user
		// Add other user routes as needed
	}
}
