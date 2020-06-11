package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/graph/generated"
	"github.com/awolk/lil-shop/backend/graph/model"

	"github.com/google/uuid"
)

func (r *mutationResolver) NewCart(ctx context.Context) (*model.Cart, error) {
	cart, err := r.Service.NewCart(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to create cart: %w", err)
	}

	return cartToModel(cart), nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	items, err := r.Service.GetItems(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to get items: %w", err)
	}

	res := make([]*model.Item, 0, len(items))
	for _, item := range items {
		model := itemToModel(item)
		res = append(res, model)
	}
	return res, nil
}

func (r *queryResolver) Cart(ctx context.Context, id string) (*model.Cart, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	cart, err := r.Service.GetCart(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("unable to get cart: %w", err)
	}

	return cartToModel(cart), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
