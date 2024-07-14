package app

import (
	usergrpc "github.com/ezraisw/tanshogyo/services/user/internal/app/user/adapter/grpc"
	userweb "github.com/ezraisw/tanshogyo/services/user/internal/app/user/adapter/web"
)

type HandlerRegistries struct {
	UserGRPC *usergrpc.UserHandlerRegistry
	UserHTTP *userweb.UserHandlerRegistry
}
