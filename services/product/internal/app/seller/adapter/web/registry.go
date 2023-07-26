package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/pwnedgod/tanshogyo/pkg/userauth"
)

type SellerHandlerRegistryOptions struct {
	SellerController   *SellerController
	UserAuthMiddleware userauth.UserAuthMiddleware
}

type SellerHandlerRegistry struct {
	o SellerHandlerRegistryOptions
}

func NewSellerHandlerRegistry(options SellerHandlerRegistryOptions) *SellerHandlerRegistry {
	return &SellerHandlerRegistry{
		o: options,
	}
}

func (h SellerHandlerRegistry) RegisterRoutes(r chi.Router) error {
	r.Route("/v1/seller", func(r chi.Router) {
		r.Use(h.o.UserAuthMiddleware)
		r.Get("/", h.o.SellerController.GetByUserIDHandler)
		r.Post("/", h.o.SellerController.RegisterHandler)
	})
	return nil
}
