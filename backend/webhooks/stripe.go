package webhooks

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/awolk/lil-shop/backend/service"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"
)

// WebhookHandler is an HTTP handler for stripe webhooks
type WebhookHandler struct {
	s              *service.Service
	endpointSecret string
}

// New constructs a new WebhookHandler
func New(s *service.Service, endpointSecret string) *WebhookHandler {
	return &WebhookHandler{
		s:              s,
		endpointSecret: endpointSecret,
	}
}

func (h *WebhookHandler) handlePaymentIntentSucceeded(ctx context.Context, pi *stripe.PaymentIntent) error {
	err := h.s.CompleteOrder(ctx, pi.ID)
	if err != nil {
		return fmt.Errorf("failed to complete order: %v", err)
	}
	return nil
}

func (h *WebhookHandler) handlePaymentIntentFailed(ctx context.Context, pi *stripe.PaymentIntent) error {
	log.Printf("order failed: %v", pi.LastPaymentError)
	return nil
}

func (h *WebhookHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Printf("Error reading request body: %v", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"),
		h.endpointSecret)

	if err != nil {
		log.Printf("Error verifying webhook signature: %v", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}

	ctx := context.Background()

	if event.Type == "payment_intent.succeeded" {
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.handlePaymentIntentSucceeded(ctx, &paymentIntent)
		if err != nil {
			log.Printf("Error handling successful payment intent notification: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else if event.Type == "payment_intent.payment_failed" {
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			log.Printf("Error parsing webhook JSON: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = h.handlePaymentIntentFailed(ctx, &paymentIntent)
		if err != nil {
			log.Printf("Error handling failed payment intent notification: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	log.Printf("Handled %s webhook", event.Type)
	w.WriteHeader(http.StatusOK)
}
