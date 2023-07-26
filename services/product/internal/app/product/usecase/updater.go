package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=updater.go -destination=mock/updater_mock.go -package=usecasemock

type ProductUpdater interface {
	Update(ctx context.Context, id string, form ProductForm) (Product, error)
}
