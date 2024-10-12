package connectrpc

import (
	"context"

	"connectrpc.com/connect"
	v1 "github.com/kayn1/guidero/gen/proto/user/v1"
)

func (s *ConnectRpcServer) Delete(context.Context, *connect.Request[v1.DeleteRequest]) (*connect.Response[v1.DeleteResponse], error) {
	panic("unimplemented")
}
