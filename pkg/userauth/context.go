package userauth

import "context"

type contextKey string

const contextKeyUser contextKey = "user"

func FromContext(ctx context.Context) *User {
	if user, ok := ctx.Value(contextKeyUser).(*User); ok {
		return user
	}
	return nil
}
