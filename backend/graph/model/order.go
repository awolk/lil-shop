package model

import "github.com/awolk/lil-shop/backend/ent"

type Order struct {
	*ent.Order
	TotalCostCents int
	ClientSecret   string
}
