package gormds

import "errors"

var (
	ErrInvalidConnection = errors.New("invalid connection")
	ErrInvalidDriver     = errors.New("invalid driver")
)
