package handlers

import (
	"net/http"
	"project/storage"
	"project/types"

	"github.com/gin-gonic/gin"
)

// BookHandler handles book-related requests
type BookHandler struct {
	store storage.Storage
}

// NewBookHandler creates a new BookHandler
func NewBookHandler(store storage.Storage) *BookHandler {
	return &BookHandler{store: store}
}

// GetBook handles GET /books/:id
func (h *BookHandler) GetBook(c *gin.Context) {
	id := c.Param("id")
	book, err := h.store.GetBook(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// CreateBook handles POST /books
func (h *BookHandler) CreateBook(c *gin.Context) {
	var book types.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.store.CreateBook(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

// UpdateBook handles PUT /books/:id
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var book types.Book
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	book.ID = id // Ensure the ID matches the URL parameter
	if err := h.store.UpdateBook(&book); err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, book)
}

// DeleteBook handles DELETE /books/:id
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := h.store.DeleteBook(id); err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.Status(http.StatusNoContent)
}
