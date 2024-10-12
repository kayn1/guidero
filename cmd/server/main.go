package main

import (
	"log"
	"os"
	"os/signal"

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
		panic(err)
	}

	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint

	err = server.Stop()
	if err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}
}
