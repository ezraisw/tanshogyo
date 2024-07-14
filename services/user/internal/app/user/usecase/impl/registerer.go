package usecaseimpl

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/util/hasher"
	"github.com/ezraisw/tanshogyo/pkg/common/util/idgen"
	"github.com/ezraisw/tanshogyo/pkg/common/util/timehelper"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/model"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/repository"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/usecase"
)

type UserRegistererOptions struct {
	UserFormValidator usecase.UserFormValidator
	UserRepository    repository.UserRepository
	Nower             timehelper.Nower
	IDGen             idgen.IDGen
	Hasher            hasher.Hasher
}

type UserRegisterer struct {
	o UserRegistererOptions
}

func NewUserRegisterer(options UserRegistererOptions) *UserRegisterer {
	return &UserRegisterer{
		o: options,
	}
}

func (s UserRegisterer) Register(ctx context.Context, form usecase.UserForm) error {
	if err := s.o.UserFormValidator.Validate(ctx, form); err != nil {
		return err
	}

	now := s.o.Nower.Now().UTC()

	user := fromForm(form)
	user.ID = s.o.IDGen.Generate()
	user.CreatedAt = now
	user.UpdatedAt = now

	if err := s.hashPasswordOf(user); err != nil {
		return err
	}

	if _, err := s.o.UserRepository.Create(ctx, user); err != nil {
		return err
	}

	return nil
}

func (s UserRegisterer) hashPasswordOf(user *model.User) error {
	hashedPassword, err := s.o.Hasher.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return nil
}
