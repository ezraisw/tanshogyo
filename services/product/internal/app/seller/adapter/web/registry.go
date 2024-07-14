package web

import (
	"github.com/ezraisw/tanshogyo/pkg/userauth"
	"github.com/go-chi/chi/v5"
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
