package users_test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/domain/users"
	"github.com/kayn1/guidero/internal/repository/mocks"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
)

type UserServiceTestSuite struct {
	ctl *gomock.Controller
	suite.Suite
	repo        *mocks.MockRepository
	userService *users.UserService
}

func (suite *UserServiceTestSuite) SetupTest() {
	ctl := gomock.NewController(suite.T())
	repo := mocks.NewMockRepository(ctl)
	userService := users.NewUserService(repo)

	suite.repo = repo
	suite.userService = userService
}

func (suite *UserServiceTestSuite) TestCreateUser() {
	ctx := context.Background()
	userRequest := domain.CreateUserRequest{
		Name:  "john_doe",
		Email: "john@example.com",
	}

	user := &domain.User{
		ID:    uuid.Must(uuid.NewRandom()),
		Name:  userRequest.Name,
		Email: userRequest.Email,
	}

	suite.repo.EXPECT().CreateUser(ctx, userRequest).Return(user, nil)
	createdUser, err := suite.userService.CreateUser(ctx, userRequest)

	suite.NoError(err)
	suite.Equal(user, createdUser)
	suite.Equal(user.Name, userRequest.Name)
	suite.Equal(user.Email, userRequest.Email)
}

func TestUserServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
