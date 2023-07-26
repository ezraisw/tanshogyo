package userauthgrpc

type ProductConfigGetter interface {
	GetProductConfig() *ProductConfig
}

type ProductConfig struct {
	Host string
	Port int
}
