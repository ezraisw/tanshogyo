package grpc

import (
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/adapter/grpc/pb"
	"google.golang.org/grpc"
)

type ProductHandlerRegistryOptions struct {
	ProductService *ProductService
}

type ProductHandlerRegistry struct {
	o ProductHandlerRegistryOptions
}

func NewProductHandlerRegistry(options ProductHandlerRegistryOptions) *ProductHandlerRegistry {
	return &ProductHandlerRegistry{
		o: options,
	}
}

func (h ProductHandlerRegistry) RegisterServices(r grpc.ServiceRegistrar) error {
	pb.RegisterProductServiceServer(r, h.o.ProductService)
	return nil
}
