package connectrpc

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kayn1/guidero/gen/proto/user/v1"
)

func (s *ConnectRpcServer) Update(context.Context, *connect.Request[v1.UpdateRequest]) (*connect.Response[v1.UpdateResponse], error) {
	panic("unimplemented")
}
