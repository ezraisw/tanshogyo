package config

import (
	"github.com/ezraisw/tanshogyo/pkg/common/adapter/grpc"
	"github.com/ezraisw/tanshogyo/pkg/common/adapter/web"
	"github.com/ezraisw/tanshogyo/pkg/common/logger"
	"github.com/ezraisw/tanshogyo/pkg/common/redis"
	"github.com/ezraisw/tanshogyo/pkg/gormds"
	productgrpc "github.com/ezraisw/tanshogyo/pkg/product/grpc"
	userauthgrpc "github.com/ezraisw/tanshogyo/pkg/userauth/grpc"
)

type Config struct {
	Stage       string
	HTTP        *web.HTTPConfig
	GRPC        *grpc.GRPCConfig
	Logger      *logger.LoggerConfig
	Datasources *DatasourcesConfig
	UserAuth    *userauthgrpc.UserAuthConfig
	Product     *productgrpc.ProductConfig
	Redis       *redis.RedisConfig
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

func (c Config) GetProductConfig() *productgrpc.ProductConfig {
	return c.Product
}

func (c Config) GetRedisConfig() *redis.RedisConfig {
	return c.Redis
}
