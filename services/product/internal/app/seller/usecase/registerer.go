package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=registerer.go -destination=mock/registerer_mock.go -package=usecasemock

type SellerRegisterer interface {
	Register(ctx context.Context, userId string, form SellerForm) error
}
