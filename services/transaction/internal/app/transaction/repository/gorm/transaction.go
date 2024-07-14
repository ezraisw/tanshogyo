package repositorygorm

import (
	"github.com/ezraisw/tanshogyo/pkg/gormds/repository"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/model"
	"gorm.io/gorm"
)

type GORMTransactionRepository struct {
	*repository.GORMRepository[model.Transaction]
}

func NewGORMTransactionRepository(db *gorm.DB) *GORMTransactionRepository {
	return &GORMTransactionRepository{
		GORMRepository: repository.NewGORMRepository[model.Transaction](db),
	}
}
