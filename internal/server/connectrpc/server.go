package connectrpc

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/kayn1/guidero/gen/proto/user/v1/v1connect"
	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/server"
)

const CONNECTRPC_SERVER = "connectrpc"

var _ server.Server = &ConnectRpcServer{}

type ConnectRpcServer struct {
	app    domain.Application
	logger slog.Logger
}

func NewConnectRpcServer(app domain.Application) *ConnectRpcServer {
	return &ConnectRpcServer{
		app:    app,
		logger: *slog.Default(),
	}
}

func (s *ConnectRpcServer) Start() error {
	mux := http.NewServeMux()
	path, handler := v1connect.NewUserServiceHandler(s)
	mux.Handle(path, handler)
	go func() {
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			s.logger.Log(context.Background(), slog.LevelError, "msg", "Failed to start server", "err", err)
		}
	}()

	s.logger.Log(context.Background(), slog.LevelInfo, "msg", "Server has been started", slog.Attr{
		Key:   "port",
		Value: slog.AnyValue(8080),
	})
	return nil
}

func (s *ConnectRpcServer) Stop() error {
	s.logger.Info("Server has been stopped")
	return nil
}

func (s *ConnectRpcServer) ServerType() string {
	return CONNECTRPC_SERVER
}
