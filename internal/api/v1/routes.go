package v1

import (
	"github.com/go-chi/chi/v5"
)

func RegisterUserRoutesV1(r chi.Router, handler *UserHandlerV1) {
	r.Route("/users", func(r chi.Router) {
		r.Post("/", handler.CreateUserHandler)
		r.Get("/{id}", handler.GetUserHandler)
	})
}
