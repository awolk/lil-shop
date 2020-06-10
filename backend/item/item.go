package item

import (
	"github.com/awolk/lil-shop/backend/ent"
	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID
	Name      string
	CostCents int
}

func entityToItem(entity *ent.Item) Item {
	return Item{
		ID:        entity.ID,
		Name:      entity.Name,
		CostCents: entity.CostCents,
	}
}
