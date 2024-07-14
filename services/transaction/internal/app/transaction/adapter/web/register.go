package web

import (
	"github.com/ezraisw/tanshogyo/pkg/userauth"
	"github.com/go-chi/chi/v5"
)

type TransactionHandlerRegistryOptions struct {
	TransactionController *TransactionController
	UserAuthMiddleware    userauth.UserAuthMiddleware
}

type TransactionHandlerRegistry struct {
	o TransactionHandlerRegistryOptions
}

func NewTransactionHandlerRegistry(options TransactionHandlerRegistryOptions) *TransactionHandlerRegistry {
	return &TransactionHandlerRegistry{
		o: options,
	}
}

func (h TransactionHandlerRegistry) RegisterRoutes(r chi.Router) error {
	r.Route("/v1/transaction", func(r chi.Router) {
		r.Use(h.o.UserAuthMiddleware)

		r.Get("/", h.o.TransactionController.ListHandler)
		r.Post("/", h.o.TransactionController.CreateHandler)

		r.Route("/cart", func(r chi.Router) {
			r.Get("/", h.o.TransactionController.GetCartHandler)
			r.Put("/", h.o.TransactionController.UpdateCartHandler)
		})
	})

	return nil
}
