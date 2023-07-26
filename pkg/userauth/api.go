package userauth

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=api.go -destination=mock/api_mock.go -package=userauthmock

type UserAPI interface {
	Authenticate(ctx context.Context, token string) (User, error)
}
