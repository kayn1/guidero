package connectrpc

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"time"

	"connectrpc.com/connect"
	crpc "connectrpc.com/connect"
	"github.com/kayn1/guidero/gen/proto/user/v1/v1connect"
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
	logger := newLogger()

	server := &ConnectRpcServer{
		app:          app,
		logger:       logger,
		interceptors: []crpc.Interceptor{},
	}

	for _, opt := range opts {
		opt(server)
	}

	return server
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

type LoggingInterceptor struct {
	logger *slog.Logger
}

// WrapStreamingClient implements connect.Interceptor.
func (l *LoggingInterceptor) WrapStreamingClient(crpc.StreamingClientFunc) crpc.StreamingClientFunc {
	panic("unimplemented")
}

// WrapStreamingHandler implements connect.Interceptor.
func (l *LoggingInterceptor) WrapStreamingHandler(crpc.StreamingHandlerFunc) crpc.StreamingHandlerFunc {
	// Log the request and response of the streaming RPC
	return func(ctx context.Context, conn crpc.StreamingHandlerConn) error {
		l.logger.Info("Streaming RPC started")

		err := conn.Receive(ctx)
		if err != nil {
			l.logger.Error("Streaming RPC failed", slog.String("error", err.Error()))
			return err
		}

		l.logger.Info("Streaming RPC completed")
		return nil
	}
}

func (l *LoggingInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
		l.logger.Info("Unary RPC started",
			slog.String("procedure", req.Spec().Procedure),
			slog.String("peer", req.Peer().Addr))

		resp, err := next(ctx, req)
		if err != nil {
			l.logger.Error("Unary RPC failed",
				slog.String("procedure", req.Spec().Procedure),
				slog.String("error", err.Error()))
			return nil, err
		}

		l.logger.Info("Unary RPC completed",
			slog.String("procedure", req.Spec().Procedure),
		)
		return resp, nil
	}
}
func NewLoggingInterceptor() *LoggingInterceptor {
	return &LoggingInterceptor{
		logger: newLogger(),
	}
}

var _ crpc.Interceptor = &LoggingInterceptor{}
