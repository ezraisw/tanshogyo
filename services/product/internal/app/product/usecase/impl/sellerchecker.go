package usecaseimpl

import (
	"context"
	"errors"

	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	commonrepository "github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/repository"
	sellerusecase "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
)

type ProductSellerCheckerOptions struct {
	SellerGetter      sellerusecase.SellerGetter
	ProductRepository repository.ProductRepository
}

type ProductSellerChecker struct {
	o ProductSellerCheckerOptions
}

func NewProductSellerChecker(options ProductSellerCheckerOptions) *ProductSellerChecker {
	return &ProductSellerChecker{
		o: options,
	}
}

func (s ProductSellerChecker) CheckSeller(ctx context.Context, userId, id string) error {
	seller, err := s.o.SellerGetter.GetByUserID(ctx, userId)
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return preseterrors.ErrUnauthorized
		}
		return err
	}

	product, err := s.o.ProductRepository.Find(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("ID").Is(entity.OperatorEquals, id)
	}, commonrepository.FindOptions{})
	if err != nil {
		return err
	}

	if product.SellerID != seller.ID {
		return preseterrors.ErrForbidden
	}

	return nil
}
