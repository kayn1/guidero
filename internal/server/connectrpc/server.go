package connectrpc

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	crpc "connectrpc.com/connect"
	"github.com/kayn1/guidero/gen/proto/user/v1/v1connect"
	"github.com/kayn1/guidero/internal"
	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/server"
)

const CONNECTRPC_SERVER = "connectrpc"

var _ server.Server = &ConnectRpcServer{}

type ConnectRpcServer struct {
	app          domain.Application
	logger       *slog.Logger
	listener     *http.Server
	interceptors []crpc.Interceptor
}

type ConnectRpcServerOption func(*ConnectRpcServer)

func NewConnectRpcServer(app domain.Application, opts ...ConnectRpcServerOption) *ConnectRpcServer {
	server := &ConnectRpcServer{
		app:          app,
		logger:       internal.Logger,
		interceptors: []crpc.Interceptor{},
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (s *ConnectRpcServer) Start() error {
	mux := http.NewServeMux()
	path, handler := v1connect.NewUserServiceHandler(
		s,
		crpc.WithInterceptors(s.interceptors...),
	)
	mux.Handle(path, handler)

	s.listener = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		err := s.listener.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			s.logger.Error("Failed to start server", slog.String("error", err.Error()))
		}
	}()

	s.logger.Info("Server has been started", slog.Int("port", 8080))

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	return s.Stop()
}

func (s *ConnectRpcServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.listener.Shutdown(ctx); err != nil {
		s.logger.Error("Failed to stop server", slog.String("error", err.Error()))
		return err
	}

	s.logger.Info("Server has been stopped")
	return nil
}

func (s *ConnectRpcServer) ServerType() string {
	return CONNECTRPC_SERVER
}

func WithInterceptors(interceptors ...crpc.Interceptor) ConnectRpcServerOption {
	return func(s *ConnectRpcServer) {
		s.interceptors = interceptors
	}
}

func WithLogger(logger *slog.Logger) ConnectRpcServerOption {
	return func(s *ConnectRpcServer) {
		s.logger = logger
	}
}
