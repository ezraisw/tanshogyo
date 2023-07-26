package web

import "github.com/go-chi/chi/v5"

type HTTPHandlerRegistry interface {
	RegisterRoutes(chi.Router) error
}

type HTTPMiddlewareRegistry interface {
	GetMiddlewares() chi.Middlewares
}
