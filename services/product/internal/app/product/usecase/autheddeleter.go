package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=authedcreator.go -destination=mock/authedcreator_mock.go -package=usecasemock

type ProductAuthedDeleter interface {
	AuthedDelete(ctx context.Context, userId, id string) error
}
