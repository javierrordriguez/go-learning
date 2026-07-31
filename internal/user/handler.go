package user

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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
	userID, err := strconv.ParseInt(r.PathValue("id"), 10, 64) // Assuming you have a way to extract the user ID from the request path

	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(r.Context(), userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	http.ResponseWriter.WriteHeader(w, http.StatusOK)
	// Assuming you have a way to serialize the user object to JSON
	json.NewEncoder(w).Encode(user)

}
