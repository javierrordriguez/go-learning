package main

import (
	"fmt"
	"go-learning/internal/user"
	"net/http"
)

func main() {
	repo := user.NewUserRepository()

	service := user.NewUserService(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{id}", user.NewHandler(service).ServeHTTP)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mux)
}
