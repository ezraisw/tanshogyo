package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=getter.go -destination=mock/getter_mock.go -package=usecasemock

type ProductGetter interface {
	Get(ctx context.Context, id string) (Product, error)
}
