package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// OrderLineItem holds the schema definition for the OrderLineItem entity.
type OrderLineItem struct {
	ent.Schema
}

// Fields of the OrderLineItem.
func (OrderLineItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("quantity").Positive(),
		field.Int("unit_cost_cents").Positive(),
		field.Bool("completed").Default(false),
	}
}

// Edges of the OrderLineItem.
func (OrderLineItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", Item.Type).Unique().Required(),
		edge.From("order", Order.Type).Ref("order_line_items").Unique().Required(),

		// keep track of the original line item to delete from cart once the order is processed
		edge.To("original_line_item", LineItem.Type).Unique(),
	}
}
