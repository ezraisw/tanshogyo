package repository

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/model"
)

type AuthenticationRepository interface {
	repository.Repository[model.Authentication]
}
