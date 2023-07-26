package userauth

import (
	"context"
	"net/http"

	"github.com/pwnedgod/tanshogyo/pkg/common/util/httphelper"
)

//go:generate go run github.com/golang/mock/mockgen -source=middleware.go -destination=mock/middleware_mock.go -package=userauthmock

type UserAuthMiddleware func(http.Handler) http.Handler

func ProvideUserAuthMiddleware(userApi UserAPI) UserAuthMiddleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token := r.Header.Get("X-Auth-Token")

			user, err := userApi.Authenticate(ctx, token)
			if err != nil {
				httphelper.RespondError(w, r, err)
				return
			}

			newCtx := context.WithValue(ctx, contextKeyUser, &user)
			r = r.WithContext(newCtx)
			next.ServeHTTP(w, r)
		})
	}
}
