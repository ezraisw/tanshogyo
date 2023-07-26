//go:build wireinject
// +build wireinject

package seller

import (
	"github.com/google/wire"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/adapter/web"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/factory"
	factoryimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/factory/impl"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/repository"
	repositorygorm "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/repository/gorm"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/usecase/impl"
)

var (
	ModuleSet = wire.NewSet(
		adapterSet,
		factorySet,
		repositorySet,
		usecaseSet,
	)

	adapterSet = wire.NewSet(
		wire.Struct(new(web.SellerControllerOptions), "*"),
		web.NewSellerController,
		wire.Struct(new(web.SellerHandlerRegistryOptions), "*"),
		web.NewSellerHandlerRegistry,
	)

	factorySet = wire.NewSet(
		wire.Struct(new(factoryimpl.SellerExistsRuleFactoryOptions), "*"),
		factoryimpl.NewSellerExistsRuleFactory,
		wire.Bind(new(factory.SellerExistsRuleFactory), new(*factoryimpl.SellerExistsRuleFactory)),
	)

	repositorySet = wire.NewSet(
		repositorygorm.NewGORMSellerRepository,
		wire.Bind(new(repository.SellerRepository), new(*repositorygorm.GORMSellerRepository)),
	)

	usecaseSet = wire.NewSet(
		usecaseimpl.NewSellerFormValidator,
		wire.Bind(new(usecase.SellerFormValidator), new(*usecaseimpl.SellerFormValidator)),

		wire.Struct(new(usecaseimpl.SellerGetterOptions), "*"),
		usecaseimpl.NewSellerGetter,
		wire.Bind(new(usecase.SellerGetter), new(*usecaseimpl.SellerGetter)),

		wire.Struct(new(usecaseimpl.SellerRegistererOptions), "*"),
		usecaseimpl.NewSellerRegisterer,
		wire.Bind(new(usecase.SellerRegisterer), new(*usecaseimpl.SellerRegisterer)),

		wire.Struct(new(usecaseimpl.SellerExistsCheckerOptions), "*"),
		usecaseimpl.NewSellerExistsChecker,
		wire.Bind(new(usecase.SellerExistsChecker), new(*usecaseimpl.SellerExistsChecker)),
	)
)
