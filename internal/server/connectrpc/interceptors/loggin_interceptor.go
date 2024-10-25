package interceptors

import (
	"context"
	"log/slog"

	"connectrpc.com/connect"
	crpc "connectrpc.com/connect"
)

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
func NewLoggingInterceptor(logger *slog.Logger) *LoggingInterceptor {
	return &LoggingInterceptor{
		logger: logger,
	}
}

var _ crpc.Interceptor = &LoggingInterceptor{}
