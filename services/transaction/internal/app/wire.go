//go:build wireinject
// +build wireinject

package app

import (
	"io"
	"os"

	"github.com/google/wire"
	"github.com/pwnedgod/tanshogyo/pkg/common/adapter/grpc"
	"github.com/pwnedgod/tanshogyo/pkg/common/adapter/web"
	"github.com/pwnedgod/tanshogyo/pkg/common/config"
	"github.com/pwnedgod/tanshogyo/pkg/common/config/viper"
	"github.com/pwnedgod/tanshogyo/pkg/common/core"
	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	"github.com/pwnedgod/tanshogyo/pkg/common/logger"
	"github.com/pwnedgod/tanshogyo/pkg/common/redis"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/idgen"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	"github.com/pwnedgod/tanshogyo/pkg/gormds"
	gormentity "github.com/pwnedgod/tanshogyo/pkg/gormds/entity"
	"github.com/pwnedgod/tanshogyo/pkg/product"
	productgrpc "github.com/pwnedgod/tanshogyo/pkg/product/grpc"
	"github.com/pwnedgod/tanshogyo/pkg/userauth"
	userauthgrpc "github.com/pwnedgod/tanshogyo/pkg/userauth/grpc"
	transactionconfig "github.com/pwnedgod/tanshogyo/services/transaction/internal/pkg/config"

	"github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction"
)

var (
	ModuleSet = wire.NewSet(
		configSet,
		grpcSet,
		gormSet,
		idgenSet,
		loggerSet,
		productSet,
		redisSet,
		timehelperSet,
		userauthSet,
		webSet,
		mainSet,
	)

	mainSet = wire.NewSet(
		wire.Struct(new(HandlerRegistries), "*"),
		wire.Struct(new(MiddlewareRegistries), "*"),
		wire.Struct(new(Runners), "*"),

		wire.Struct(new(ApplicationOptions), "*"),
		NewApplication,
	)
)

// PKG bindings.
var (
	grpcSet = wire.NewSet(
		wire.Struct(new(grpc.GRPCRunnerOptions), "*"),
		grpc.NewGRPCRunner,
	)

	gormSet = wire.NewSet(
		gormds.ProvideDB,
		gormds.NewConnector,

		wire.Struct(new(gormentity.GORMMigratorOptions), "*"),
		gormentity.NewGORMMigrator,
		wire.Bind(new(entity.Migrator), new(*gormentity.GORMMigrator)),
	)

	idgenSet = wire.NewSet(
		idgen.ProvideIDGen,
	)

	loggerSet = wire.NewSet(
		wire.InterfaceValue(new(io.Writer), os.Stdout),
		logger.ProvideLogger,
		wire.Struct(new(logger.LoggerMiddlewareRegistryOptions), "*"),
		logger.NewLoggerMiddlewareRegistry,
	)

	productSet = wire.NewSet(
		productgrpc.NewProductGRPCClient,
		productgrpc.NewGRPCProductAPI,
		wire.Bind(new(product.ProductAPI), new(*productgrpc.GRPCProductAPI)),
	)

	redisSet = wire.NewSet(
		redis.ProvideUniversalClient,
	)

	timehelperSet = wire.NewSet(
		timehelper.ProvideNower,
	)

	userauthSet = wire.NewSet(
		userauth.ProvideUserAuthMiddleware,
		userauthgrpc.NewUserGRPCClient,
		userauthgrpc.NewGRPCUserAPI,
		wire.Bind(new(userauth.UserAPI), new(*userauthgrpc.GRPCUserAPI)),
	)

	webSet = wire.NewSet(
		wire.Struct(new(web.WebRunnerOptions), "*"),
		web.NewWebRunner,
		web.NewEssentialsMiddlewareRegistry,
	)
)

// Configuration bindings.
var (
	configSet = wire.NewSet(
		wire.Value(config.BinderProperties{
			Paths:     []string{"./configs"},
			FileName:  "app-config",
			EnvPrefix: "TRANSACTION",
		}),
		viper.NewViperBinder,
		wire.Bind(new(config.Binder), new(*viper.ViperBinder)),
		transactionconfig.ProvideConfig,

		wire.Bind(new(core.StageGetter), new(*transactionconfig.Config)),
		wire.Bind(new(gormds.GORMConfigGetter), new(*transactionconfig.Config)),
		wire.Bind(new(logger.LoggerConfigGetter), new(*transactionconfig.Config)),
		wire.Bind(new(grpc.GRPCConfigGetter), new(*transactionconfig.Config)),
		wire.Bind(new(web.HTTPConfigGetter), new(*transactionconfig.Config)),
		wire.Bind(new(redis.RedisConfigGetter), new(*transactionconfig.Config)),
		wire.Bind(new(userauthgrpc.UserAuthConfigGetter), new(*transactionconfig.Config)),
		wire.Bind(new(productgrpc.ProductConfigGetter), new(*transactionconfig.Config)),
	)
)

func InjectApplication() (*Application, func(), error) {
	panic(wire.Build(
		ModuleSet,
		transaction.ModuleSet,
	))
}
