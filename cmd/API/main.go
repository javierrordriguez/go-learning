package main

import (
	"context"
	"fmt"
	"go-learning/internal/user"
)

func main() {
	repo := user.NewUserRepository()

	service := user.NewUserService(repo)

	user, err := service.GetUser(context.Background(), 1)

	if err != nil {
		// Handle error
	}

	if user != nil {
		fmt.Printf("User: %+v\n", user)
	}
}
