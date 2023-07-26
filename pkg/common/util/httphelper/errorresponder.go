package httphelper

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
)

type (
	errStatusCode struct {
		matcher    preseterrors.Matcher
		statusCode int
	}
)

var errStatusCodes []errStatusCode = make([]errStatusCode, 0)

func RespondError(w http.ResponseWriter, r *http.Request, err error) {
	var body any
	if marshalableErr, ok := err.(preseterrors.Marshalable); ok {
		var marshalErr error
		if body, marshalErr = marshalableErr.MarshalAs(); marshalErr != nil {
			err = marshalErr
		}
	}

	if body == nil {
		body = map[string]string{
			"message": err.Error(),
		}
	}

	render.Status(r, GetStatusCode(err))
	render.DefaultResponder(w, r, body)
}

func GetStatusCode(err error) int {
	for _, esc := range errStatusCodes {
		if esc.matcher(err) {
			return esc.statusCode
		}
	}
	return http.StatusInternalServerError
}

func RegisterError(matcher preseterrors.Matcher, statusCode int) {
	errStatusCodes = append(errStatusCodes, errStatusCode{matcher: matcher, statusCode: statusCode})
}
