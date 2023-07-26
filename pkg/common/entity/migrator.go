package entity

type Migrator interface {
	Migrate(models []any) error
}
