package connectrpc

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kayn1/guidero/gen/proto/user/v1"
	"github.com/kayn1/guidero/internal/domain"
)

func (s *ConnectRpcServer) Create(ctx context.Context, req *connect.Request[v1.CreateRequest]) (*connect.Response[v1.CreateResponse], error) {
	user, err := s.app.UserService.CreateUser(context.Background(), domain.CreateUserRequest{
		Email: req.Msg.Email,
		Name:  req.Msg.Name,
	})

	if err != nil {
		return nil, err
	}

	res := &connect.Response[v1.CreateResponse]{
		Msg: &v1.CreateResponse{
			User: &v1.User{
				Id:    user.ID.String(),
				Email: user.Email,
				Name:  user.Name,
			},
		},
	}

	return res, nil
}
