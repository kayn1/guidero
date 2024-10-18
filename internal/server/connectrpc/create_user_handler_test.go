package connectrpc_test

import (
	"context"
	"testing"

	"connectrpc.com/connect"
	"github.com/google/uuid"
	v1 "github.com/kayn1/guidero/gen/proto/user/v1"
	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/domain/mocks"
	"github.com/kayn1/guidero/internal/server/connectrpc"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateUserHandler(t *testing.T) {
	mockUserService := mocks.NewMockUserService(gomock.NewController(t))

	app := domain.Application{
		UserService: mockUserService,
	}
	server := connectrpc.NewConnectRpcServer(app)

	req := &connect.Request[v1.CreateRequest]{
		Msg: &v1.CreateRequest{
			Email: "john@example.com",
			Name:  "John Doe",
		},
	}

	testUser := &domain.User{
		ID:    uuid.New(),
		Email: "john@example.com",
		Name:  "John Doe",
	}

	mockUserService.EXPECT().CreateUser(gomock.Any(), domain.CreateUserRequest{
		Email: "john@example.com",
		Name:  "John Doe",
	}).Return(testUser, nil)

	res, err := server.Create(context.Background(), req)

	assert.NoError(t, err)

	expectedRes := &connect.Response[v1.CreateResponse]{
		Msg: &v1.CreateResponse{
			User: &v1.User{
				Id:    testUser.ID.String(),
				Email: testUser.Email,
				Name:  testUser.Name,
			},
		},
	}
	assert.Equal(t, expectedRes, res)
}
