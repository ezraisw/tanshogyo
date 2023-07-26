package reflhelper

import (
	"errors"
	"reflect"
)

type (
	StructCatalog struct {
		rvStruct reflect.Value
	}

	Collected[T any] struct {
		Name  string
		Value T
	}
)

var (
	ErrIncompatibleType = errors.New("incompatible type")
)

func NewStructCatalog(i any) *StructCatalog {
	rv := reflect.ValueOf(i)
	if rv.Kind() != reflect.Struct {
		panic("not a struct")
	}
	return &StructCatalog{rvStruct: rv}
}

func (c StructCatalog) Count() int {
	return c.rvStruct.NumField()
}

func (c StructCatalog) ForEach(cb func(reflect.Value) error) error {
	for i := 0; i < c.rvStruct.NumField(); i++ {
		if err := cb(c.rvStruct.Field(i)); err != nil {
			return err
		}
	}
	return nil
}

// Collect all field values of type T in a slice.
func Collect[T any](c *StructCatalog) []Collected[T] {
	els := make([]Collected[T], 0)
	c.ForEach(func(v reflect.Value) error {
		if el, ok := v.Interface().(T); ok {
			els = append(els, Collected[T]{
				Name:  reflect.Indirect(v).Type().Name(),
				Value: el,
			})
		}
		return nil
	})
	return els
}
