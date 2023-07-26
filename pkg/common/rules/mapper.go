package rules

import (
	"fmt"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
)

func MapErrors(err error) error {
	switch v := err.(type) {
	case validation.InternalError:
		return v.InternalError()
	case validation.Errors:
		return &preseterrors.ValidationError{
			Message:     "encountered a validation error",
			FieldErrors: mapValidationErrors("", v),
		}
	}
	return err
}

func mapValidationErrors(parentField string, valErrs validation.Errors) []*preseterrors.FieldError {
	var fieldErrs []*preseterrors.FieldError
	for field, e := range valErrs {
		targetField := fmt.Sprintf("%s.%s", parentField, field)
		if strings.TrimSpace(parentField) == "" {
			targetField = field
		}
		switch v := e.(type) {
		case validation.Error:
			fieldErrs = append(fieldErrs, ToFieldError(targetField, v))
		case validation.Errors:
			childFieldErrs := mapValidationErrors(targetField, v)
			fieldErrs = append(fieldErrs, childFieldErrs...)
		}
	}
	return fieldErrs
}

func ToFieldError(field string, ozzoErr validation.Error) *preseterrors.FieldError {
	return &preseterrors.FieldError{
		Field:   field,
		Code:    ozzoErr.Code(),
		Message: ozzoErr.Error(),
	}
}

func ToFieldErrorWithParams(field string, ozzoErr validation.Error, params map[string]any) *preseterrors.FieldError {
	ozzoErr = ozzoErr.SetParams(params)

	return &preseterrors.FieldError{
		Field:   field,
		Code:    ozzoErr.Code(),
		Message: ozzoErr.Error(),
	}
}
