package grpc

type GRPCConfigGetter interface {
	GetGRPCConfig() *GRPCConfig
}

type GRPCConfig struct {
	Host            string
	Port            int
	CertificateFile string
	PrivateKeyFile  string
}
