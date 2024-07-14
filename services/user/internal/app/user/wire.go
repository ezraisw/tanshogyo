//go:build wireinject
// +build wireinject

package user

import (
	"github.com/google/wire"

	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/adapter/grpc"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/adapter/web"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/factory"
	factoryimpl "github.com/ezraisw/tanshogyo/services/user/internal/app/user/factory/impl"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/repository"
	repositorygorm "github.com/ezraisw/tanshogyo/services/user/internal/app/user/repository/gorm"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase"
	usecaseimpl "github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase/impl"
)

var (
	ModuleSet = wire.NewSet(
		adapterSet,
		factorySet,
		repositorySet,
		usecaseSet,
	)

	adapterSet = wire.NewSet(
		wire.Struct(new(grpc.UserServiceOptions), "*"),
		grpc.NewUserService,
		wire.Struct(new(grpc.UserHandlerRegistryOptions), "*"),
		grpc.NewUserHandlerRegistry,

		wire.Struct(new(web.UserControllerOptions), "*"),
		web.NewUserController,
		wire.Struct(new(web.UserHandlerRegistryOptions), "*"),
		web.NewUserHandlerRegistry,
	)

	factorySet = wire.NewSet(
		wire.Struct(new(factoryimpl.UserUniqueRuleFactoryOptions), "*"),
		factoryimpl.NewUserUniqueRuleFactory,
		wire.Bind(new(factory.UserUniqueRuleFactory), new(*factoryimpl.UserUniqueRuleFactory)),
	)

	repositorySet = wire.NewSet(
		repositorygorm.NewGORMAuthenticationRepository,
		wire.Bind(new(repository.AuthenticationRepository), new(*repositorygorm.GORMAuthenticationRepository)),

		repositorygorm.NewGORMUserRepository,
		wire.Bind(new(repository.UserRepository), new(*repositorygorm.GORMUserRepository)),
	)

	usecaseSet = wire.NewSet(
		wire.Struct(new(usecaseimpl.UserUniqueCheckerOptions), "*"),
		usecaseimpl.NewUserUniqueChecker,
		wire.Bind(new(usecase.UserUniqueChecker), new(*usecaseimpl.UserUniqueChecker)),

		wire.Struct(new(usecaseimpl.UserFormValidatorOptions), "*"),
		usecaseimpl.NewUserFormValidator,
		wire.Bind(new(usecase.UserFormValidator), new(*usecaseimpl.UserFormValidator)),

		wire.Struct(new(usecaseimpl.UserAuthenticatorOptions), "*"),
		usecaseimpl.NewUserAuthenticator,
		wire.Bind(new(usecase.UserAuthenticator), new(*usecaseimpl.UserAuthenticator)),

		wire.Struct(new(usecaseimpl.UserLoginerOptions), "*"),
		usecaseimpl.NewUserLoginer,
		wire.Bind(new(usecase.UserLoginer), new(*usecaseimpl.UserLoginer)),

		wire.Struct(new(usecaseimpl.UserRegistererOptions), "*"),
		usecaseimpl.NewUserRegisterer,
		wire.Bind(new(usecase.UserRegisterer), new(*usecaseimpl.UserRegisterer)),
	)
)
