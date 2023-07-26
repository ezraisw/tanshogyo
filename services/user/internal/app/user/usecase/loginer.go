package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=loginer.go -destination=mock/loginer_mock.go -package=usecasemock

type UserLoginer interface {
	Login(ctx context.Context, form LoginForm) (AuthenticationResult, error)
}
