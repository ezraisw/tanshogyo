package factory

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

//go:generate go run github.com/golang/mock/mockgen -source=existsrule.go -destination=mock/existsrule_mock.go -package=factorymock

type Field struct {
	Name  string
	Value any
}

type SellerExistsRuleFactory interface {
	Make(fieldName string, otherFields ...Field) validation.Rule
}
