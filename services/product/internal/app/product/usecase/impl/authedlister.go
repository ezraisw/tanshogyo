package usecaseimpl

import (
	"context"
	"errors"

	"github.com/ezraisw/tanshogyo/pkg/common/entity"
	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/repository"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
	sellerusecase "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
)

type ProductAuthedListerOptions struct {
	SellerGetter      sellerusecase.SellerGetter
	ProductRepository repository.ProductRepository
}

type ProductAuthedLister struct {
	o ProductAuthedListerOptions
}

func NewProductAuthedLister(options ProductAuthedListerOptions) *ProductAuthedLister {
	return &ProductAuthedLister{
		o: options,
	}
}

func (s ProductAuthedLister) AuthedList(ctx context.Context, userId string, limit, offset int) (usecase.ProductList, error) {
	seller, err := s.o.SellerGetter.GetByUserID(ctx, userId)
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return usecase.ProductList{}, preseterrors.ErrUnauthorized
		}
		return usecase.ProductList{}, err
	}
	return list(ctx, s.o.ProductRepository, limit, offset, func(p entity.Prober) entity.Expression {
		return p.Field("SellerID").Is(entity.OperatorEquals, seller.ID)
	})
}
