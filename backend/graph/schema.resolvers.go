package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/graph/generated"
	"github.com/awolk/lil-shop/backend/graph/model"
	"github.com/awolk/lil-shop/backend/item"
)

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	items, err := r.ItemService.GetItems(ctx)
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
	// TODO: implement
	return &model.Cart{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func itemToModel(item item.Item) *model.Item {
	return &model.Item{
		ID:        item.ID.String(),
		Name:      item.Name,
		CostCents: item.CostCents,
	}
}
