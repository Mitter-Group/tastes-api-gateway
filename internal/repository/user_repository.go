package repository

import (
	"context"

	"github.com/chunnior/api-gateway/internal/domain"
)

// UserRepository is an interface that represents the methods that a user repository should have.
type UserRepository interface {
	GetByID(ctx context.Context, id string) (*domain.User, error)
	Create(ctx context.Context, user *domain.User) error
	Update(ctx context.Context, id string, user *domain.User) error
	Delete(ctx context.Context, id string) error
}

// userRepository is a struct that represents a user repository.
type userRepository struct {
	// Here you can add the database connection or any other dependencies that your repository might need.
}

// NewUserRepository creates a new user repository.
func NewUserRepository() UserRepository {
	return &userRepository{}
}

// GetByID gets a user by id.
func (r *userRepository) GetByID(ctx context.Context, id string) (*domain.User, error) {
	// Here you should implement the logic to get a user by id from your database.
	return nil, nil
}

// Create creates a new user.
func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	// Here you should implement the logic to create a new user in your database.
	return nil
}

// Update updates a user.
func (r *userRepository) Update(ctx context.Context, id string, user *domain.User) error {
	// Here you should implement the logic to update a user in your database.
	return nil
}

// Delete deletes a user.
func (r *userRepository) Delete(ctx context.Context, id string) error {
	// Here you should implement the logic to delete a user from your database.
	return nil
}
