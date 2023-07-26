package app

import (
	"net/http"

	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/grpchelper"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/httphelper"
	transactionerrors "github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/errors"
	"google.golang.org/grpc/codes"
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
	httphelper.RegisterError(preseterrors.ErrIs(transactionerrors.ErrEmptyCart), http.StatusUnprocessableEntity)
	httphelper.RegisterError(preseterrors.ErrIs(transactionerrors.ErrInvalidCart), http.StatusUnprocessableEntity)
	grpchelper.RegisterError(preseterrors.ErrIs(transactionerrors.ErrEmptyCart), codes.InvalidArgument)
	grpchelper.RegisterError(preseterrors.ErrIs(transactionerrors.ErrInvalidCart), codes.InvalidArgument)
}
