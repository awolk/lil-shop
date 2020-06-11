package service

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent"
	"github.com/awolk/lil-shop/backend/ent/cart"
	"github.com/google/uuid"
)

type Service struct {
	client *ent.Client
}

func New(client *ent.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) GetItems(ctx context.Context) ([]Item, error) {
	entities, err := s.client.Item.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch items: %w", err)
	}

	res := make([]Item, 0, len(entities))
	for _, entity := range entities {
		res = append(res, entityToItem(entity))
	}

	return res, nil
}

func (s *Service) NewItem(ctx context.Context, name string, costCents int) (Item, error) {
	entity, err := s.client.Item.Create().SetName(name).SetCostCents(costCents).Save(ctx)
	if err != nil {
		return Item{}, fmt.Errorf("unable to save item: %w", err)
	}

	return entityToItem(entity), nil
}

func (s *Service) NewCart(ctx context.Context) (Cart, error) {
	entity, err := s.client.Cart.Create().Save(ctx)
	if err != nil {
		return Cart{}, fmt.Errorf("unable to save cart: %w", err)
	}

	cart, err := entityToCart(ctx, entity)
	if err != nil {
		return Cart{}, fmt.Errorf("unable to map cart: %w", err)
	}

	return cart, nil
}

func (s *Service) GetCart(ctx context.Context, id uuid.UUID) (Cart, error) {
	entity, err := s.client.Cart.Query().
		Where(cart.ID(id)).
		WithLineItems(func(q *ent.LineItemQuery) {
			q.WithItem()
		}).
		Only(ctx)

	if err != nil {
		return Cart{}, fmt.Errorf("unable to load cart: %w", err)
	}

	cart, err := entityToCart(ctx, entity)
	if err != nil {
		return Cart{}, fmt.Errorf("unable to map cart: %w", err)
	}

	return cart, nil
}
