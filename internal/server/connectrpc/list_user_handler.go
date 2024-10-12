package connectrpc

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kayn1/guidero/gen/proto/user/v1"
)

func (s *ConnectRpcServer) List(context.Context, *connect.Request[v1.ListRequest]) (*connect.Response[v1.ListResponse], error) {
	panic("unimplemented")
}
