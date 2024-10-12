package inmemory

import (
	"context"

	"github.com/brianvoe/gofakeit/v7"
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
func (i *InMemoryRepository) CreateUser(ctx context.Context, userRequest domain.CreateUserRequest) (*domain.User, error) {

	// Generate random user using fakeit
	user := domain.User{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		ID:    uuid.Must(uuid.NewRandom()),
	}

	return &domain.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
