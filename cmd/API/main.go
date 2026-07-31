package main

import (
	"context"
	"fmt"
	"go-learning/internal/user"
	"net/http"
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

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{id}", user.NewHandler(service).ServeHTTP)

	http.ListenAndServe(":8080", mux)
}
