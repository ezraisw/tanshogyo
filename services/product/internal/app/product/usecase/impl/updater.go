package usecaseimpl

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/entity"
	commonrepository "github.com/ezraisw/tanshogyo/pkg/common/repository"
	"github.com/ezraisw/tanshogyo/pkg/common/util/timehelper"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/repository"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/usecase"
)

type ProductUpdaterOptions struct {
	ProductFormValidator usecase.ProductFormValidator
	ProductRepository    repository.ProductRepository
	Nower                timehelper.Nower
}

type ProductUpdater struct {
	o ProductUpdaterOptions
}

func NewProductUpdater(options ProductUpdaterOptions) *ProductUpdater {
	return &ProductUpdater{
		o: options,
	}
}

func (s ProductUpdater) Update(ctx context.Context, id string, form usecase.ProductForm) (usecase.Product, error) {
	err := s.o.ProductFormValidator.Validate(ctx, form)
	if err != nil {
		return usecase.Product{}, err
	}

	product, err := s.o.ProductRepository.Find(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("ID").Is(entity.OperatorEquals, id)
	}, commonrepository.FindOptions{})
	if err != nil {
		return usecase.Product{}, err
	}

	now := s.o.Nower.Now().UTC()
	product.SellerID = form.SellerID
	product.Name = form.Name
	product.Description = form.Description
	product.Price = form.Price
	product.Quantity = form.Quantity
	product.UpdatedAt = now

	if err := s.o.ProductRepository.Update(ctx, product); err != nil {
		return usecase.Product{}, err
	}

	return toDto(product), nil
}
