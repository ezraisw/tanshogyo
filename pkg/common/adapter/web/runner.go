package web

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/reflhelper"
	"github.com/rs/zerolog"
)

type WebRunnerOptions struct {
	Logger *zerolog.Logger
}

type WebRunner struct {
	o WebRunnerOptions

	router chi.Router
	server *http.Server
}

func NewWebRunner(configGetter HTTPConfigGetter, options WebRunnerOptions) *WebRunner {
	router := chi.NewRouter()
	router.NotFound(notFoundHandler)
	server := &http.Server{
		Addr:    configGetter.GetHTTPConfig().GetAddress(),
		Handler: router,
	}
	return &WebRunner{
		o:      options,
		router: router,
		server: server,
	}
}

func (r WebRunner) Run(context.Context) error {
	r.o.Logger.Info().
		Str("addr", r.server.Addr).
		Msg("starting web server")

	if err := r.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	r.o.Logger.Info().
		Msg("exited web server")

	return nil
}

func (r WebRunner) Stop(ctx context.Context) error {
	r.o.Logger.Info().
		Msg("shutting down web server")

	return r.server.Shutdown(ctx)
}

func (r WebRunner) AddHandlerRegistries(catalog *reflhelper.StructCatalog) error {
	for _, c := range reflhelper.Collect[HTTPHandlerRegistry](catalog) {
		r.o.Logger.Debug().
			Str("name", c.Name).
			Msg("adding http handler registry")
		registry := c.Value
		if err := registry.RegisterRoutes(r.router); err != nil {
			return err
		}
	}

	routes := r.router.Routes()
	r.o.Logger.Info().
		Int("count", len(routes)).
		Msg("routes registered")

	return nil
}

func (r WebRunner) AddMiddlewareRegistries(catalog *reflhelper.StructCatalog) error {
	for _, c := range reflhelper.Collect[HTTPMiddlewareRegistry](catalog) {
		r.o.Logger.Debug().
			Str("name", c.Name).
			Msg("adding http middleware registry")
		registry := c.Value
		r.router.Use(registry.GetMiddlewares()...)
	}

	middlewares := r.router.Middlewares()
	r.o.Logger.Info().
		Int("count", len(middlewares)).
		Msg("global middlewares registered")

	return nil
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusNotFound)
	render.DefaultResponder(w, r, map[string]string{"message": preseterrors.ErrNotFound.Error()})
}
