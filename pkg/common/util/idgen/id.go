package idgen

import "github.com/google/uuid"

type (
	IDGen interface {
		Generate() string
	}

	IDGenFunc func() string
)

func (f IDGenFunc) Generate() string {
	return f()
}

func ProvideIDGen() IDGen {
	return IDGenFunc(uuid.NewString)
}
