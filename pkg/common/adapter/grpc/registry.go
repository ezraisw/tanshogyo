package grpc

import "google.golang.org/grpc"

type GRPCHandlerRegistry interface {
	RegisterServices(registrar grpc.ServiceRegistrar) error
}
