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
	"github.com/pwnedgod/tanshogyo/pkg/common/util/idgen"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	"github.com/pwnedgod/tanshogyo/pkg/gormds"
	gormentity "github.com/pwnedgod/tanshogyo/pkg/gormds/entity"
	"github.com/pwnedgod/tanshogyo/pkg/userauth"
	userauthgrpc "github.com/pwnedgod/tanshogyo/pkg/userauth/grpc"
	productconfig "github.com/pwnedgod/tanshogyo/services/product/internal/pkg/config"

	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller"
)

var (
	ModuleSet = wire.NewSet(
		configSet,
		grpcSet,
		gormSet,
		idgenSet,
		loggerSet,
		userauthSet,
		timehelperSet,
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
			EnvPrefix: "PRODUCT",
		}),
		viper.NewViperBinder,
		wire.Bind(new(config.Binder), new(*viper.ViperBinder)),
		productconfig.ProvideConfig,

		wire.Bind(new(core.StageGetter), new(*productconfig.Config)),
		wire.Bind(new(gormds.GORMConfigGetter), new(*productconfig.Config)),
		wire.Bind(new(logger.LoggerConfigGetter), new(*productconfig.Config)),
		wire.Bind(new(grpc.GRPCConfigGetter), new(*productconfig.Config)),
		wire.Bind(new(web.HTTPConfigGetter), new(*productconfig.Config)),
		wire.Bind(new(userauthgrpc.UserAuthConfigGetter), new(*productconfig.Config)),
	)
)

func InjectApplication() (*Application, func(), error) {
	panic(wire.Build(
		ModuleSet,
		product.ModuleSet,
		seller.ModuleSet,
	))
}
