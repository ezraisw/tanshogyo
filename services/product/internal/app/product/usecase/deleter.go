package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=deleter.go -destination=mock/deleter_mock.go -package=usecasemock

type ProductDeleter interface {
	Delete(ctx context.Context, id string) error
}
