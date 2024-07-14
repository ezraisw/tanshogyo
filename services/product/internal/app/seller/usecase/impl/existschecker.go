package usecaseimpl

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/entity"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/repository"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
)

type SellerExistsCheckerOptions struct {
	SellerRepository repository.SellerRepository
}

type SellerExistsChecker struct {
	o SellerExistsCheckerOptions
}

func NewSellerExistsChecker(options SellerExistsCheckerOptions) *SellerExistsChecker {
	return &SellerExistsChecker{
		o: options,
	}
}

func (s SellerExistsChecker) CheckExists(ctx context.Context, fields []usecase.Field) (bool, error) {
	exists, err := s.o.SellerRepository.Exists(ctx, s.makeClause(fields))
	if err != nil {
		return false, err
	}
	return exists, err
}

func (s SellerExistsChecker) makeClause(fields []usecase.Field) entity.Clause {
	return func(p entity.Prober) (expr entity.Expression) {
		for i, field := range fields {
			if i > 0 {
				p = expr.And()
			}
			expr = p.Field(field.Name).Is(entity.OperatorEquals, field.Value)
		}
		return
	}
}
