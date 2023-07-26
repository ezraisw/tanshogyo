//go:build wireinject
// +build wireinject

package product

import (
	"github.com/google/wire"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/adapter/grpc"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/adapter/web"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/middleware"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/repository"
	repositorygorm "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/repository/gorm"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
	usecaseimpl "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase/impl"
)

var (
	ModuleSet = wire.NewSet(
		adapterSet,
		middlewareSet,
		repositorySet,
		usecaseSet,
	)

	adapterSet = wire.NewSet(
		wire.Struct(new(grpc.ProductServiceOptions), "*"),
		grpc.NewProductService,
		wire.Struct(new(grpc.ProductHandlerRegistryOptions), "*"),
		grpc.NewProductHandlerRegistry,

		wire.Struct(new(web.ProductControllerOptions), "*"),
		web.NewProductController,
		wire.Struct(new(web.ProductHandlerRegistryOptions), "*"),
		web.NewProductHandlerRegistry,
	)

	middlewareSet = wire.NewSet(
		middleware.ProvideSellerCheckerMiddleware,
	)

	repositorySet = wire.NewSet(
		repositorygorm.NewGORMProductRepository,
		wire.Bind(new(repository.ProductRepository), new(*repositorygorm.GORMProductRepository)),
	)

	usecaseSet = wire.NewSet(
		wire.Struct(new(usecaseimpl.ProductFormValidatorOptions), "*"),
		usecaseimpl.NewProductFormValidator,
		wire.Bind(new(usecase.ProductFormValidator), new(*usecaseimpl.ProductFormValidator)),

		wire.Struct(new(usecaseimpl.ProductListerOptions), "*"),
		usecaseimpl.NewProductLister,
		wire.Bind(new(usecase.ProductLister), new(*usecaseimpl.ProductLister)),

		wire.Struct(new(usecaseimpl.ProductGetterOptions), "*"),
		usecaseimpl.NewProductGetter,
		wire.Bind(new(usecase.ProductGetter), new(*usecaseimpl.ProductGetter)),

		wire.Struct(new(usecaseimpl.ProductCreatorOptions), "*"),
		usecaseimpl.NewProductCreator,
		wire.Bind(new(usecase.ProductCreator), new(*usecaseimpl.ProductCreator)),

		wire.Struct(new(usecaseimpl.ProductUpdaterOptions), "*"),
		usecaseimpl.NewProductUpdater,
		wire.Bind(new(usecase.ProductUpdater), new(*usecaseimpl.ProductUpdater)),

		wire.Struct(new(usecaseimpl.ProductDeleterOptions), "*"),
		usecaseimpl.NewProductDeleter,
		wire.Bind(new(usecase.ProductDeleter), new(*usecaseimpl.ProductDeleter)),

		wire.Struct(new(usecaseimpl.ProductAuthedListerOptions), "*"),
		usecaseimpl.NewProductAuthedLister,
		wire.Bind(new(usecase.ProductAuthedLister), new(*usecaseimpl.ProductAuthedLister)),

		wire.Struct(new(usecaseimpl.ProductAuthedCreatorOptions), "*"),
		usecaseimpl.NewProductAuthedCreator,
		wire.Bind(new(usecase.ProductAuthedCreator), new(*usecaseimpl.ProductAuthedCreator)),

		wire.Struct(new(usecaseimpl.ProductAuthedUpdaterOptions), "*"),
		usecaseimpl.NewProductAuthedUpdater,
		wire.Bind(new(usecase.ProductAuthedUpdater), new(*usecaseimpl.ProductAuthedUpdater)),

		wire.Struct(new(usecaseimpl.ProductAuthedDeleterOptions), "*"),
		usecaseimpl.NewProductAuthedDeleter,
		wire.Bind(new(usecase.ProductAuthedDeleter), new(*usecaseimpl.ProductAuthedDeleter)),

		wire.Struct(new(usecaseimpl.ProductSellerCheckerOptions), "*"),
		usecaseimpl.NewProductSellerChecker,
		wire.Bind(new(usecase.ProductSellerChecker), new(*usecaseimpl.ProductSellerChecker)),
	)
)
