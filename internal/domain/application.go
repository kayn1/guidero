package domain

import (
	"context"

	"github.com/google/uuid"
)

type Application struct {
	UserService UserService
}

type UserService interface {
	CreateUser(ctx context.Context, user CreateUserRequest) (*User, error)
}

type User struct {
	ID    uuid.UUID
	Name  string
	Email string
}

type CreateUserRequest struct {
	Email string
	Name  string
}

func NewApplication(userService UserService) Application {
	return Application{
		UserService: userService,
	}
}
