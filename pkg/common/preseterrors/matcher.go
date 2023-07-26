package preseterrors

import "errors"

type (
	Matcher func(error) bool
)

func ErrIs(target error) Matcher {
	return func(err error) bool { return errors.Is(err, target) }
}
