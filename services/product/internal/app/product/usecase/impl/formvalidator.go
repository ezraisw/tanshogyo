package usecaseimpl

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pwnedgod/tanshogyo/pkg/common/rules"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/factory"
)

type ProductFormValidatorOptions struct {
	SellerExistsRuleFactory factory.SellerExistsRuleFactory
}

type ProductFormValidator struct {
	o ProductFormValidatorOptions
}

func NewProductFormValidator(options ProductFormValidatorOptions) *ProductFormValidator {
	return &ProductFormValidator{
		o: options,
	}
}

func (v ProductFormValidator) Validate(ctx context.Context, form usecase.ProductForm) error {
	valErr := validation.ValidateStructWithContext(ctx, &form,
		validation.Field(&form.SellerID, validation.Required, v.o.SellerExistsRuleFactory.Make("ID")),
		validation.Field(&form.Name, validation.Required),
		validation.Field(&form.Description, validation.Required),
		validation.Field(&form.Price, validation.Required, validation.Min(0).Exclusive()),
		validation.Field(&form.Quantity, validation.Min(0)),
	)
	return rules.MapErrors(valErr)
}
