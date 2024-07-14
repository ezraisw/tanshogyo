package config

import (
	"github.com/ezraisw/tanshogyo/pkg/common/adapter/grpc"
	"github.com/ezraisw/tanshogyo/pkg/common/adapter/web"
	"github.com/ezraisw/tanshogyo/pkg/common/logger"
	"github.com/ezraisw/tanshogyo/pkg/gormds"
)

type Config struct {
	Stage       string
	HTTP        *web.HTTPConfig
	GRPC        *grpc.GRPCConfig
	Logger      *logger.LoggerConfig
	Datasources *DatasourcesConfig
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
