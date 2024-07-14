package usecaseimpl

import (
	"context"
	"errors"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/product"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/cache"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/usecase"
)

type TransactionCartGetterOptions struct {
	CartCache  cache.CartCache
	ProductAPI product.ProductAPI
}

type TransactionCartGetter struct {
	o TransactionCartGetterOptions
}

func NewTransactionCartGetter(options TransactionCartGetterOptions) *TransactionCartGetter {
	return &TransactionCartGetter{
		o: options,
	}
}

func (s TransactionCartGetter) GetCart(ctx context.Context, userId string) (usecase.CartInfo, error) {
	cacheCart, err := s.o.CartCache.Get(ctx, userId)
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return usecase.CartInfo{Details: []usecase.CartInfoDetail{}}, nil
		}
		return usecase.CartInfo{}, err
	}

	cart := usecase.CartInfo{
		Details: make([]usecase.CartInfoDetail, 0, len(cacheCart.Details)),
	}

	for _, detail := range cacheCart.Details {
		p, err := s.o.ProductAPI.Get(ctx, detail.ProductID)
		if err != nil && !errors.Is(err, preseterrors.ErrNotFound) { // Zero price if not found.
			return usecase.CartInfo{}, err
		}

		subTotal := int(p.Price) * detail.Quantity
		cart.TotalPrice += subTotal

		cart.Details = append(cart.Details, usecase.CartInfoDetail{
			SubTotal:  int(p.Price) * detail.Quantity,
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
		})
	}

	return cart, nil
}
