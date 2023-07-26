package usecaseimpl

import (
	"context"

	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	commonrepository "github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
)

type ProductGetterOptions struct {
	ProductRepository repository.ProductRepository
}

type ProductGetter struct {
	o ProductGetterOptions
}

func NewProductGetter(options ProductGetterOptions) *ProductGetter {
	return &ProductGetter{
		o: options,
	}
}

func (s ProductGetter) Get(ctx context.Context, id string) (usecase.Product, error) {
	product, err := s.o.ProductRepository.Find(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("ID").Is(entity.OperatorEquals, id)
	}, commonrepository.FindOptions{})
	if err != nil {
		return usecase.Product{}, err
	}
	return toDto(product), nil
}
