package repository

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/user/internal/app/user/model"
)

type UserRepository interface {
	repository.Repository[model.User]
}
