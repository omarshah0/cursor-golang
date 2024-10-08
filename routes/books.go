package routes

import (
	"project/handlers"

	"github.com/gin-gonic/gin"
)

// SetupBookRoutes sets up the book-related routes
func SetupBookRoutes(router *gin.Engine, handler *handlers.BookHandler) {
	bookGroup := router.Group("/books")
	{
		bookGroup.GET("/:id", handler.GetBook)
		bookGroup.POST("/", handler.CreateBook)
		bookGroup.PUT("/:id", handler.UpdateBook)    // Added route for updating a book
		bookGroup.DELETE("/:id", handler.DeleteBook) // Added route for deleting a book
		// Add other book routes as needed
	}
}
