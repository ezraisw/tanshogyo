package errors

import "errors"

var (
	ErrEmptyCart   = errors.New("empty cart")
	ErrInvalidCart = errors.New("invalid cart")
)
