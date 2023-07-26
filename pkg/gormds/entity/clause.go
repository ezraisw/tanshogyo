package entity

import (
	"fmt"

	"github.com/pwnedgod/tanshogyo/pkg/common/entity"
	"gorm.io/gorm"
)

func ParseClause[T any](tx *gorm.DB, clause entity.Clause) *gorm.DB {
	if clause == nil {
		return tx
	}
	if attacher, ok := clause(&conjunction[T]{}).(expression); ok {
		return attacher.exprAttach(tx)
	}
	return tx
}

type expression interface {
	exprAttach(*gorm.DB) *gorm.DB
}

type isComparison[T any] struct {
	pField *field[T]

	operator string
	value    any
}

func (c *isComparison[T]) And() entity.Prober {
	return &conjunction[T]{pExpression: c}
}

func (c *isComparison[T]) Or() entity.Prober {
	return &conjunction[T]{pExpression: c, or: true}
}

func (c isComparison[T]) exprAttach(tx *gorm.DB) *gorm.DB {
	return c.pField.fieldAttach(tx, "%s "+c.operator+" ?", c.value)
}

type isLikeComparison[T any] struct {
	pField *field[T]

	pattern string
}

func (c *isLikeComparison[T]) And() entity.Prober {
	return &conjunction[T]{pExpression: c}
}

func (c *isLikeComparison[T]) Or() entity.Prober {
	return &conjunction[T]{pExpression: c, or: true}
}

func (c isLikeComparison[T]) exprAttach(tx *gorm.DB) *gorm.DB {
	return c.pField.fieldAttach(tx, "%s LIKE ?", c.pattern)
}

type isInComparison[T any] struct {
	pField *field[T]

	values []any
}

func (c *isInComparison[T]) And() entity.Prober {
	return &conjunction[T]{pExpression: c}
}

func (c *isInComparison[T]) Or() entity.Prober {
	return &conjunction[T]{pExpression: c, or: true}
}

func (c isInComparison[T]) exprAttach(tx *gorm.DB) *gorm.DB {
	return c.pField.fieldAttach(tx, "%s IN ?", c.values...)
}

type isBetweenComparison[T any] struct {
	pField *field[T]

	start any
	end   any
}

func (c *isBetweenComparison[T]) And() entity.Prober {
	return &conjunction[T]{pExpression: c}
}

func (c *isBetweenComparison[T]) Or() entity.Prober {
	return &conjunction[T]{pExpression: c, or: true}
}

func (c isBetweenComparison[T]) exprAttach(tx *gorm.DB) *gorm.DB {
	return c.pField.fieldAttach(tx, "%s BETWEEN ? AND ?", c.start, c.end)
}

type group[T any] struct {
	pConjunction *conjunction[T]

	clause entity.Clause
	not    bool
}

func (g *group[T]) And() entity.Prober {
	return &conjunction[T]{pExpression: g}
}

func (g *group[T]) Or() entity.Prober {
	return &conjunction[T]{pExpression: g, or: true}
}

func (g group[T]) exprAttach(tx *gorm.DB) *gorm.DB {
	if g.clause == nil {
		return g.pConjunction.child(tx)
	}

	// The conjunction for the start of the group.
	// Similar logic as ParseClause.
	var groupTx *gorm.DB
	if attacher, ok := g.clause(&conjunction[T]{}).(expression); ok {
		groupTx = attacher.exprAttach(tx)
	}

	if groupTx == nil {
		return g.pConjunction.child(tx)
	}

	if g.not {
		groupTx = tx.Not(groupTx)
	}

	return g.pConjunction.conjAttach(tx, groupTx)
}

type field[T any] struct {
	pConjunction *conjunction[T]

	name string
}

func (f *field[T]) Is(operator string, value any) entity.Expression {
	return &isComparison[T]{pField: f, operator: operator, value: value}
}

func (f *field[T]) IsLike(pattern string) entity.Expression {
	return &isLikeComparison[T]{pField: f, pattern: pattern}
}

func (f *field[T]) IsIn(values ...any) entity.Expression {
	return &isInComparison[T]{pField: f, values: values}
}

func (f *field[T]) IsBetween(start, end any) entity.Expression {
	return &isBetweenComparison[T]{pField: f, start: start, end: end}
}

func (f field[T]) fieldAttach(tx *gorm.DB, queryFormat string, values ...any) *gorm.DB {
	query := fmt.Sprintf(queryFormat, ColumnName[T](tx, f.name))
	return f.pConjunction.conjAttach(tx, query, values...)
}

type conjunction[T any] struct {
	pExpression expression

	or bool
}

func (c *conjunction[T]) Field(name string) entity.Field {
	return &field[T]{pConjunction: c, name: name}
}

func (c *conjunction[T]) Group(clause entity.Clause) entity.Expression {
	return &group[T]{pConjunction: c, clause: clause}
}

func (c *conjunction[T]) NotGroup(clause entity.Clause) entity.Expression {
	return &group[T]{pConjunction: c, clause: clause, not: true}
}

func (c conjunction[T]) conjAttach(tx *gorm.DB, query any, args ...any) *gorm.DB {
	tx = c.child(tx)
	if c.or {
		return tx.Or(query, args...)
	}
	return tx.Where(query, args...)
}

func (c conjunction[T]) child(tx *gorm.DB) *gorm.DB {
	// Base condition.
	if c.pExpression != nil {
		tx = c.pExpression.exprAttach(tx)
	}
	return tx
}
