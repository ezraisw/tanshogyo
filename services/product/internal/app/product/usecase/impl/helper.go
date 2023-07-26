package usecaseimpl

import (
	"context"

	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	commonrepository "github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
)

func fromForm(form usecase.ProductForm) *model.Product {
	return &model.Product{
		SellerID:    form.SellerID,
		Name:        form.Name,
		Description: form.Description,
		Price:       form.Price,
		Quantity:    form.Quantity,
	}
}

func toDto(product *model.Product) usecase.Product {
	return usecase.Product{
		ID:          product.ID,
		SellerID:    product.SellerID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}
}

func list(ctx context.Context, r repository.ProductRepository, limit, offset int, clause entity.Clause) (usecase.ProductList, error) {
	count, err := r.Count(ctx, clause)
	if err != nil {
		return usecase.ProductList{}, err
	}

	// Revert to defaults when invalid values are given.
	if limit <= 0 {
		limit = LimitDefault
	}

	if offset < 0 {
		offset = OffsetDefault
	}

	// Do not return anything if offset exceeds the count.
	// Reduces the amount of database connections.
	if count <= offset {
		return usecase.ProductList{
			Count:  count,
			Limit:  limit,
			Offset: offset,
			Data:   []usecase.Product{},
		}, nil
	}

	products, err := r.FindMany(ctx, clause, commonrepository.FindManyOptions{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return usecase.ProductList{}, err
	}

	dtos := make([]usecase.Product, 0, len(products))
	for _, product := range products {
		dtos = append(dtos, toDto(product))
	}

	return usecase.ProductList{
		Count:  count,
		Limit:  limit,
		Offset: offset,
		Data:   dtos,
	}, nil
}
