package entity

import (
	"reflect"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func ColumnName[T any](tx *gorm.DB, name string) string {
	var e T
	rt := reflect.TypeOf(e)

	var tableName string
	if tabler, ok := (any)(e).(schema.Tabler); ok {
		tableName = tabler.TableName()
	} else {
		tableName = tx.NamingStrategy.TableName(rt.Name())
	}

	return tx.NamingStrategy.ColumnName(tableName, name)
}
