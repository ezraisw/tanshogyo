package gormds

import "gorm.io/gorm"

func ProvideDB(connector *Connector) (*gorm.DB, error) {
	return connector.DB()
}
