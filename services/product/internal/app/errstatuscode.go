package app

import (
	"net/http"

	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/common/util/grpchelper"
	"github.com/ezraisw/tanshogyo/pkg/common/util/httphelper"
	"google.golang.org/grpc/codes"

	sellererrors "github.com/ezraisw/tanshogyo/services/product/internal/app/seller/errors"
)

func init() {
	httphelper.RegisterError(preseterrors.ErrIs(preseterrors.ErrNotFound), http.StatusNotFound)
	httphelper.RegisterError(preseterrors.ErrIs(preseterrors.ErrUnauthorized), http.StatusUnauthorized)
	httphelper.RegisterError(preseterrors.ErrIs(preseterrors.ErrForbidden), http.StatusForbidden)
	httphelper.RegisterError(preseterrors.IsValidationError, http.StatusUnprocessableEntity)

	grpchelper.RegisterError(preseterrors.ErrIs(preseterrors.ErrNotFound), codes.NotFound)
	grpchelper.RegisterError(preseterrors.ErrIs(preseterrors.ErrUnauthorized), codes.Unauthenticated)
	grpchelper.RegisterError(preseterrors.ErrIs(preseterrors.ErrForbidden), codes.PermissionDenied)
	grpchelper.RegisterError(preseterrors.IsValidationError, codes.InvalidArgument)

	// Business
	httphelper.RegisterError(preseterrors.ErrIs(sellererrors.ErrAlreadyHasSellerAccount), http.StatusUnprocessableEntity)
	grpchelper.RegisterError(preseterrors.ErrIs(sellererrors.ErrAlreadyHasSellerAccount), codes.AlreadyExists)
}
