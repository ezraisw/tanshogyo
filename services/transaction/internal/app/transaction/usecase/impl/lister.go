package usecaseimpl

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/entity"
	commonrepository "github.com/ezraisw/tanshogyo/pkg/common/repository"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/repository"
	"github.com/ezraisw/tanshogyo/services/transaction/internal/app/transaction/usecase"
)

type TransactionListerOptions struct {
	TransactionRepository repository.TransactionRepository
}

type TransactionLister struct {
	o TransactionListerOptions
}

func NewTransactionLister(options TransactionListerOptions) *TransactionLister {
	return &TransactionLister{
		o: options,
	}
}

func (s TransactionLister) List(ctx context.Context, userId string, limit, offset int) (usecase.TransactionList, error) {
	clause := func(p entity.Prober) entity.Expression {
		return p.Field("UserID").Is(entity.OperatorEquals, userId)
	}

	count, err := s.o.TransactionRepository.Count(ctx, clause)
	if err != nil {
		return usecase.TransactionList{}, err
	}

	// Revert to defaults when invalid values are given.
	if limit <= 0 {
		limit = LimitDefault
	}

	if offset < 0 {
		offset = OffsetDefault
	}

	// Do not return anything if offset exceeds the count.
	// Reduces the amount of database connections.
	if count <= offset {
		return usecase.TransactionList{
			Count:  count,
			Limit:  limit,
			Offset: offset,
			Data:   []usecase.Transaction{},
		}, nil
	}

	products, err := s.o.TransactionRepository.FindMany(ctx, clause, commonrepository.FindManyOptions{
		Orderings: []commonrepository.Ordering{{By: "CreatedAt", Desc: true}},
		Limit:     limit,
		Offset:    offset,
		Relations: []string{"Details"},
	})
	if err != nil {
		return usecase.TransactionList{}, err
	}

	dtos := make([]usecase.Transaction, 0, len(products))
	for _, product := range products {
		dtos = append(dtos, toDto(product))
	}

	return usecase.TransactionList{
		Count:  count,
		Limit:  limit,
		Offset: offset,
		Data:   dtos,
	}, nil
}
