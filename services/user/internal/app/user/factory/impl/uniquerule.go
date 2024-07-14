package factoryimpl

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/rules"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/factory"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserUniqueRuleFactoryOptions struct {
	UserUniqueChecker usecase.UserUniqueChecker
}

type UserUniqueRuleFactory struct {
	o UserUniqueRuleFactoryOptions
}

func NewUserUniqueRuleFactory(options UserUniqueRuleFactoryOptions) *UserUniqueRuleFactory {
	return &UserUniqueRuleFactory{
		o: options,
	}
}

func (f UserUniqueRuleFactory) Make(excludedId string, fieldName string, otherFields ...factory.Field) validation.Rule {
	return validation.WithContext(func(ctx context.Context, value any) error {
		fields := []usecase.Field{
			{Name: fieldName, Value: value},
		}

		for _, otherField := range otherFields {
			fields = append(fields, usecase.Field{Name: otherField.Name, Value: otherField.Value})
		}

		unique, err := f.o.UserUniqueChecker.CheckUnique(ctx, excludedId, fields)
		if err != nil {
			return validation.NewInternalError(err)
		}

		if !unique {
			return rules.ErrUnique
		}

		return nil
	})
}
