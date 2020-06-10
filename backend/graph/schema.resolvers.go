package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/awolk/lil-shop/backend/graph/generated"
	"github.com/awolk/lil-shop/backend/graph/model"
)

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	// TODO: implement
	return nil, nil
}

func (r *queryResolver) Cart(ctx context.Context, id string) (*model.Cart, error) {
	// TODO: implement
	return &model.Cart{}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
