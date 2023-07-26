package userauthgrpc

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserGRPCClient struct {
	addr string
}

func NewUserGRPCClient(configGetter UserAuthConfigGetter) *UserGRPCClient {
	config := configGetter.GetUserAuthConfig()
	return &UserGRPCClient{
		addr: fmt.Sprintf("%s:%d", config.Host, config.Port),
	}
}

func (c UserGRPCClient) Dial() (*grpc.ClientConn, error) {
	return grpc.Dial(c.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
}
