package users

import (
	"context"

	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/repository"
)

var _ domain.UserService = &UserService{}

func NewUserService(repo repository.Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

type UserService struct {
	repo repository.Repository
}

func (u *UserService) CreateUser(ctx context.Context, user domain.CreateUserRequest) (*domain.User, error) {
	return u.repo.CreateUser(ctx, user)
}
