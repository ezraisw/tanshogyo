package repository

import (
	"context"

	"github.com/ezraisw/tanshogyo/pkg/common/entity"
)

//go:generate go run github.com/golang/mock/mockgen -source=repository.go -destination=mock/repository_mock.go -package=repositorymock

type Repository[T any] interface {
	Exists(context.Context, entity.Clause) (bool, error)
	Count(context.Context, entity.Clause) (int, error)
	FindMany(context.Context, entity.Clause, FindManyOptions) ([]*T, error)
	Find(context.Context, entity.Clause, FindOptions) (*T, error)
	Create(context.Context, *T) (*T, error)
	Update(context.Context, *T) error
	Delete(context.Context, entity.Clause) error
}
