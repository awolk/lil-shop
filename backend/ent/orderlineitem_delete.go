// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent/orderlineitem"
	"github.com/awolk/lil-shop/backend/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// OrderLineItemDelete is the builder for deleting a OrderLineItem entity.
type OrderLineItemDelete struct {
	config
	hooks      []Hook
	mutation   *OrderLineItemMutation
	predicates []predicate.OrderLineItem
}

// Where adds a new predicate to the delete builder.
func (olid *OrderLineItemDelete) Where(ps ...predicate.OrderLineItem) *OrderLineItemDelete {
	olid.predicates = append(olid.predicates, ps...)
	return olid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (olid *OrderLineItemDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(olid.hooks) == 0 {
		affected, err = olid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderLineItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			olid.mutation = mutation
			affected, err = olid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(olid.hooks) - 1; i >= 0; i-- {
			mut = olid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, olid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (olid *OrderLineItemDelete) ExecX(ctx context.Context) int {
	n, err := olid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (olid *OrderLineItemDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: orderlineitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: orderlineitem.FieldID,
			},
		},
	}
	if ps := olid.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, olid.driver, _spec)
}

// OrderLineItemDeleteOne is the builder for deleting a single OrderLineItem entity.
type OrderLineItemDeleteOne struct {
	olid *OrderLineItemDelete
}

// Exec executes the deletion query.
func (olido *OrderLineItemDeleteOne) Exec(ctx context.Context) error {
	n, err := olido.olid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orderlineitem.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (olido *OrderLineItemDeleteOne) ExecX(ctx context.Context) {
	olido.olid.ExecX(ctx)
}
