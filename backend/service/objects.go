package service

import (
	"github.com/google/uuid"
)

type Item struct {
	ID        uuid.UUID
	Name      string
	CostCents int
}

type LineItem struct {
	ID       uuid.UUID
	Item     Item
	Quantity int
}

type Cart struct {
	ID        uuid.UUID
	LineItems []LineItem
}
