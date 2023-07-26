package userauthgrpc

type UserAuthConfigGetter interface {
	GetUserAuthConfig() *UserAuthConfig
}

type UserAuthConfig struct {
	Host string
	Port int
}
