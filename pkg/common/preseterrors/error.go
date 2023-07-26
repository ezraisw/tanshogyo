package preseterrors

import "errors"

var (
	ErrTimeout          = errors.New("timeout")
	ErrInvalidArguments = errors.New("invalid arguments")
	ErrNotFound         = errors.New("not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrForbidden        = errors.New("forbidden")
	ErrInternalProblem  = errors.New("internal problem")
)

type Marshalable interface {
	MarshalAs() (map[string]any, error)
}
