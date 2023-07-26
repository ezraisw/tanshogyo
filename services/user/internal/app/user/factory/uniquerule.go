package factory

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

//go:generate go run github.com/golang/mock/mockgen -source=uniquerule.go -destination=mock/uniquerule_mock.go -package=factorymock

type Field struct {
	Name  string
	Value any
}

type UserUniqueRuleFactory interface {
	Make(excludedId string, fieldName string, otherFields ...Field) validation.Rule
}
