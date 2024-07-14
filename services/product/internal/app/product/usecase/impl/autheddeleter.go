package usecaseimpl

import (
	"context"
	"errors"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	sellerusecase "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
)

type ProductAuthedDeleterOptions struct {
	SellerGetter   sellerusecase.SellerGetter
	ProductDeleter usecase.ProductDeleter
}

type ProductAuthedDeleter struct {
	o ProductAuthedDeleterOptions
}

func NewProductAuthedDeleter(options ProductAuthedDeleterOptions) *ProductAuthedDeleter {
	return &ProductAuthedDeleter{
		o: options,
	}
}

func (s ProductAuthedDeleter) AuthedDelete(ctx context.Context, userId, id string) error {
	_, err := s.o.SellerGetter.GetByUserID(ctx, userId)
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return preseterrors.ErrUnauthorized
		}
		return err
	}
	return s.o.ProductDeleter.Delete(ctx, id)
}
