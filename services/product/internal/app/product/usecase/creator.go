package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=creator.go -destination=mock/creator_mock.go -package=usecasemock

type ProductCreator interface {
	Create(ctx context.Context, form ProductForm) (Product, error)
}
