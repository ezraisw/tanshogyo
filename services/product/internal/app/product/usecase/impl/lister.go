package usecaseimpl

import (
	"context"

	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
)

type ProductListerOptions struct {
	ProductRepository repository.ProductRepository
}

type ProductLister struct {
	o ProductListerOptions
}

func NewProductLister(options ProductListerOptions) *ProductLister {
	return &ProductLister{
		o: options,
	}
}

func (s ProductLister) List(ctx context.Context, limit, offset int) (usecase.ProductList, error) {
	return list(ctx, s.o.ProductRepository, limit, offset, nil)
}
