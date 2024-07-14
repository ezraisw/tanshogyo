package grpc

import (
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/adapter/grpc/pb"
	"google.golang.org/grpc"
)

type UserHandlerRegistryOptions struct {
	UserService *UserService
}

type UserHandlerRegistry struct {
	o UserHandlerRegistryOptions
}

func NewUserHandlerRegistry(options UserHandlerRegistryOptions) *UserHandlerRegistry {
	return &UserHandlerRegistry{
		o: options,
	}
}

func (h UserHandlerRegistry) RegisterServices(r grpc.ServiceRegistrar) error {
	pb.RegisterUserServiceServer(r, h.o.UserService)
	return nil
}
