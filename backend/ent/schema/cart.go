package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
)

// Cart holds the schema definition for the Cart entity.
type Cart struct {
	ent.Schema
}

// Fields of the Cart.
func (Cart) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("payment_intent_id").Optional(),
	}
}

// Edges of the Cart.
func (Cart) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("line_items", LineItem.Type),
	}
}
