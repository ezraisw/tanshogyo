package usecaseimpl

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/entity"
	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	commonrepository "github.com/ezraisw/tanshogyo/pkg/common/repository"
	"github.com/ezraisw/tanshogyo/pkg/common/util/hasher"
	"github.com/ezraisw/tanshogyo/pkg/common/util/idgen"
	"github.com/ezraisw/tanshogyo/pkg/common/util/timehelper"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/model"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/repository"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase"
)

type UserLoginerOptions struct {
	UserRepository           repository.UserRepository
	AuthenticationRepository repository.AuthenticationRepository
	Nower                    timehelper.Nower
	IDGen                    idgen.IDGen
	Hasher                   hasher.Hasher
}

type UserLoginer struct {
	o UserLoginerOptions
}

func NewUserLoginer(options UserLoginerOptions) *UserLoginer {
	return &UserLoginer{
		o: options,
	}
}

func (s UserLoginer) Login(ctx context.Context, form usecase.LoginForm) (usecase.AuthenticationResult, error) {
	user, err := s.o.UserRepository.Find(ctx, func(p entity.Prober) entity.Expression {
		return p.Field("Username").Is(entity.OperatorEquals, form.Username)
	}, commonrepository.FindOptions{})
	if err != nil {
		return usecase.AuthenticationResult{}, handleRepoErrorForAuth(err)
	}

	matched, err := s.o.Hasher.Compare(form.Password, user.Password)
	if err != nil {
		return usecase.AuthenticationResult{}, err
	}

	if !matched {
		return usecase.AuthenticationResult{}, preseterrors.ErrUnauthorized
	}

	token, err := s.createToken(ctx, user.ID)
	if err != nil {
		return usecase.AuthenticationResult{}, err
	}

	return usecase.AuthenticationResult{Token: token}, nil
}

func (s UserLoginer) createToken(ctx context.Context, userId string) (string, error) {
	now := s.o.Nower.Now().UTC()

	authentication := &model.Authentication{
		Token:     s.o.IDGen.Generate(),
		UserID:    userId,
		CreatedAt: now,
		UpdatedAt: now,
		ExpiredAt: now.Add(DurationAuthenticationExpired),
	}

	if _, err := s.o.AuthenticationRepository.Create(ctx, authentication); err != nil {
		return "", err
	}

	return authentication.Token, nil
}
