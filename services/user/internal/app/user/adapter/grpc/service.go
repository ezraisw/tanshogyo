package grpc

import (
	context "context"

	"github.com/pwnedgod/tanshogyo/pkg/common/util/grpchelper"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/adapter/grpc/pb"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServiceOptions struct {
	UserAuthenticator usecase.UserAuthenticator
}

type UserService struct {
	pb.UnimplementedUserServiceServer

	o UserServiceOptions
}

func NewUserService(options UserServiceOptions) *UserService {
	return &UserService{
		o: options,
	}
}

func (s UserService) Authenticate(ctx context.Context, req *pb.AuthenticationRequest) (*pb.User, error) {
	dto, err := s.o.UserAuthenticator.Authenticate(ctx, req.Token)
	if err != nil {
		return nil, status.Error(grpchelper.GetCode(err), err.Error())
	}

	msg := &pb.User{
		Id:        dto.ID,
		Username:  dto.Username,
		Email:     dto.Email,
		Name:      dto.Name,
		CreatedAt: timestamppb.New(dto.CreatedAt),
		UpdatedAt: timestamppb.New(dto.UpdatedAt),
	}
	return msg, nil
}
