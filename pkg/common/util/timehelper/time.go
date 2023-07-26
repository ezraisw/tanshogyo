package timehelper

import "time"

type (
	Nower interface {
		Now() time.Time
	}

	NowerFunc func() time.Time
)

func (f NowerFunc) Now() time.Time {
	return f()
}

func ProvideNower() Nower {
	return NowerFunc(time.Now)
}
