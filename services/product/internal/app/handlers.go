package app

import (
	productgrpc "github.com/ezraisw/tanshogyo/services/product/internal/app/product/adapter/grpc"
	productweb "github.com/ezraisw/tanshogyo/services/product/internal/app/product/adapter/web"
	sellerweb "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/adapter/web"
)

type HandlerRegistries struct {
	ProductGRPC *productgrpc.ProductHandlerRegistry
	ProductHTTP *productweb.ProductHandlerRegistry
	SellerHTTP  *sellerweb.SellerHandlerRegistry
}
