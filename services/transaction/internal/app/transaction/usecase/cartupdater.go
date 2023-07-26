package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=cartupdater.go -destination=mock/cartupdater_mock.go -package=usecasemock

type TransactionCartUpdater interface {
	UpdateCart(ctx context.Context, userId string, cart Cart) error
}
