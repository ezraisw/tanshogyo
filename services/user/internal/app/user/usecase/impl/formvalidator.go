package usecaseimpl

import (
	"context"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/pwnedgod/tanshogyo/pkg/common/rules"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/factory"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
)

type UserFormValidatorOptions struct {
	UserUniqueRuleFactory factory.UserUniqueRuleFactory
}

type UserFormValidator struct {
	o UserFormValidatorOptions
}

func NewUserFormValidator(options UserFormValidatorOptions) *UserFormValidator {
	return &UserFormValidator{
		o: options,
	}
}

func (s UserFormValidator) Validate(ctx context.Context, form usecase.UserForm) error {
	valErr := validation.ValidateStructWithContext(ctx, &form,
		validation.Field(&form.Username, validation.Required, validation.Length(4, 0), s.o.UserUniqueRuleFactory.Make("", "Username")),
		validation.Field(&form.Password, validation.Required, validation.Length(8, 0), rules.HasLowerCaseUpperCaseAndDigits),
		validation.Field(&form.Email, validation.Required, is.EmailFormat, s.o.UserUniqueRuleFactory.Make("", "Email")),
		validation.Field(&form.Name, validation.Required),
	)
	return rules.MapErrors(valErr)
}
