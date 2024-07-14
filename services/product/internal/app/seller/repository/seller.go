package repository

import (
	"github.com/ezraisw/tanshogyo/pkg/common/repository"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/seller/model"
)

type SellerRepository interface {
	repository.Repository[model.Seller]
}
