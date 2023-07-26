package usecaseimpl

import (
	"context"

	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	commonrepository "github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
)

type SellerGetterOptions struct {
	SellerRepository repository.SellerRepository
}

type SellerGetter struct {
	o SellerGetterOptions
}

func NewSellerGetter(options SellerGetterOptions) *SellerGetter {
	return &SellerGetter{
		o: options,
	}
}

func (s SellerGetter) GetByUserID(ctx context.Context, userId string) (usecase.Seller, error) {
	seller, err := s.o.SellerRepository.Find(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("UserID").Is(entity.OperatorEquals, userId)
	}, commonrepository.FindOptions{})
	if err != nil {
		return usecase.Seller{}, err
	}
	return toDto(seller), nil
}
