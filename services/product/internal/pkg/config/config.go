package config

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/adapter/grpc"
	"github.com/pwnedgod/tanshogyo/pkg/common/adapter/web"
	"github.com/pwnedgod/tanshogyo/pkg/common/logger"
	"github.com/pwnedgod/tanshogyo/pkg/gormds"
	userauthgrpc "github.com/pwnedgod/tanshogyo/pkg/userauth/grpc"
)

type Config struct {
	Stage       string
	HTTP        *web.HTTPConfig
	GRPC        *grpc.GRPCConfig
	Logger      *logger.LoggerConfig
	Datasources *DatasourcesConfig
	UserAuth    *userauthgrpc.UserAuthConfig
}

type DatasourcesConfig struct {
	GORM *gormds.GORMConfig
}

func (c Config) GetStage() string {
	return c.Stage
}

func (c Config) GetHTTPConfig() *web.HTTPConfig {
	return c.HTTP
}

func (c Config) GetGRPCConfig() *grpc.GRPCConfig {
	return c.GRPC
}

func (c Config) GetLoggerConfig() *logger.LoggerConfig {
	return c.Logger
}

func (c Config) GetGORMConfig() *gormds.GORMConfig {
	return c.Datasources.GORM
}

func (c Config) GetUserAuthConfig() *userauthgrpc.UserAuthConfig {
	return c.UserAuth
}
