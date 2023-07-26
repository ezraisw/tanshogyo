package usecaseimpl

import (
	"context"
	"errors"

	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"github.com/pwnedgod/tanshogyo/pkg/product"
	"github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/cache"
	transactionerrors "github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/errors"
	"github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/usecase"
)

type TransactionCartUpdaterOptions struct {
	CartCache  cache.CartCache
	ProductAPI product.ProductAPI
}

type TransactionCartUpdater struct {
	o TransactionCartUpdaterOptions
}

func NewTransactionCartUpdater(options TransactionCartUpdaterOptions) *TransactionCartUpdater {
	return &TransactionCartUpdater{
		o: options,
	}
}

func (s TransactionCartUpdater) UpdateCart(ctx context.Context, userId string, cart usecase.Cart) error {
	cacheCart := cache.Cart{
		// At least as an array.
		Details: make([]cache.CartDetail, 0, len(cart.Details)),
	}

	productIds := make(map[string]bool)
	for _, detail := range cart.Details {
		if err := s.checkDetail(ctx, productIds, detail); err != nil {
			return err
		}

		cacheCart.Details = append(cacheCart.Details, cache.CartDetail(detail))
	}

	if err := s.o.CartCache.Set(ctx, userId, cacheCart); err != nil {
		return err
	}

	return nil
}

func (s TransactionCartUpdater) checkDetail(ctx context.Context, productIds map[string]bool, detail usecase.CartDetail) error {
	if detail.Quantity < 0 {
		return transactionerrors.ErrInvalidCart
	}

	if _, ok := productIds[detail.ProductID]; ok {
		return transactionerrors.ErrInvalidCart
	}

	p, err := s.o.ProductAPI.Get(ctx, detail.ProductID)
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return transactionerrors.ErrInvalidCart
		}
		return err
	}

	if detail.Quantity > p.Quantity {
		return transactionerrors.ErrInvalidCart
	}

	productIds[detail.ProductID] = true
	return nil
}
