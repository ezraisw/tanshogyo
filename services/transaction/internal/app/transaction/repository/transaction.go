package repository

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/model"
)

type TransactionRepository interface {
	repository.Repository[model.Transaction]
}
