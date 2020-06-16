package payments

import (
	"fmt"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type PaymentsService struct {
	client paymentintent.Client
}

func New(stripePrivateKey string) *PaymentsService {
	client := paymentintent.Client{
		B:   stripe.GetBackend(stripe.APIBackend),
		Key: stripePrivateKey,
	}

	return &PaymentsService{
		client,
	}
}

func (s *PaymentsService) NewPaymentIntent(costCents int) (string, error) {
	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(costCents)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}

	pi, err := s.client.New(params)
	if err != nil {
		return "", fmt.Errorf("failed created payment intent: %w", err)
	}

	return pi.ClientSecret, nil
}
