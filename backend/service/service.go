package service

import (
	"context"
	"fmt"

	"github.com/awolk/lil-shop/backend/ent"
	"github.com/awolk/lil-shop/backend/ent/cart"
	"github.com/awolk/lil-shop/backend/ent/item"
	"github.com/awolk/lil-shop/backend/ent/lineitem"
	"github.com/awolk/lil-shop/backend/graph/model"
	"github.com/awolk/lil-shop/backend/payments"
	"github.com/google/uuid"
)

// Service handles shop business logic
type Service struct {
	globalClient    *ent.Client
	paymentsService *payments.PaymentsService
}

// New constructs a new Service
func New(client *ent.Client, paymentsService *payments.PaymentsService) *Service {
	return &Service{
		globalClient:    client,
		paymentsService: paymentsService,
	}
}

// GetItems fetches all shop items
func (s *Service) GetItems(ctx context.Context) ([]*ent.Item, error) {
	client := s.client(ctx)
	items, err := client.Item.Query().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch items: %w", err)
	}

	return items, nil
}

// LineItemsFromCart fetches the line items in a cart
func (s *Service) LineItemsFromCart(ctx context.Context, cart *ent.Cart) ([]*ent.LineItem, error) {
	client := s.client(ctx)
	items, err := cart.Edges.LineItemsOrErr()
	if err != nil {
		if ent.IsNotLoaded(err) {
			return client.Cart.QueryLineItems(cart).WithItem().All(ctx)
		}
		return nil, err
	}
	return items, nil
}

// ItemFromLineItem fetches the associated item for a line-item
func (s *Service) ItemFromLineItem(ctx context.Context, lineItem *ent.LineItem) (*ent.Item, error) {
	client := s.client(ctx)
	item, err := lineItem.Edges.ItemOrErr()
	if err != nil {
		if ent.IsNotLoaded(err) {
			return client.LineItem.QueryItem(lineItem).Only(ctx)
		}
		return nil, err
	}
	return item, nil
}

// OrderLineItemsFromOrder fetches the line items in an order
func (s *Service) OrderLineItemsFromOrder(ctx context.Context, order *ent.Order) ([]*ent.OrderLineItem, error) {
	client := s.client(ctx)
	items, err := order.Edges.OrderLineItemsOrErr()
	if err != nil {
		if ent.IsNotLoaded(err) {
			return client.Order.QueryOrderLineItems(order).WithItem().All(ctx)
		}
		return nil, err
	}
	return items, nil
}

// ItemFromOrderLineItem fetches the associated item for a line-item in an order
func (s *Service) ItemFromOrderLineItem(ctx context.Context, lineItem *ent.OrderLineItem) (*ent.Item, error) {
	client := s.client(ctx)
	item, err := lineItem.Edges.ItemOrErr()
	if err != nil {
		if ent.IsNotLoaded(err) {
			return client.OrderLineItem.QueryItem(lineItem).Only(ctx)
		}
		return nil, err
	}
	return item, nil
}

// NewItem creates a new item for purchase in the shop
func (s *Service) NewItem(ctx context.Context, name string, costCents int) (*ent.Item, error) {
	client := s.client(ctx)
	item, err := client.Item.Create().SetName(name).SetCostCents(costCents).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to save item: %w", err)
	}

	return item, nil
}

// NewCart creates a new cart
func (s *Service) NewCart(ctx context.Context) (*ent.Cart, error) {
	client := s.client(ctx)
	cart, err := client.Cart.Create().Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to save cart: %w", err)
	}

	return cart, nil
}

// GetCart fetches a cart given its ID
func (s *Service) GetCart(ctx context.Context, id uuid.UUID) (*ent.Cart, error) {
	client := s.client(ctx)
	cart, err := client.Cart.Query().
		Where(cart.ID(id)).
		WithLineItems(func(q *ent.LineItemQuery) {
			q.WithItem()
		}).
		Only(ctx)

	if err != nil {
		return nil, fmt.Errorf("unable to load cart: %w", err)
	}

	return cart, nil
}

// AddItemToCart adds a given quantity of an item to a cart. If there is a
// preexisting line item, it is replaces with a new line item that sums their
// quantities
func (s *Service) AddItemToCart(
	ctx context.Context,
	itemID uuid.UUID,
	quantity int,
	cartID uuid.UUID) error {

	return s.withTx(ctx, func(ctx context.Context) error {
		client := s.client(ctx)
		oldLineItem, err := client.LineItem.
			Query().
			WithItem().
			Where(
				lineitem.HasCartWith(cart.ID(cartID)),
				lineitem.HasItemWith(item.ID(itemID)),
			).Only(ctx)

		if err != nil && !ent.IsNotFound(err) {
			return fmt.Errorf("unable to fetch line item: %w", err)
		} else if err == nil {
			// found preexisting line item of same item:
			// add to old quantity and delete old line item
			quantity += oldLineItem.Quantity

			err := client.LineItem.DeleteOne(oldLineItem).Exec(ctx)
			if err != nil {
				return fmt.Errorf("unable to delete old line item: %w", err)
			}
		}

		_, err = client.LineItem.
			Create().
			SetCartID(cartID).
			SetItemID(itemID).
			SetQuantity(quantity).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("unable to create new line item: %w", err)
		}

		return nil
	})
}

// CheckoutCart creates an order given an existing cart, returning the order
// annotated with total cost and payment client secret
func (s *Service) CheckoutCart(ctx context.Context, cartID uuid.UUID) (*model.Order, error) {
	var res *model.Order
	err := s.withTx(ctx, func(ctx context.Context) error {
		client := s.client(ctx)
		cart, err := s.GetCart(ctx, cartID)
		if err != nil {
			return fmt.Errorf("failed to load cart: %w", err)
		}

		// calculate total cost
		totalCostCents := 0
		for _, lineItem := range cart.Edges.LineItems {
			totalCostCents += lineItem.Quantity * lineItem.Edges.Item.CostCents
		}

		// create payment intent
		pi, err := s.paymentsService.NewPaymentIntent(ctx, totalCostCents)
		if err != nil {
			return fmt.Errorf("failed to create payment intent: %w", err)
		}

		// create new order
		order, err := client.Order.Create().SetPaymentIntentID(pi.ID).Save(ctx)
		if err != nil {
			return fmt.Errorf("failed to create order: %w", err)
		}

		// populate new order with line items
		var orderLineItems []*ent.OrderLineItem
		for _, lineItem := range cart.Edges.LineItems {
			orderLineItem, err := client.OrderLineItem.Create().
				SetItem(lineItem.Edges.Item).
				SetOrder(order).
				SetOriginalLineItem(lineItem).
				SetQuantity(lineItem.Quantity).
				SetUnitCostCents(lineItem.Edges.Item.CostCents).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to create order line item: %w", err)
			}

			orderLineItems = append(orderLineItems, orderLineItem)
		}

		res = &model.Order{
			Order:          order,
			TotalCostCents: totalCostCents,
			ClientSecret:   pi.ClientSecret,
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, err
}
