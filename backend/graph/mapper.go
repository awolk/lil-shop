package graph

import (
	"github.com/awolk/lil-shop/backend/graph/model"
	"github.com/awolk/lil-shop/backend/service"
)

func itemToModel(item service.Item) *model.Item {
	return &model.Item{
		ID:        item.ID.String(),
		Name:      item.Name,
		CostCents: item.CostCents,
	}
}

func lineItemToModel(lineItem service.LineItem) *model.LineItem {
	return &model.LineItem{
		ID:       lineItem.ID.String(),
		Item:     itemToModel(lineItem.Item),
		Quantity: lineItem.Quantity,
	}
}

func cartToModel(cart service.Cart) *model.Cart {
	lineItemModels := make([]*model.LineItem, 0, len(cart.LineItems))
	for _, lineItem := range cart.LineItems {
		lineItemModels = append(lineItemModels, lineItemToModel(lineItem))
	}

	return &model.Cart{
		ID:        cart.ID.String(),
		LineItems: lineItemModels,
	}
}
