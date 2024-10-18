package main

import (
	"log"

	"github.com/kayn1/guidero/internal/domain"
	"github.com/kayn1/guidero/internal/domain/users"
	"github.com/kayn1/guidero/internal/repository/inmemory"
	"github.com/kayn1/guidero/internal/server/connectrpc"
)

func main() {
	repo := inmemory.NewRepository()
	userService := users.NewUserService(repo)
	app := domain.NewApplication(userService)
	server := connectrpc.NewConnectRpcServer(app)

	err := server.Start()
	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
