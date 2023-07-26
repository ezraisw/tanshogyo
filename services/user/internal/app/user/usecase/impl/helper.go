package usecaseimpl

import (
	"errors"

	"github.com/pwnedgod/tanshogyo/pkg/common/preseterrors"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/model"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/usecase"
)

func fromForm(form usecase.UserForm) *model.User {
	return &model.User{
		Username: form.Username,
		Password: form.Password,
		Email:    form.Email,
		Name:     form.Name,
	}
}

func toDto(user *model.User) usecase.User {
	return usecase.User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func handleRepoErrorForAuth(err error) error {
	if errors.Is(err, preseterrors.ErrNotFound) {
		return preseterrors.ErrUnauthorized
	}
	return err
}
