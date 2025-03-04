package v2

import (
	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutesV2(r chi.Router, handler *UserHandlerV2) {
	r.Route("/users", func(r chi.Router) { // тут не надо `v2` — он будет выше, при регистрации
		r.Post("/", handler.CreateUser)
		r.Get("/{id}", handler.GetUserByID)
	})
}
