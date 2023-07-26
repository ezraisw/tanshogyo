package factoryimpl

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pwnedgod/tanshogyo/pkg/common/rules"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/factory"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
)

type SellerExistsRuleFactoryOptions struct {
	SellerExistsChecker usecase.SellerExistsChecker
}

type SellerExistsRuleFactory struct {
	o SellerExistsRuleFactoryOptions
}

func NewSellerExistsRuleFactory(options SellerExistsRuleFactoryOptions) *SellerExistsRuleFactory {
	return &SellerExistsRuleFactory{
		o: options,
	}
}

func (f SellerExistsRuleFactory) Make(fieldName string, otherFields ...factory.Field) validation.Rule {
	return validation.WithContext(func(ctx context.Context, value any) error {
		fields := []usecase.Field{
			{Name: fieldName, Value: value},
		}

		for _, otherField := range otherFields {
			fields = append(fields, usecase.Field{Name: otherField.Name, Value: otherField.Value})
		}

		exists, err := f.o.SellerExistsChecker.CheckExists(ctx, fields)
		if err != nil {
			return validation.NewInternalError(err)
		}

		if !exists {
			return rules.ErrExists
		}

		return nil
	})
}
