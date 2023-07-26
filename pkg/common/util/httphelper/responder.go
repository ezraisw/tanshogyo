package httphelper

import (
	"net/http"
	"strings"

	"github.com/go-chi/render"
)

func RespondStatusCode(w http.ResponseWriter, r *http.Request, statusCode int) {
	render.Status(r, statusCode)
	render.DefaultResponder(w, r, map[string]any{
		"message": strings.ToLower(http.StatusText(statusCode)),
	})
}

func Respond(w http.ResponseWriter, r *http.Request, statusCode int, body any) {
	render.Status(r, statusCode)
	render.DefaultResponder(w, r, body)
}
