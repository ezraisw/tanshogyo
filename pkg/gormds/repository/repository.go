package repository

import (
	"context"
	"errors"

	"github.com/ezraisw/tanshogyo/pkg/common/entity"
	"github.com/ezraisw/tanshogyo/pkg/common/preseterrors"
	"github.com/ezraisw/tanshogyo/pkg/common/repository"
	gormentity "github.com/ezraisw/tanshogyo/pkg/gormds/entity"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GORMRepository[T any] struct {
	db *gorm.DB
}

// Compile time implementation test.
var _ repository.Repository[struct{}] = (*GORMRepository[struct{}])(nil)

func NewGORMRepository[T any](db *gorm.DB) *GORMRepository[T] {
	return &GORMRepository[T]{db: db}
}

func (r GORMRepository[T]) Q(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

func (r GORMRepository[T]) HandleError(tx *gorm.DB) error {
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return preseterrors.ErrNotFound
		}
		return tx.Error
	}
	return nil
}

func (r GORMRepository[T]) Exists(ctx context.Context, clause entity.Clause) (bool, error) {
	tx := gormentity.ParseClause[T](r.Q(ctx), clause).Model((*T)(nil)).Find(&struct{}{})
	if err := r.HandleError(tx); err != nil {
		return false, err
	}
	return tx.RowsAffected > 0, nil
}

func (r GORMRepository[T]) Count(ctx context.Context, clause entity.Clause) (int, error) {
	var count int64
	tx := gormentity.ParseClause[T](r.Q(ctx), clause).Model((*T)(nil)).Count(&count)
	if err := r.HandleError(tx); err != nil {
		return 0, err
	}
	return int(count), nil
}

func (r GORMRepository[T]) FindMany(ctx context.Context, clause entity.Clause, options repository.FindManyOptions) ([]*T, error) {
	var ents []*T
	tx := r.parseManyOptions(gormentity.ParseClause[T](r.Q(ctx), clause), options).Find(&ents)
	if err := r.HandleError(tx); err != nil {
		return nil, err
	}
	return ents, nil
}

func (r GORMRepository[T]) parseManyOptions(tx *gorm.DB, options repository.FindManyOptions) *gorm.DB {
	for _, ordering := range options.Orderings {
		tx = tx.Order(clause.OrderByColumn{
			Column: clause.Column{Name: gormentity.ColumnName[T](tx, ordering.By)},
			Desc:   ordering.Desc,
		})
	}
	return r.addPreload(tx.Limit(options.Limit).Offset(options.Offset), options.Relations)
}

func (r GORMRepository[T]) Find(ctx context.Context, clause entity.Clause, options repository.FindOptions) (*T, error) {
	e := new(T)
	tx := r.parseOptions(gormentity.ParseClause[T](r.Q(ctx), clause), options).First(e)
	if err := r.HandleError(tx); err != nil {
		return nil, err
	}
	return e, nil
}

func (r GORMRepository[T]) parseOptions(tx *gorm.DB, options repository.FindOptions) *gorm.DB {
	return r.addPreload(tx, options.Relations)
}

func (r GORMRepository[T]) addPreload(tx *gorm.DB, relations []string) *gorm.DB {
	for _, relation := range relations {
		tx = tx.Preload(relation)
	}
	return tx
}

func (r GORMRepository[T]) Create(ctx context.Context, e *T) (*T, error) {
	tx := r.Q(ctx).Create(e)
	if err := r.HandleError(tx); err != nil {
		return nil, err
	}
	return e, nil
}

func (r GORMRepository[T]) Update(ctx context.Context, e *T) error {
	tx := r.Q(ctx).Save(e)
	if err := r.HandleError(tx); err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return preseterrors.ErrNotFound
	}
	return nil
}

func (r GORMRepository[T]) Delete(ctx context.Context, clause entity.Clause) error {
	var e T
	tx := gormentity.ParseClause[T](r.Q(ctx), clause).Delete(&e)
	if err := r.HandleError(tx); err != nil {
		return err
	}
	if tx.RowsAffected == 0 {
		return preseterrors.ErrNotFound
	}
	return nil
}
