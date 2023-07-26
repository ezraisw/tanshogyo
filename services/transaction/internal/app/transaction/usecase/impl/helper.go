package usecaseimpl

import (
	"github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/model"
	"github.com/pwnedgod/tanshogyo/services/transaction/internal/app/transaction/usecase"
)

func toDto(transaction *model.Transaction) usecase.Transaction {
	dto := usecase.Transaction{
		ID:        transaction.ID,
		UserID:    transaction.UserID,
		Details:   make([]usecase.TransactionDetail, 0, len(transaction.Details)),
		CreatedAt: transaction.CreatedAt,
		UpdatedAt: transaction.UpdatedAt,
	}

	for _, detail := range transaction.Details {
		dto.Details = append(dto.Details, usecase.TransactionDetail{
			ID:        detail.ID,
			ProductID: detail.ProductID,
			Price:     detail.Price,
			Quantity:  detail.Quantity,
		})
		dto.TotalPrice += detail.Price * int64(detail.Quantity)
	}

	return dto
}
