// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent/cart"
	"github.com/awolk/lil-shop/backend/ent/item"
	"github.com/awolk/lil-shop/backend/ent/lineitem"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// LineItemCreate is the builder for creating a LineItem entity.
type LineItemCreate struct {
	config
	mutation *LineItemMutation
	hooks    []Hook
}

// SetQuantity sets the quantity field.
func (lic *LineItemCreate) SetQuantity(i int) *LineItemCreate {
	lic.mutation.SetQuantity(i)
	return lic
}

// SetID sets the id field.
func (lic *LineItemCreate) SetID(u uuid.UUID) *LineItemCreate {
	lic.mutation.SetID(u)
	return lic
}

// SetItemID sets the item edge to Item by id.
func (lic *LineItemCreate) SetItemID(id uuid.UUID) *LineItemCreate {
	lic.mutation.SetItemID(id)
	return lic
}

// SetItem sets the item edge to Item.
func (lic *LineItemCreate) SetItem(i *Item) *LineItemCreate {
	return lic.SetItemID(i.ID)
}

// SetCartID sets the cart edge to Cart by id.
func (lic *LineItemCreate) SetCartID(id uuid.UUID) *LineItemCreate {
	lic.mutation.SetCartID(id)
	return lic
}

// SetCart sets the cart edge to Cart.
func (lic *LineItemCreate) SetCart(c *Cart) *LineItemCreate {
	return lic.SetCartID(c.ID)
}

// Mutation returns the LineItemMutation object of the builder.
func (lic *LineItemCreate) Mutation() *LineItemMutation {
	return lic.mutation
}

// Save creates the LineItem in the database.
func (lic *LineItemCreate) Save(ctx context.Context) (*LineItem, error) {
	if _, ok := lic.mutation.Quantity(); !ok {
		return nil, errors.New("ent: missing required field \"quantity\"")
	}
	if v, ok := lic.mutation.Quantity(); ok {
		if err := lineitem.QuantityValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"quantity\": %w", err)
		}
	}
	if _, ok := lic.mutation.ID(); !ok {
		v := lineitem.DefaultID()
		lic.mutation.SetID(v)
	}
	if _, ok := lic.mutation.ItemID(); !ok {
		return nil, errors.New("ent: missing required edge \"item\"")
	}
	if _, ok := lic.mutation.CartID(); !ok {
		return nil, errors.New("ent: missing required edge \"cart\"")
	}
	var (
		err  error
		node *LineItem
	)
	if len(lic.hooks) == 0 {
		node, err = lic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LineItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lic.mutation = mutation
			node, err = lic.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(lic.hooks) - 1; i >= 0; i-- {
			mut = lic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (lic *LineItemCreate) SaveX(ctx context.Context) *LineItem {
	v, err := lic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (lic *LineItemCreate) sqlSave(ctx context.Context) (*LineItem, error) {
	var (
		li    = &LineItem{config: lic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: lineitem.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: lineitem.FieldID,
			},
		}
	)
	if id, ok := lic.mutation.ID(); ok {
		li.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lic.mutation.Quantity(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: lineitem.FieldQuantity,
		})
		li.Quantity = value
	}
	if nodes := lic.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   lineitem.ItemTable,
			Columns: []string{lineitem.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := lic.mutation.CartIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   lineitem.CartTable,
			Columns: []string{lineitem.CartColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: cart.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, lic.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return li, nil
}
