package user

import "net/http"

type Handler struct {
	service *UserService
}

func NewHandler(service *UserService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Implementation for serving HTTP requests
}