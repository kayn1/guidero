package inmemory

import (
	"context"

	"github.com/google/uuid"
	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/repository"
)

var _ repository.Repository = &InMemoryRepository{}

type InMemoryRepository struct {
}

func NewRepository() *InMemoryRepository {
	return &InMemoryRepository{}
}

// CreateUser implements repository.Repository.
func (i *InMemoryRepository) CreateUser(ctx context.Context, user domain.CreateUserRequest) (*domain.User, error) {
	// Create user in memory

	return &domain.User{
		ID:    uuid.Must(uuid.NewRandom()),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
