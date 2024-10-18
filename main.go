package main

import (
	"fmt"

	"github.com/kayn1/guidero/pkg/client"
)

func main() {
	client, err := client.NewClient()
	if err != nil {
		panic(err)
	}

	res, err := client.CreateUser(
		"test@user.com",
		"Test User",
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("User created: %v\n", res)
}
