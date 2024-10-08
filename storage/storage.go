package storage

import "project/types"

// Storage defines the interface for database operations
type Storage interface {
	GetUsers() ([]*types.User, error)
	GetUser(id string) (*types.User, error)
	CreateUser(user *types.User) error
	UpdateUser(user *types.User) error
	DeleteUser(id string) error

	GetBook(id string) (*types.Book, error)
	CreateBook(book *types.Book) error
	UpdateBook(book *types.Book) error
	DeleteBook(id string) error
	// Add other database operations as needed
}
