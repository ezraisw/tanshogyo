//go:build wireinject
// +build wireinject

package transaction

import (
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/adapter/web"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/cache"
	cacheredis "github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/cache/redis"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/repository"
	repositorygorm "github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/repository/gorm"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/usecase/impl"
	"github.com/google/wire"
)

var (
	ModuleSet = wire.NewSet(
		adapterSet,
		cacheSet,
		repositorySet,
		usecaseSet,
	)

	adapterSet = wire.NewSet(
		wire.Struct(new(web.TransactionControllerOptions), "*"),
		web.NewTransactionController,
		wire.Struct(new(web.TransactionHandlerRegistryOptions), "*"),
		web.NewTransactionHandlerRegistry,
	)

	cacheSet = wire.NewSet(
		wire.Struct(new(cacheredis.RedisCartCacheOptions), "*"),
		cacheredis.NewRedisCartCache,
		wire.Bind(new(cache.CartCache), new(*cacheredis.RedisCartCache)),
	)

	repositorySet = wire.NewSet(
		repositorygorm.NewGORMTransactionRepository,
		wire.Bind(new(repository.TransactionRepository), new(*repositorygorm.GORMTransactionRepository)),
	)

	usecaseSet = wire.NewSet(
		wire.Struct(new(usecaseimpl.TransactionListerOptions), "*"),
		usecaseimpl.NewTransactionLister,
		wire.Bind(new(usecase.TransactionLister), new(*usecaseimpl.TransactionLister)),

		wire.Struct(new(usecaseimpl.TransactionCreatorOptions), "*"),
		usecaseimpl.NewTransactionCreator,
		wire.Bind(new(usecase.TransactionCreator), new(*usecaseimpl.TransactionCreator)),

		wire.Struct(new(usecaseimpl.TransactionCartGetterOptions), "*"),
		usecaseimpl.NewTransactionCartGetter,
		wire.Bind(new(usecase.TransactionCartGetter), new(*usecaseimpl.TransactionCartGetter)),

		wire.Struct(new(usecaseimpl.TransactionCartUpdaterOptions), "*"),
		usecaseimpl.NewTransactionCartUpdater,
		wire.Bind(new(usecase.TransactionCartUpdater), new(*usecaseimpl.TransactionCartUpdater)),
	)
)
