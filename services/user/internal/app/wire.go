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
	"github.com/pwnedgod/tanshogyo/pkg/common/util/hasher"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/idgen"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	"github.com/pwnedgod/tanshogyo/pkg/gormds"
	gormentity "github.com/pwnedgod/tanshogyo/pkg/gormds/entity"
	userconfig "github.com/pwnedgod/tanshogyo/services/user/internal/pkg/config"

	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user"
)

var (
	ModuleSet = wire.NewSet(
		configSet,
		grpcSet,
		gormSet,
		hasherSet,
		idgenSet,
		loggerSet,
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

	hasherSet = wire.NewSet(
		hasher.ProvideHasher,
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
			EnvPrefix: "USER",
		}),
		viper.NewViperBinder,
		wire.Bind(new(config.Binder), new(*viper.ViperBinder)),
		userconfig.ProvideConfig,

		wire.Bind(new(core.StageGetter), new(*userconfig.Config)),
		wire.Bind(new(gormds.GORMConfigGetter), new(*userconfig.Config)),
		wire.Bind(new(logger.LoggerConfigGetter), new(*userconfig.Config)),
		wire.Bind(new(grpc.GRPCConfigGetter), new(*userconfig.Config)),
		wire.Bind(new(web.HTTPConfigGetter), new(*userconfig.Config)),
	)
)

func InjectApplication() (*Application, func(), error) {
	panic(wire.Build(
		ModuleSet,
		user.ModuleSet,
	))
}
