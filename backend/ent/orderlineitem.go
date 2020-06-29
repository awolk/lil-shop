// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/awolk/lil-shop/backend/ent/item"
	"github.com/awolk/lil-shop/backend/ent/lineitem"
	"github.com/awolk/lil-shop/backend/ent/order"
	"github.com/awolk/lil-shop/backend/ent/orderlineitem"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
)

// OrderLineItem is the model entity for the OrderLineItem schema.
type OrderLineItem struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// UnitCostCents holds the value of the "unit_cost_cents" field.
	UnitCostCents int `json:"unit_cost_cents,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderLineItemQuery when eager-loading is set.
	Edges                              OrderLineItemEdges `json:"edges"`
	order_order_line_items             *uuid.UUID
	order_line_item_item               *uuid.UUID
	order_line_item_original_line_item *uuid.UUID
}

// OrderLineItemEdges holds the relations/edges for other nodes in the graph.
type OrderLineItemEdges struct {
	// Item holds the value of the item edge.
	Item *Item
	// Order holds the value of the order edge.
	Order *Order
	// OriginalLineItem holds the value of the original_line_item edge.
	OriginalLineItem *LineItem
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderLineItemEdges) ItemOrErr() (*Item, error) {
	if e.loadedTypes[0] {
		if e.Item == nil {
			// The edge item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: item.Label}
		}
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderLineItemEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[1] {
		if e.Order == nil {
			// The edge order was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// OriginalLineItemOrErr returns the OriginalLineItem value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderLineItemEdges) OriginalLineItemOrErr() (*LineItem, error) {
	if e.loadedTypes[2] {
		if e.OriginalLineItem == nil {
			// The edge original_line_item was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: lineitem.Label}
		}
		return e.OriginalLineItem, nil
	}
	return nil, &NotLoadedError{edge: "original_line_item"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderLineItem) scanValues() []interface{} {
	return []interface{}{
		&uuid.UUID{},     // id
		&sql.NullInt64{}, // quantity
		&sql.NullInt64{}, // unit_cost_cents
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*OrderLineItem) fkValues() []interface{} {
	return []interface{}{
		&uuid.UUID{}, // order_order_line_items
		&uuid.UUID{}, // order_line_item_item
		&uuid.UUID{}, // order_line_item_original_line_item
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderLineItem fields.
func (oli *OrderLineItem) assignValues(values ...interface{}) error {
	if m, n := len(values), len(orderlineitem.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	if value, ok := values[0].(*uuid.UUID); !ok {
		return fmt.Errorf("unexpected type %T for field id", values[0])
	} else if value != nil {
		oli.ID = *value
	}
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quantity", values[0])
	} else if value.Valid {
		oli.Quantity = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field unit_cost_cents", values[1])
	} else if value.Valid {
		oli.UnitCostCents = int(value.Int64)
	}
	values = values[2:]
	if len(values) == len(orderlineitem.ForeignKeys) {
		if value, ok := values[0].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field order_order_line_items", values[0])
		} else if value != nil {
			oli.order_order_line_items = value
		}
		if value, ok := values[1].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field order_line_item_item", values[1])
		} else if value != nil {
			oli.order_line_item_item = value
		}
		if value, ok := values[2].(*uuid.UUID); !ok {
			return fmt.Errorf("unexpected type %T for field order_line_item_original_line_item", values[2])
		} else if value != nil {
			oli.order_line_item_original_line_item = value
		}
	}
	return nil
}

// QueryItem queries the item edge of the OrderLineItem.
func (oli *OrderLineItem) QueryItem() *ItemQuery {
	return (&OrderLineItemClient{config: oli.config}).QueryItem(oli)
}

// QueryOrder queries the order edge of the OrderLineItem.
func (oli *OrderLineItem) QueryOrder() *OrderQuery {
	return (&OrderLineItemClient{config: oli.config}).QueryOrder(oli)
}

// QueryOriginalLineItem queries the original_line_item edge of the OrderLineItem.
func (oli *OrderLineItem) QueryOriginalLineItem() *LineItemQuery {
	return (&OrderLineItemClient{config: oli.config}).QueryOriginalLineItem(oli)
}

// Update returns a builder for updating this OrderLineItem.
// Note that, you need to call OrderLineItem.Unwrap() before calling this method, if this OrderLineItem
// was returned from a transaction, and the transaction was committed or rolled back.
func (oli *OrderLineItem) Update() *OrderLineItemUpdateOne {
	return (&OrderLineItemClient{config: oli.config}).UpdateOne(oli)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (oli *OrderLineItem) Unwrap() *OrderLineItem {
	tx, ok := oli.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderLineItem is not a transactional entity")
	}
	oli.config.driver = tx.drv
	return oli
}

// String implements the fmt.Stringer.
func (oli *OrderLineItem) String() string {
	var builder strings.Builder
	builder.WriteString("OrderLineItem(")
	builder.WriteString(fmt.Sprintf("id=%v", oli.ID))
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", oli.Quantity))
	builder.WriteString(", unit_cost_cents=")
	builder.WriteString(fmt.Sprintf("%v", oli.UnitCostCents))
	builder.WriteByte(')')
	return builder.String()
}

// OrderLineItems is a parsable slice of OrderLineItem.
type OrderLineItems []*OrderLineItem

func (oli OrderLineItems) config(cfg config) {
	for _i := range oli {
		oli[_i].config = cfg
	}
}
