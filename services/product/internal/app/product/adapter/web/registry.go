package web

import (
	"github.com/ezraisw/tanshogyo/pkg/userauth"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/middleware"
	"github.com/go-chi/chi/v5"
)

type ProductHandlerRegistryOptions struct {
	ProductController       *ProductController
	UserAuthMiddleware      userauth.UserAuthMiddleware
	SellerCheckerMiddleware middleware.SellerCheckerMiddleware
}

type ProductHandlerRegistry struct {
	o ProductHandlerRegistryOptions
}

func NewProductHandlerRegistry(options ProductHandlerRegistryOptions) *ProductHandlerRegistry {
	return &ProductHandlerRegistry{
		o: options,
	}
}

func (h ProductHandlerRegistry) RegisterRoutes(r chi.Router) error {
	r.Route("/v1/seller/product", func(r chi.Router) {
		r.Use(h.o.UserAuthMiddleware)
		r.Get("/", h.o.ProductController.AuthedListHandler)
		r.Post("/", h.o.ProductController.AuthedCreateHandler)

		r.Group(func(r chi.Router) {
			r.Use(h.o.SellerCheckerMiddleware)
			r.Put("/{id}", h.o.ProductController.AuthedUpdateHandler)
			r.Delete("/{id}", h.o.ProductController.AuthedDeleteHandler)
		})
	})

	r.Route("/v1/product", func(r chi.Router) {
		r.Get("/", h.o.ProductController.ListHandler)
		r.Get("/{id}", h.o.ProductController.GetHandler)

		// Debugging routes.
		// r.Post("/", h.o.ProductController.CreateHandler)
		// r.Put("/{id}", h.o.ProductController.UpdateHandler)
		// r.Delete("/{id}", h.o.ProductController.DeleteHandler)
	})
	return nil
}
