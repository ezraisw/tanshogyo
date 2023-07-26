package repositorygorm

import (
	"github.com/pwnedgod/tanshogyo/pkg/gormds/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
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
