package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent"
	"github.com/awolk/lil-shop/backend/graph/generated"
	"github.com/awolk/lil-shop/backend/graph/model"
	"github.com/google/uuid"
)

func (r *cartResolver) LineItems(ctx context.Context, obj *ent.Cart) ([]*ent.LineItem, error) {
	items, err := r.Service.LineItemsFromCart(ctx, obj)
	if err != nil {
		return nil, fmt.Errorf("unable to load line items: %w", err)
	}
	return items, nil
}

func (r *lineItemResolver) Item(ctx context.Context, obj *ent.LineItem) (*ent.Item, error) {
	item, err := r.Service.ItemFromLineItem(ctx, obj)
	if err != nil {
		return nil, fmt.Errorf("unable to load item: %w", err)
	}
	return item, nil
}

func (r *mutationResolver) NewCart(ctx context.Context) (*ent.Cart, error) {
	cart, err := r.Service.NewCart(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to create cart: %w", err)
	}

	return cart, nil
}

func (r *mutationResolver) AddItemToCart(ctx context.Context, itemID string, quantity int, cartID string) (*bool, error) {
	itemUUID, err := uuid.Parse(itemID)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	cartUUID, err := uuid.Parse(cartID)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	err = r.Service.AddItemToCart(ctx, itemUUID, quantity, cartUUID)
	if err != nil {
		return nil, fmt.Errorf("unable to add item to cart: %w", err)
	}

	return nil, nil
}

func (r *mutationResolver) CheckoutCart(ctx context.Context, cartID string) (*model.Order, error) {
	cartUUID, err := uuid.Parse(cartID)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}
	_ = cartUUID

	order, err := r.Service.CheckoutCart(ctx, cartUUID)
	if err != nil {
		return nil, fmt.Errorf("failed to checkout cart: %w", err)
	}

	return order, nil
}

func (r *orderResolver) OrderLineItems(ctx context.Context, obj *model.Order) ([]*ent.OrderLineItem, error) {
	orderLineItems, err := r.Service.OrderLineItemsFromOrder(ctx, obj.Order)
	if err != nil {
		return nil, fmt.Errorf("unable to load order line items: %w", err)
	}
	return orderLineItems, nil
}

func (r *orderLineItemResolver) Item(ctx context.Context, obj *ent.OrderLineItem) (*ent.Item, error) {
	item, err := r.Service.ItemFromOrderLineItem(ctx, obj)
	if err != nil {
		return nil, fmt.Errorf("unable to load item: %w", err)
	}
	return item, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*ent.Item, error) {
	items, err := r.Service.GetItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get items: %w", err)
	}

	return items, nil
}

func (r *queryResolver) Cart(ctx context.Context, id string) (*ent.Cart, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	cart, err := r.Service.GetCart(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("unable to get cart: %w", err)
	}

	return cart, nil
}

// Cart returns generated.CartResolver implementation.
func (r *Resolver) Cart() generated.CartResolver { return &cartResolver{r} }

// LineItem returns generated.LineItemResolver implementation.
func (r *Resolver) LineItem() generated.LineItemResolver { return &lineItemResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Order returns generated.OrderResolver implementation.
func (r *Resolver) Order() generated.OrderResolver { return &orderResolver{r} }

// OrderLineItem returns generated.OrderLineItemResolver implementation.
func (r *Resolver) OrderLineItem() generated.OrderLineItemResolver { return &orderLineItemResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type cartResolver struct{ *Resolver }
type lineItemResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type orderResolver struct{ *Resolver }
type orderLineItemResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
