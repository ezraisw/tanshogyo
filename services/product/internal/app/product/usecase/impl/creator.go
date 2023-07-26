package usecaseimpl

import (
	"context"

	"github.com/pwnedgod/tanshogyo/pkg/common/util/idgen"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
)

type ProductCreatorOptions struct {
	ProductFormValidator usecase.ProductFormValidator
	ProductRepository    repository.ProductRepository
	Nower                timehelper.Nower
	IDGen                idgen.IDGen
}

type ProductCreator struct {
	o ProductCreatorOptions
}

func NewProductCreator(options ProductCreatorOptions) *ProductCreator {
	return &ProductCreator{
		o: options,
	}
}

func (s ProductCreator) Create(ctx context.Context, form usecase.ProductForm) (usecase.Product, error) {
	err := s.o.ProductFormValidator.Validate(ctx, form)
	if err != nil {
		return usecase.Product{}, err
	}

	now := s.o.Nower.Now().UTC()
	id := s.o.IDGen.Generate()

	product := fromForm(form)
	product.ID = id
	product.CreatedAt = now
	product.UpdatedAt = now

	if _, err := s.o.ProductRepository.Create(ctx, product); err != nil {
		return usecase.Product{}, err
	}

	return toDto(product), nil
}
