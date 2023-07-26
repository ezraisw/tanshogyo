package repository

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/seller/model"
)

type SellerRepository interface {
	repository.Repository[model.Seller]
}
