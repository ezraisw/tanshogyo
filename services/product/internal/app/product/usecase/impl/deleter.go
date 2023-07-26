package usecaseimpl

import (
	"context"

	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/repository"
)

type ProductDeleterOptions struct {
	ProductRepository repository.ProductRepository
}

type ProductDeleter struct {
	o ProductDeleterOptions
}

func NewProductDeleter(options ProductDeleterOptions) *ProductDeleter {
	return &ProductDeleter{
		o: options,
	}
}

func (s ProductDeleter) Delete(ctx context.Context, id string) error {
	return s.o.ProductRepository.Delete(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("ID").Is(entity.OperatorEquals, id)
	})
}
