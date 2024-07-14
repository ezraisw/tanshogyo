package grpchelper

import (
	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func HandleError(err error) error {
	if s, ok := status.FromError(err); ok {
		switch s.Code() {
		case codes.InvalidArgument:
			return preseterrors.ErrInvalidArguments
		case codes.NotFound:
			return preseterrors.ErrNotFound
		case codes.PermissionDenied:
			return preseterrors.ErrForbidden
		case codes.Unauthenticated:
			return preseterrors.ErrUnauthorized
		}
	}
	return err
}
