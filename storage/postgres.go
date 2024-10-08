package storage

import (
	"errors"
	"fmt"
	"project/types"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// PostgresStorage implements the Storage interface using PostgreSQL via GORM
type PostgresStorage struct {
	db *gorm.DB
}

// NewPostgresStorage initializes PostgreSQL storage
func NewPostgresStorage(dsn string) (*PostgresStorage, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	// Migrate the schema
	if err := db.AutoMigrate(&types.User{}, &types.Book{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database schema: %w", err)
	}

	return &PostgresStorage{db: db}, nil
}

// GetUsers retrieves all users from PostgreSQL
func (p *PostgresStorage) GetUsers() ([]*types.User, error) {
	var users []*types.User
	result := p.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

// GetUser retrieves a user by ID from PostgreSQL
func (p *PostgresStorage) GetUser(id string) (*types.User, error) {
	var user types.User
	result := p.db.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

// CreateUser adds a new user to PostgreSQL
func (p *PostgresStorage) CreateUser(user *types.User) error {
	result := p.db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateUser updates an existing user in PostgreSQL
func (p *PostgresStorage) UpdateUser(user *types.User) error {
	result := p.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteUser removes a user by ID from PostgreSQL
func (p *PostgresStorage) DeleteUser(id string) error {
	result := p.db.Delete(&types.User{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("user not found")
	}
	return nil
}

// GetBook retrieves a book by ID from PostgreSQL
func (p *PostgresStorage) GetBook(id string) (*types.Book, error) {
	var book types.Book
	result := p.db.First(&book, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, result.Error
	}
	return &book, nil
}

// CreateBook adds a new book to PostgreSQL
func (p *PostgresStorage) CreateBook(book *types.Book) error {
	result := p.db.Create(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// UpdateBook updates an existing book in PostgreSQL
func (p *PostgresStorage) UpdateBook(book *types.Book) error {
	result := p.db.Save(book)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// DeleteBook removes a book by ID from PostgreSQL
func (p *PostgresStorage) DeleteBook(id string) error {
	result := p.db.Delete(&types.Book{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("book not found")
	}
	return nil
}

// ... implement other database operations as needed ...
