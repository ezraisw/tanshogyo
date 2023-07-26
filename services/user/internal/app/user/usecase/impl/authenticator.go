package usecaseimpl

import (
	"context"
	"errors"

	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	commonrepository "github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/pkg/common/util/timehelper"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/repository"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
)

type UserAuthenticatorOptions struct {
	AuthenticationRepository repository.AuthenticationRepository
	UserRepository           repository.UserRepository
	Nower                    timehelper.Nower
}

type UserAuthenticator struct {
	o UserAuthenticatorOptions
}

func NewUserAuthenticator(options UserAuthenticatorOptions) *UserAuthenticator {
	return &UserAuthenticator{
		o: options,
	}
}

func (s UserAuthenticator) Authenticate(ctx context.Context, token string) (usecase.User, error) {
	now := s.o.Nower.Now().UTC()

	authentication, err := s.o.AuthenticationRepository.Find(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("Token").Is(entity.OperatorEquals, token).And().
			Field("ExpiredAt").Is(entity.OperatorGT, now)
	}, commonrepository.FindOptions{})
	if err != nil {
		if errors.Is(err, preseterrors.ErrNotFound) {
			return usecase.User{}, preseterrors.ErrUnauthorized
		}
		return usecase.User{}, handleRepoErrorForAuth(err)
	}

	user, err := s.o.UserRepository.Find(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("ID").Is(entity.OperatorEquals, authentication.UserID)
	}, commonrepository.FindOptions{})
	if err != nil {
		return usecase.User{}, handleRepoErrorForAuth(err)
	}

	return toDto(user), nil
}
