package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=getter.go -destination=mock/getter_mock.go -package=usecasemock

type SellerGetter interface {
	GetByUserID(ctx context.Context, userId string) (Seller, error)
}
