package repository

import (
	"github.com/ezraisw/tanshogyo/pkg/common/repository"
	"github.com/ezraisw/tanshogyo/services/product/internal/app/product/model"
)

type ProductRepository interface {
	repository.Repository[model.Product]
}
