package userauthgrpc

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ProductGRPCClient struct {
	addr string
}

func NewProductGRPCClient(configGetter ProductConfigGetter) *ProductGRPCClient {
	config := configGetter.GetProductConfig()
	return &ProductGRPCClient{
		addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
	}
}

func (c ProductGRPCClient) Dial() (*grpc.ClientConn, error) {
	return grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
