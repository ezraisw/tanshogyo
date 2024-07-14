package usecaseimpl

import (
	"context"
	"errors"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/common/util/idgen"
	"github.com/ezraisw/tanshogyo/pkg/common/util/timehelper"
	"github.com/ezraisw/tanshogyo/pkg/product"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/cache"
	transactionerrors "github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/errors"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/model"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/repository"
)

type TransactionCreatorOptions struct {
	CartCache             cache.CartCache
	ProductAPI            product.ProductAPI
	TransactionRepository repository.TransactionRepository
	Nower                 timehelper.Nower
	IDGen                 idgen.IDGen
}

type TransactionCreator struct {
	o TransactionCreatorOptions
}

func NewTransactionCreator(options TransactionCreatorOptions) *TransactionCreator {
	return &TransactionCreator{
		o: options,
	}
}

func (s TransactionCreator) Create(ctx context.Context, userId string) error {
	cacheCart, err := s.o.CartCache.Get(ctx, userId)
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return transactionerrors.ErrEmptyCart
		}
		return err
	}

	if len(cacheCart.Details) == 0 {
		return transactionerrors.ErrEmptyCart
	}

	now := s.o.Nower.Now()
	transaction := &model.Transaction{
		ID:        s.o.IDGen.Generate(),
		UserID:    userId,
		CreatedAt: now,
		UpdatedAt: now,

		Details: make([]model.TransactionDetail, 0, len(cacheCart.Details)),
	}

	modifiedProducts := make([]product.Product, 0, len(cacheCart.Details))
	for _, detail := range cacheCart.Details {
		p, err := s.o.ProductAPI.Get(ctx, detail.ProductID)
		if err != nil {
			if errors.Is(err, preseterrors.ErrNotFound) {
				return transactionerrors.ErrInvalidCart
			}
			return err
		}

		if p.Quantity < detail.Quantity {
			return transactionerrors.ErrInvalidCart
		}

		transaction.Details = append(transaction.Details, model.TransactionDetail{
			ID:            idgen.ProvideIDGen().Generate(),
			TransactionID: transaction.ID,
			ProductID:     p.ID,
			Price:         p.Price,
			Quantity:      detail.Quantity,
		})

		p.Quantity -= detail.Quantity
		modifiedProducts = append(modifiedProducts, p)
	}

	if _, err := s.o.TransactionRepository.Create(ctx, transaction); err != nil {
		return err
	}

	for _, p := range modifiedProducts {
		if _, err := s.o.ProductAPI.Update(ctx, p.ID, product.ProductForm{
			SellerID:    p.SellerID,
			Name:        p.Name,
			Description: p.Description,
			Price:       p.Price,
			Quantity:    p.Quantity,
		}); err != nil {
			return err
		}
	}

	if err := s.o.CartCache.Delete(ctx, userId); err != nil {
		return err
	}

	return nil
}
