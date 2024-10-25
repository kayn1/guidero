package internal

import (
	"log/slog"
	"os"
)

var Logger, _ = NewLogger()

type loggerOptions struct {
	name     string
	version  string
	logLevel slog.Level
}

type LoggerOption func(*loggerOptions) error

func NewLogger(opts ...LoggerOption) (*slog.Logger, error) {
	logger := &loggerOptions{
		name:     "logger",
		version:  "0.0.1",
		logLevel: slog.LevelInfo,
	}

	for _, opt := range opts {
		err := opt(logger)
		if err != nil {
			return nil, err
		}
	}

	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logger.logLevel,
	})

	slogger := slog.New(handler)

	slogger = slogger.With(
		slog.String("service", logger.name),
		slog.String("version", logger.version),
	)

	return slogger, nil
}

func WithName(name string) LoggerOption {
	return func(lo *loggerOptions) error {
		lo.name = name
		return nil
	}
}

func WithVersion(version string) LoggerOption {
	return func(lo *loggerOptions) error {
		lo.version = version
		return nil
	}
}

func WithLogLevel(logLevel slog.Level) LoggerOption {
	return func(lo *loggerOptions) error {
		lo.logLevel = logLevel
		return nil
	}
}
