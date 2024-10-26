package main

import (
	"os"

	"github.com/kayn1/guidero/internal"
	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/domain/users"
	"github.com/kayn1/guidero/internal/repository/inmemory"
	"github.com/kayn1/guidero/internal/server/connectrpc"
	"github.com/kayn1/guidero/internal/server/connectrpc/interceptors"
)

func main() {
	repo := inmemory.NewRepository()
	userService := users.NewUserService(repo)
	app := domain.NewApplication(userService)
	server := connectrpc.NewConnectRpcServer(app,
		connectrpc.WithLogger(internal.Logger),
		connectrpc.WithInterceptors(
			interceptors.NewLoggingInterceptor(internal.Logger),
		),
	)

	err := server.Start()
	if err != nil {
		internal.Logger.Error("failed to start server", "error", err)
		os.Exit(1)
	}
}
