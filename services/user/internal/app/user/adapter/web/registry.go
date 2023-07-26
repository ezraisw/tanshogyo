package web

import (
	"github.com/go-chi/chi/v5"
)

type UserHandlerRegistryOptions struct {
	UserController *UserController
}

type UserHandlerRegistry struct {
	o UserHandlerRegistryOptions
}

func NewUserHandlerRegistry(options UserHandlerRegistryOptions) *UserHandlerRegistry {
	return &UserHandlerRegistry{
		o: options,
	}
}

func (h UserHandlerRegistry) RegisterRoutes(r chi.Router) error {
	r.Route("/v1/user", func(r chi.Router) {
		r.Get("/", h.o.UserController.AuthenticateHandler)
		r.Post("/login", h.o.UserController.LoginHandler)
		r.Post("/register", h.o.UserController.RegisterHandler)
	})
	return nil
}
