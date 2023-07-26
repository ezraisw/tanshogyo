package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=cartgetter.go -destination=mock/cartgetter_mock.go -package=usecasemock

type TransactionCartGetter interface {
	GetCart(ctx context.Context, userId string) (CartInfo, error)
}
