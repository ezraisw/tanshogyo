package usecaseimpl

import (
	"context"
	"errors"

	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
	sellerusecase "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
)

type ProductAuthedCreatorOptions struct {
	SellerGetter   sellerusecase.SellerGetter
	ProductCreator usecase.ProductCreator
}

type ProductAuthedCreator struct {
	o ProductAuthedCreatorOptions
}

func NewProductAuthedCreator(options ProductAuthedCreatorOptions) *ProductAuthedCreator {
	return &ProductAuthedCreator{
		o: options,
	}
}

func (s ProductAuthedCreator) AuthedCreate(ctx context.Context, userId string, form usecase.AuthedProductForm) (usecase.Product, error) {
	seller, err := s.o.SellerGetter.GetByUserID(ctx, userId)
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return usecase.Product{}, preseterrors.ErrUnauthorized
		}
		return usecase.Product{}, err
	}
	return s.o.ProductCreator.Create(ctx, usecase.ProductForm{
		SellerID:    seller.ID,
		Name:        form.Name,
		Description: form.Description,
		Price:       form.Price,
		Quantity:    form.Quantity,
	})
}
