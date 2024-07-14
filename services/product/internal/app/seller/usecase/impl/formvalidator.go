package usecaseimpl

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/rules"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/usecase"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SellerFormValidator struct {
}

func NewSellerFormValidator() *SellerFormValidator {
	return &SellerFormValidator{}
}

func (s SellerFormValidator) Validate(ctx context.Context, form usecase.SellerForm) error {
	valErr := validation.ValidateStructWithContext(ctx, &form,
		validation.Field(&form.ShopName, validation.Required),
	)
	return rules.MapErrors(valErr)
}
