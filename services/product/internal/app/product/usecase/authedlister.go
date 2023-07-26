package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=authedlister.go -destination=mock/authedlister_mock.go -package=usecasemock

type ProductAuthedLister interface {
	AuthedList(ctx context.Context, userId string, limit, offset int) (ProductList, error)
}
