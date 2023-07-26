package entity

import (
	"reflect"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type GORMMigratorOptions struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

type GORMMigrator struct {
	o GORMMigratorOptions
}

func NewGORMMigrator(options GORMMigratorOptions) *GORMMigrator {
	return &GORMMigrator{
		o: options,
	}
}

func (m GORMMigrator) Migrate(models []any) error {
	for _, model := range models {
		rt := reflect.Indirect(reflect.ValueOf(model)).Type()

		m.o.Logger.Debug().
			Str("model", rt.Name()).
			Msg("migrating model")

		if err := m.o.DB.Migrator().AutoMigrate(model); err != nil {
			return err
		}
	}

	return nil
}
