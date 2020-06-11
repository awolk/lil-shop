package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// LineItem holds the schema definition for the LineItem entity.
type LineItem struct {
	ent.Schema
}

// Fields of the LineItem.
func (LineItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.Int("quantity").Positive(),
	}
}

// Edges of the LineItem.
func (LineItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", Item.Type).Unique().Required(),
		edge.From("cart", Cart.Type).Ref("line_items").Unique().Required(),
	}
}
