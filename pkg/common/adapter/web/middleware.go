package web

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type EssentialsMiddlewareRegistry struct{}

func NewEssentialsMiddlewareRegistry() *EssentialsMiddlewareRegistry {
	return &EssentialsMiddlewareRegistry{}
}

func (r EssentialsMiddlewareRegistry) GetMiddlewares() chi.Middlewares {
	return chi.Middlewares{
		middleware.Recoverer,
		middleware.CleanPath,
		middleware.RedirectSlashes,
		middleware.RealIP,
	}
}
