package connectrpc

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/kayn1/guidero/gen/proto/user/v1/v1connect"
	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/server"
)

const CONNECTRPC_SERVER = "connectrpc"

var _ server.Server = &ConnectRpcServer{}

type ConnectRpcServer struct {
	app      domain.Application
	logger   slog.Logger
	listener *http.Server
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

	s.listener = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		err := s.listener.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.listener.Shutdown(ctx); err != nil {
		s.logger.Log(context.Background(), slog.LevelError, "msg", "Failed to stop server", "err", err)
		return err
	}

	s.logger.Log(context.Background(), slog.LevelInfo, "msg", "Server has been stopped", slog.Attr{})
	return nil
}

func (s *ConnectRpcServer) ServerType() string {
	return CONNECTRPC_SERVER
}
