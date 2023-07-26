package cache

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=cart.go -destination=mock/cart_mock.go -package=cachemock

type CartCache interface {
	Get(ctx context.Context, userId string) (Cart, error)
	Set(ctx context.Context, userId string, cart Cart) error
	Delete(ctx context.Context, userId string) error
}
