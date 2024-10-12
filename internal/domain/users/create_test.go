package users_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/kayn1/guidero/internal/domain"
// 	"github.com/kayn1/guidero/internal/domain/users"
// )

// func TestCreateUser(t *testing.T) {
// 	repo := &repository.MockRepository{}
// 	userService := users.NewUserService(repo)

// 	ctx := context.Background()
// 	user := domain.CreateUserRequest{
// 		Name:  "john_doe",
// 		Email: "john@example.com",
// 	}

// 	// Call the CreateUser method
// 	createdUser, err := userService.CreateUser(ctx, user)

// 	// Check if there was an error
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 	}

// 	// Check if the createdUser is not nil
// 	if createdUser == nil {
// 		t.Error("Expected a non-nil createdUser, but got nil")
// 	}

// 	// Add more assertions as needed
// }
