package rules

import validation "github.com/go-ozzo/ozzo-validation/v4"

var (
	ErrUnique = validation.NewError("validation_unique", "value is not unique")
	ErrExists = validation.NewError("validation_exists", "value does not exist")
)
