package usecase

import "context"

type SellerExistsChecker interface {
	CheckExists(ctx context.Context, fields []Field) (bool, error)
}
