package repositorygorm

import (
	"github.com/ezraisw/tanshogyo/pkg/gormds/repository"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/model"
	"gorm.io/gorm"
)

type GORMProductRepository struct {
	*repository.GORMRepository[model.Product]
}

func NewGORMProductRepository(db *gorm.DB) *GORMProductRepository {
	return &GORMProductRepository{
		GORMRepository: repository.NewGORMRepository[model.Product](db),
	}
}
