package payments

import (
	"context"
	"fmt"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
)

type PaymentsService struct {
	client *client.API
}

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

func New(stripePrivateKey string) *PaymentsService {
	client := client.New(stripePrivateKey, nil)

	return &PaymentsService{
		client,
	}
}

func (s *PaymentsService) NewPaymentIntent(ctx context.Context, costCents int) (*PaymentIntent, error) {
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

func (s *PaymentsService) UpdatePaymentIntent(ctx context.Context, id string, costCents int) (*PaymentIntent, error) {
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

func (s *PaymentsService) GetPaymentIntent(ctx context.Context, id string) (*PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Params: stripe.Params{Context: ctx},
	}
	pi, err := s.client.PaymentIntents.Get(id, params)
	if err != nil {
		return nil, fmt.Errorf("failed to find payment intent: %w", err)
	}

	return mapPaymentIntent(pi), nil
}
