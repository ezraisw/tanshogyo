package middleware

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/httphelper"
	"github.com/pwnedgod/tanshogyo/pkg/userauth"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/usecase"
)

type SellerCheckerMiddleware func(http.Handler) http.Handler

func ProvideSellerCheckerMiddleware(sellerChecker usecase.ProductSellerChecker) SellerCheckerMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			authUser := userauth.FromContext(ctx)
			if authUser == nil {
				// Should be handled by middleware.
				httphelper.RespondError(w, r, preseterrors.ErrInternalProblem)
				return
			}

			if err := sellerChecker.CheckSeller(ctx, authUser.ID, chi.URLParam(r, "id")); err != nil {
				httphelper.RespondError(w, r, err)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
