package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=registerer.go -destination=mock/registerer_mock.go -package=usecasemock

type UserRegisterer interface {
	Register(ctx context.Context, form UserForm) error
}
