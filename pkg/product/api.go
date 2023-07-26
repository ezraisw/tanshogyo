package product

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=api.go -destination=mock/api_mock.go -package=userauthmock

type ProductAPI interface {
	Get(ctx context.Context, id string) (Product, error)
	Update(ctx context.Context, id string, form ProductForm) (Product, error)
}
