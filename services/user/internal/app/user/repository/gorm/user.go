package repositorygorm

import (
	"github.com/ezraisw/tanshogyo/pkg/gormds/repository"
	"github.com/ezraisw/tanshogyo/services/user/internal/app/user/model"
	"gorm.io/gorm"
)

type GORMUserRepository struct {
	*repository.GORMRepository[model.User]
}

func NewGORMUserRepository(db *gorm.DB) *GORMUserRepository {
	return &GORMUserRepository{
		GORMRepository: repository.NewGORMRepository[model.User](db),
	}
}
