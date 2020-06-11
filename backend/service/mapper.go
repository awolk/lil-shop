package service

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent"
)

func entityToItem(entity *ent.Item) Item {
	return Item{
		ID:        entity.ID,
		Name:      entity.Name,
		CostCents: entity.CostCents,
	}
}

func entityToLineItem(ctx context.Context, entity *ent.LineItem) (LineItem, error) {
	itemEntity, err := entity.Edges.ItemOrErr()
	if err != nil {
		if ent.IsNotLoaded(err) {
			itemEntity, err = entity.QueryItem().Only(ctx)
		}
	}
	if err != nil {
		return LineItem{}, fmt.Errorf("unable to get item for line item: %w", err)
	}

	return LineItem{
		ID:       entity.ID,
		Item:     entityToItem(itemEntity),
		Quantity: entity.Quantity,
	}, nil
}

func entityToCart(ctx context.Context, entity *ent.Cart) (Cart, error) {
	lineItemEntities, err := entity.Edges.LineItemsOrErr()
	if err != nil {
		if ent.IsNotLoaded(err) {
			lineItemEntities, err = entity.QueryLineItems().WithItem().All(ctx)
		}
	}
	if err != nil {
		return Cart{}, fmt.Errorf("unable to get line items for cart: %w", err)
	}

	lineItems := make([]LineItem, 0, len(lineItemEntities))
	for _, lineItemEntity := range lineItemEntities {
		lineItem, err := entityToLineItem(ctx, lineItemEntity)
		if err != nil {
			return Cart{}, fmt.Errorf("unable to map line item: %w", err)
		}
		lineItems = append(lineItems, lineItem)
	}

	return Cart{
		ID:        entity.ID,
		LineItems: lineItems,
	}, nil
}
