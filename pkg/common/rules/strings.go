package rules

import validation "github.com/go-ozzo/ozzo-validation/v4"

var (
	ErrLowerCaseUpperCaseAndDigits = validation.NewError("validation_lower_upper_digits", "must have lower case, upper case, and digits")

	HasLowerCaseUpperCaseAndDigits = validation.NewStringRuleWithError(lowerCaseUpperCaseAndDigitsValidator, ErrLowerCaseUpperCaseAndDigits)
)

func lowerCaseUpperCaseAndDigitsValidator(str string) bool {
	digits := false
	lowerCase := false
	upperCase := false

	for _, c := range str {
		if c >= '0' && c <= '9' {
			digits = true
		}

		if c >= 'a' && c <= 'z' {
			lowerCase = true
		}

		if c >= 'A' && c <= 'Z' {
			upperCase = true
		}
	}

	return digits && lowerCase && upperCase
}
