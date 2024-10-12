package connectrpc

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kayn1/guidero/gen/proto/user/v1"
)

func (s *ConnectRpcServer) Get(context.Context, *connect.Request[v1.GetRequest]) (*connect.Response[v1.GetResponse], error) {
	panic("unimplemented")
}
