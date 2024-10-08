package storage

import (
	"errors"
	"project/types"
	"sync"
)

// MemoryStorage implements the Storage interface with in-memory storage
type MemoryStorage struct {
	users map[string]*types.User
	books map[string]*types.Book
	mu    sync.RWMutex
}

// NewMemoryStorage initializes the in-memory storage with dummy data and returns a Storage interface
func NewMemoryStorage(dsn string) (Storage, error) {
	m := &MemoryStorage{
		users: make(map[string]*types.User),
		books: make(map[string]*types.Book),
	}

	// Add dummy users
	m.users["1"] = &types.User{
		ID:    "1",
		Name:  "Alice Johnson",
		Email: "alice.johnson@example.com",
	}
	m.users["2"] = &types.User{
		ID:    "2",
		Name:  "Bob Smith",
		Email: "bob.smith@example.com",
	}

	// Add dummy books
	m.books["1"] = &types.Book{
		ID:     "1",
		Title:  "Go Programming Essentials",
		Author: "John Doe",
	}
	m.books["2"] = &types.Book{
		ID:     "2",
		Title:  "Mastering Gin Framework",
		Author: "Jane Smith",
	}

	return m, nil
}

// GetUsers retrieves all users
func (m *MemoryStorage) GetUsers() ([]*types.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	users := make([]*types.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}
	return users, nil
}

// GetUser retrieves a user by ID
func (m *MemoryStorage) GetUser(id string) (*types.User, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

// CreateUser adds a new user
func (m *MemoryStorage) CreateUser(user *types.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.users[user.ID]; exists {
		return errors.New("user already exists")
	}
	m.users[user.ID] = user
	return nil
}

// UpdateUser updates an existing user
func (m *MemoryStorage) UpdateUser(user *types.User) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.users[user.ID]; !exists {
		return errors.New("user not found")
	}
	m.users[user.ID] = user
	return nil
}

// DeleteUser removes a user by ID
func (m *MemoryStorage) DeleteUser(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.users[id]; !exists {
		return errors.New("user not found")
	}
	delete(m.users, id)
	return nil
}

// GetBook retrieves a book by ID
func (m *MemoryStorage) GetBook(id string) (*types.Book, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	book, exists := m.books[id]
	if !exists {
		return nil, errors.New("book not found")
	}
	return book, nil
}

// CreateBook adds a new book
func (m *MemoryStorage) CreateBook(book *types.Book) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.books[book.ID]; exists {
		return errors.New("book already exists")
	}
	m.books[book.ID] = book
	return nil
}

// UpdateBook updates an existing book
func (m *MemoryStorage) UpdateBook(book *types.Book) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.books[book.ID]; !exists {
		return errors.New("book not found")
	}
	m.books[book.ID] = book
	return nil
}

// DeleteBook removes a book by ID
func (m *MemoryStorage) DeleteBook(id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.books[id]; !exists {
		return errors.New("book not found")
	}
	delete(m.books, id)
	return nil
}

// ... other database operation implementations ...
