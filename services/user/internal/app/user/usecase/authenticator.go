package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=authenticator.go -destination=mock/authenticator_mock.go -package=usecasemock

type UserAuthenticator interface {
	Authenticate(ctx context.Context, token string) (User, error)
}
