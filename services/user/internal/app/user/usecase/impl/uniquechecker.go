package usecaseimpl

import (
	"context"

	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/repository"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
)

type UserUniqueCheckerOptions struct {
	UserRepository repository.UserRepository
}

type UserUniqueChecker struct {
	o UserUniqueCheckerOptions
}

func NewUserUniqueChecker(options UserUniqueCheckerOptions) *UserUniqueChecker {
	return &UserUniqueChecker{
		o: options,
	}
}

func (s UserUniqueChecker) CheckUnique(ctx context.Context, excludedId string, fields []usecase.Field) (bool, error) {
	exists, err := s.o.UserRepository.Exists(ctx, s.makeClause(excludedId, fields))
	if err != nil {
		return false, err
	}
	return !exists, err
}

func (s UserUniqueChecker) makeClause(excludedId string, fields []usecase.Field) entity.Clause {
	if excludedId == "" {
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

	return func(p entity.Prober) (expr entity.Expression) {
		expr = p.Field("ID").Is(entity.OperatorNotEquals, excludedId)
		for _, field := range fields {
			expr = expr.And().Field(field.Name).Is(entity.OperatorEquals, field.Value)
		}
		return
	}
}
