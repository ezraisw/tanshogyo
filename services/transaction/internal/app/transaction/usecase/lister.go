package usecase

import "context"

type TransactionLister interface {
	List(ctx context.Context, userId string, limit, offset int) (TransactionList, error)
}
