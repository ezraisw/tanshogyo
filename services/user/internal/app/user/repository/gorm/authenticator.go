package repositorygorm

import (
	"github.com/ezraisw/tanshogyo/pkg/gormds/repository"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/model"
	"gorm.io/gorm"
)

type GORMAuthenticationRepository struct {
	*repository.GORMRepository[model.Authentication]
}

func NewGORMAuthenticationRepository(db *gorm.DB) *GORMAuthenticationRepository {
	return &GORMAuthenticationRepository{
		GORMRepository: repository.NewGORMRepository[model.Authentication](db),
	}
}
