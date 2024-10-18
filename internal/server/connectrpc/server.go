package connectrpc

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/kayn1/guidero/gen/proto/user/v1/v1connect"
	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/server"
)

const CONNECTRPC_SERVER = "connectrpc"

var _ server.Server = &ConnectRpcServer{}

type ConnectRpcServer struct {
	app      domain.Application
	logger   *slog.Logger
	listener *http.Server
}

func NewConnectRpcServer(app domain.Application) *ConnectRpcServer {
	logger := newLogger()

	return &ConnectRpcServer{
		app:    app,
		logger: logger,
	}
}

func newLogger() *slog.Logger {
	// Create a JSON handler with common options
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo, // Set the log level to Info
	})

	// Create a new logger with the handler
	logger := slog.New(handler)

	// Add common context fields
	logger = logger.With(
		slog.String("service", "connectrpc"),
		slog.String("version", "1.0.0"),
	)

	return logger
}

func (s *ConnectRpcServer) Start() error {
	mux := http.NewServeMux()
	path, handler := v1connect.NewUserServiceHandler(s)
	mux.Handle(path, handler)

	s.listener = &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Start the server in a goroutine
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
