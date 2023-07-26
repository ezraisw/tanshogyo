package repository

import (
	"github.com/pwnedgod/tanshogyo/pkg/common/repository"
	"github.com/pwnedgod/tanshogyo/services/product/internal/app/product/model"
)

type ProductRepository interface {
	repository.Repository[model.Product]
}
