package repositorygorm

import (
	"github.com/pwnedgod/tanshogyo/pkg/gormds/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/model"
	"gorm.io/gorm"
)

type GORMSellerRepository struct {
	*repository.GORMRepository[model.Seller]
}

func NewGORMSellerRepository(db *gorm.DB) *GORMSellerRepository {
	return &GORMSellerRepository{
		GORMRepository: repository.NewGORMRepository[model.Seller](db),
	}
}
