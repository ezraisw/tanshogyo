package usecaseimpl

import (
	"context"
	"errors"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	sellerusecase "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
)

type ProductAuthedUpdaterOptions struct {
	SellerGetter   sellerusecase.SellerGetter
	ProductUpdater usecase.ProductUpdater
}

type ProductAuthedUpdater struct {
	o ProductAuthedUpdaterOptions
}

func NewProductAuthedUpdater(options ProductAuthedUpdaterOptions) *ProductAuthedUpdater {
	return &ProductAuthedUpdater{
		o: options,
	}
}

func (s ProductAuthedUpdater) AuthedUpdate(ctx context.Context, userId, id string, form usecase.AuthedProductForm) (usecase.Product, error) {
	seller, err := s.o.SellerGetter.GetByUserID(ctx, userId)
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return usecase.Product{}, preseterrors.ErrUnauthorized
		}
		return usecase.Product{}, err
	}
	return s.o.ProductUpdater.Update(ctx, id, usecase.ProductForm{
		SellerID:    seller.ID,
		Name:        form.Name,
		Description: form.Description,
		Price:       form.Price,
		Quantity:    form.Quantity,
	})
}
