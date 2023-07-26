package repositorygorm

import (
	"github.com/pwnedgod/tanshogyo/pkg/gormds/repository"
	"github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/model"
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
