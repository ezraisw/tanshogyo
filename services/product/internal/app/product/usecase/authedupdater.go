package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=autheddeleter.go -destination=mock/autheddeleter_mock.go -package=usecasemock

type ProductAuthedUpdater interface {
	AuthedUpdate(ctx context.Context, userId, id string, form AuthedProductForm) (Product, error)
}
