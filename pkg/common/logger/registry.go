package logger

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type LoggerMiddlewareRegistryOptions struct {
	Logger *zerolog.Logger
}

type LoggerMiddlewareRegistry struct {
	o LoggerMiddlewareRegistryOptions
}

func NewLoggerMiddlewareRegistry(options LoggerMiddlewareRegistryOptions) *LoggerMiddlewareRegistry {
	return &LoggerMiddlewareRegistry{
		o: options,
	}
}

func (m LoggerMiddlewareRegistry) GetMiddlewares() chi.Middlewares {
	return chi.Middlewares{
		m.loggerMiddleware,
	}
}

func (m LoggerMiddlewareRegistry) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.o.Logger.Info().
			Str("method", r.Method).
			Str("uri", r.RequestURI).
			Str("host", r.Host).
			Msg("http request")

		next.ServeHTTP(w, r)
	})
}
