package service

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent"
	"github.com/awolk/lil-shop/backend/ent/cart"
	"github.com/awolk/lil-shop/backend/ent/item"
	"github.com/awolk/lil-shop/backend/ent/lineitem"
	"github.com/awolk/lil-shop/backend/payments"
	"github.com/google/uuid"
)

type Service struct {
	client          *ent.Client
	paymentsService *payments.PaymentsService
}

func New(client *ent.Client, paymentsService *payments.PaymentsService) *Service {
	return &Service{
		client:          client,
		paymentsService: paymentsService,
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

func (s *Service) AddItemToCart(
	ctx context.Context,
	itemID uuid.UUID,
	quantity int,
	cartID uuid.UUID) (LineItem, error) {

	var lineItem LineItem

	err := s.withTx(ctx, func(tx *ent.Tx) error {
		lineItemEntity, err := tx.LineItem.
			Query().
			WithItem().
			Where(
				lineitem.HasCartWith(cart.ID(cartID)),
				lineitem.HasItemWith(item.ID(itemID)),
			).Only(ctx)

		if err != nil {
			if ent.IsNotFound(err) {
				lineItemEntity, err = tx.LineItem.Create().
					SetItemID(itemID).
					SetCartID(cartID).
					SetQuantity(quantity).
					Save(ctx)
			} else {
				return fmt.Errorf("unable to fetch line item: %w", err)
			}
		} else {
			lineItemEntity, err = lineItemEntity.Update().SetQuantity(lineItemEntity.Quantity + quantity).Save(ctx)
		}

		if err != nil {
			return fmt.Errorf("unable to create or update line item: %w", err)
		}

		lineItem, err = entityToLineItem(ctx, lineItemEntity)
		if err != nil {
			return fmt.Errorf("unable to map line item: %w", err)
		}

		return nil
	})

	if err != nil {
		return LineItem{}, err
	}

	return lineItem, nil
}

type CheckoutReply struct {
	TotalCostCents int
	ClientSecret   string
}

func (s *Service) CheckoutCart(ctx context.Context, cartID uuid.UUID) (*CheckoutReply, error) {
	var pi *payments.PaymentIntent
	var totalCostCents int

	err := s.withTx(ctx, func(tx *ent.Tx) error {
		cartEnt, err := tx.Cart.Get(ctx, cartID)
		if err != nil {
			return fmt.Errorf("failed to load cart: %w", err)
		}

		var cost struct {
			Sum int `json:"sum"`
		}

		err = tx.Cart.
			Query().
			Where(cart.ID(cartID)).
			WithLineItems(func(liq *ent.LineItemQuery) {
				liq.WithItem().Select("quantity * cost_cents as line_item_cost_cents")
			}).
			GroupBy(cart.FieldID).
			Aggregate(ent.Sum("line_item_cost_cents")).
			Scan(ctx, &cost)
		if err != nil {
			return fmt.Errorf("failed to sum cart cost: %w", err)
		}

		totalCostCents = cost.Sum

		if cartEnt.PaymentIntentID == "" {
			// need to create new payment intent
			pi, err = s.paymentsService.NewPaymentIntent(ctx, totalCostCents)
		} else {
			// need to update payment intent with new cost
			pi, err = s.paymentsService.UpdatePaymentIntent(ctx, cartEnt.PaymentIntentID, totalCostCents)
		}
		if err != nil {
			return fmt.Errorf("failed to create/update payment intent: %w", err)
		}

		cartEnt.Update().SetPaymentIntentID(pi.ID).Save(ctx)

		return nil
	})
	if err != nil {
		return nil, err
	}

	return &CheckoutReply{
		TotalCostCents: totalCostCents,
		ClientSecret:   pi.ClientSecret,
	}, nil
}
