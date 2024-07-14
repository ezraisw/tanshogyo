package repository

import (
	"github.com/ezraisw/tanshogyo/pkg/common/repository"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/model"
)

type TransactionRepository interface {
	repository.Repository[model.Transaction]
}
