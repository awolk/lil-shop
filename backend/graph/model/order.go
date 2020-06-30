package model

import "github.com/awolk/lil-shop/backend/ent"

// Order holds information about an order passed to the client
type Order struct {
	*ent.Order
	TotalCostCents int
	ClientSecret   string
}
