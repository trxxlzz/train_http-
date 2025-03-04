package v1

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"training/internal/models"
	"training/internal/service/v1"
)

type UserHandlerV1 struct {
	service v1.UserService
}

func NewUserHandler(s v1.UserService) *UserHandlerV1 {
	return &UserHandlerV1{service: s}
}

func (h *UserHandlerV1) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandlerV1) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByID(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

//func RegisterUserRoutes(r chi.Router, handler *UserHandler) {
//	r.Post("/users", handler.CreateUserHandler)
//	r.Get("/users/{id}", handler.GetUserHandler)
//}
