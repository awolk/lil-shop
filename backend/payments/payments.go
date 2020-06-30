package payments

import (
	"context"
	"fmt"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

// Service abstracts payment collection
type Service struct {
	client *client.API
}

// PaymentIntent is a planned payment through Stripe
type PaymentIntent struct {
	ID           string
	ClientSecret string
}

func mapPaymentIntent(pi *stripe.PaymentIntent) *PaymentIntent {
	return &PaymentIntent{
		ID:           pi.ID,
		ClientSecret: pi.ClientSecret,
	}
}

// New constructs a new payments Service
func New(stripePrivateKey string) *Service {
	client := client.New(stripePrivateKey, nil)

	return &Service{
		client,
	}
}

// NewPaymentIntent creates a new payment intent with stripe
func (s *Service) NewPaymentIntent(ctx context.Context, costCents int) (*PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Params:   stripe.Params{Context: ctx},
		Amount:   stripe.Int64(int64(costCents)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	pi, err := s.client.PaymentIntents.New(params)
	if err != nil {
		return nil, fmt.Errorf("failed to create payment intent: %w", err)
	}

	return mapPaymentIntent(pi), nil
}

// UpdatePaymentIntent changes the cost of a given payment intent
func (s *Service) UpdatePaymentIntent(ctx context.Context, id string, costCents int) (*PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Params: stripe.Params{Context: ctx},
		Amount: stripe.Int64(int64(costCents)),
	}

	pi, err := s.client.PaymentIntents.Update(id, params)
	if err != nil {
		return nil, fmt.Errorf("failed to update payment intent: %w", err)
	}

	return mapPaymentIntent(pi), nil
}

// GetPaymentIntent fetches a payment intent from Stripe given its ID
func (s *Service) GetPaymentIntent(ctx context.Context, id string) (*PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Params: stripe.Params{Context: ctx},
	}
	pi, err := s.client.PaymentIntents.Get(id, params)
	if err != nil {
		return nil, fmt.Errorf("failed to find payment intent: %w", err)
	}

	return mapPaymentIntent(pi), nil
}
