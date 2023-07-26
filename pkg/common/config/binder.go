package config

type Binder interface {
	BindTo(i any) error
}
