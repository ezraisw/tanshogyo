package usecase

import "context"

type ProductQuantityUpdater interface {
	UpdateQuantity(ctx context.Context, id string, quantity int) error
}
