package golearning

import (
	"context"
	"fmt"
	"go-learning/internal/repository/memory"
	"go-learning/internal/service"
)

func main() {
	repo := memory.NewUserRepository()

	service := service.NewUserService(repo)

	user, err := service.GetUser(context.Background(), 1)

	if err != nil {
		// Handle error
	}

	if user != nil {
		fmt.Printf("User: %+v\n", user)
	}
}
