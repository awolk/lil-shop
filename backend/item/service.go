package item

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent"
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
