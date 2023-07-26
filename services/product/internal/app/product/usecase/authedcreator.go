package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=authedcreator.go -destination=mock/authedcreator_mock.go -package=usecasemock

type ProductAuthedCreator interface {
	AuthedCreate(ctx context.Context, userId string, form AuthedProductForm) (Product, error)
}
