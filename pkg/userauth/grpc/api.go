package userauthgrpc

import (
	"context"

	"github.com/pwnedgod/tanshogyo/pkg/common/util/grpchelper"
	"github.com/pwnedgod/tanshogyo/pkg/userauth"
	"github.com/pwnedgod/tanshogyo/pkg/userauth/grpc/pb"
)

type GRPCUserAPI struct {
	client *UserGRPCClient
}

func NewGRPCUserAPI(client *UserGRPCClient) *GRPCUserAPI {
	return &GRPCUserAPI{client: client}
}

func (a GRPCUserAPI) Authenticate(ctx context.Context, token string) (userauth.User, error) {
	conn, err := a.client.Dial()
	if err != nil {
		return userauth.User{}, err
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)
	user, err := client.Authenticate(ctx, &pb.AuthenticationRequest{
		Token: token,
	})
	if err != nil {
		return userauth.User{}, grpchelper.HandleError(err)
	}

	return fromPb(user), nil
}

func fromPb(user *pb.User) userauth.User {
	return userauth.User{
		ID:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.AsTime(),
		UpdatedAt: user.UpdatedAt.AsTime(),
	}
}
