// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent/cart"
	"github.com/awolk/lil-shop/backend/ent/lineitem"
	"github.com/awolk/lil-shop/backend/ent/predicate"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// CartUpdate is the builder for updating Cart entities.
type CartUpdate struct {
	config
	hooks      []Hook
	mutation   *CartMutation
	predicates []predicate.Cart
}

// Where adds a new predicate for the builder.
func (cu *CartUpdate) Where(ps ...predicate.Cart) *CartUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// AddLineItemIDs adds the line_items edge to LineItem by ids.
func (cu *CartUpdate) AddLineItemIDs(ids ...uuid.UUID) *CartUpdate {
	cu.mutation.AddLineItemIDs(ids...)
	return cu
}

// AddLineItems adds the line_items edges to LineItem.
func (cu *CartUpdate) AddLineItems(l ...*LineItem) *CartUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cu.AddLineItemIDs(ids...)
}

// Mutation returns the CartMutation object of the builder.
func (cu *CartUpdate) Mutation() *CartMutation {
	return cu.mutation
}

// RemoveLineItemIDs removes the line_items edge to LineItem by ids.
func (cu *CartUpdate) RemoveLineItemIDs(ids ...uuid.UUID) *CartUpdate {
	cu.mutation.RemoveLineItemIDs(ids...)
	return cu
}

// RemoveLineItems removes line_items edges to LineItem.
func (cu *CartUpdate) RemoveLineItems(l ...*LineItem) *CartUpdate {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cu.RemoveLineItemIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CartUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CartUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CartUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CartUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cart.Table,
			Columns: cart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cart.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if nodes := cu.mutation.RemovedLineItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.LineItemsTable,
			Columns: []string{cart.LineItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: lineitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.LineItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.LineItemsTable,
			Columns: []string{cart.LineItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: lineitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CartUpdateOne is the builder for updating a single Cart entity.
type CartUpdateOne struct {
	config
	hooks    []Hook
	mutation *CartMutation
}

// AddLineItemIDs adds the line_items edge to LineItem by ids.
func (cuo *CartUpdateOne) AddLineItemIDs(ids ...uuid.UUID) *CartUpdateOne {
	cuo.mutation.AddLineItemIDs(ids...)
	return cuo
}

// AddLineItems adds the line_items edges to LineItem.
func (cuo *CartUpdateOne) AddLineItems(l ...*LineItem) *CartUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cuo.AddLineItemIDs(ids...)
}

// Mutation returns the CartMutation object of the builder.
func (cuo *CartUpdateOne) Mutation() *CartMutation {
	return cuo.mutation
}

// RemoveLineItemIDs removes the line_items edge to LineItem by ids.
func (cuo *CartUpdateOne) RemoveLineItemIDs(ids ...uuid.UUID) *CartUpdateOne {
	cuo.mutation.RemoveLineItemIDs(ids...)
	return cuo
}

// RemoveLineItems removes line_items edges to LineItem.
func (cuo *CartUpdateOne) RemoveLineItems(l ...*LineItem) *CartUpdateOne {
	ids := make([]uuid.UUID, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return cuo.RemoveLineItemIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CartUpdateOne) Save(ctx context.Context) (*Cart, error) {

	var (
		err  error
		node *Cart
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CartMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CartUpdateOne) SaveX(ctx context.Context) *Cart {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CartUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CartUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CartUpdateOne) sqlSave(ctx context.Context) (c *Cart, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   cart.Table,
			Columns: cart.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: cart.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Cart.ID for update")}
	}
	_spec.Node.ID.Value = id
	if nodes := cuo.mutation.RemovedLineItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.LineItemsTable,
			Columns: []string{cart.LineItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: lineitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.LineItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   cart.LineItemsTable,
			Columns: []string{cart.LineItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: lineitem.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Cart{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{cart.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
