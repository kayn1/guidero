//go:generate mockgen -destination=mocks/user_repository_mock.go -package=mocks github.com/kayn1/guidero/internal/repository UserRepository

package repository

import (
	"context"

	"github.com/kayn1/guidero/internal/domain"
)

type Repository interface {
	UserRepository
}

type UserRepository interface {
	CreateUser(ctx context.Context, user domain.CreateUserRequest) (*domain.User, error)
}
