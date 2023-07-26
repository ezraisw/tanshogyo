package app

import (
	productgrpc "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/adapter/grpc"
	productweb "github.com/pwnedgod/tanshogyo/services/product/internal/app/product/adapter/web"
	sellerweb "github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/adapter/web"
)

type HandlerRegistries struct {
	ProductGRPC *productgrpc.ProductHandlerRegistry
	ProductHTTP *productweb.ProductHandlerRegistry
	SellerHTTP  *sellerweb.SellerHandlerRegistry
}
