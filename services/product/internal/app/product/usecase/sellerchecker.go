package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=sellerchecker.go -destination=mock/sellerchecker_mock.go -package=usecasemock

type ProductSellerChecker interface {
	CheckSeller(ctx context.Context, userId, id string) error
}
