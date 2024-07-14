package repository

import (
	"github.com/ezraisw/tanshogyo/pkg/common/repository"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/model"
)

type AuthenticationRepository interface {
	repository.Repository[model.Authentication]
}
