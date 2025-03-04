package v2

import (
	"encoding/json"
	"net/http"
	"strconv"
	"training/internal/models"
	"training/internal/service/v2"

	"github.com/go-chi/chi/v5"
)

type UserHandlerV2 struct {
	service v2.UserServiceV2
}

func NewUserHandlerV2(service v2.UserServiceV2) *UserHandlerV2 {
	return &UserHandlerV2{service: service}
}

func (h *UserHandlerV2) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.UserV2
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUserV2(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func (h *UserHandlerV2) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserByIDV2(id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
