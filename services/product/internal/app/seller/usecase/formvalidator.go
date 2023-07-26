package usecase

import "context"

//go:generate go run github.com/golang/mock/mockgen -source=formvalidator.go -destination=mock/formvalidator_mock.go -package=usecasemock

type SellerFormValidator interface {
	Validate(ctx context.Context, form SellerForm) error
}
