package main

import (
	"fmt"
	"os"
)

type config struct {
	stripePrivateKey     string
	stripeEndpointSecret string
	port                 string
}

func loadConfig() (*config, error) {
	stripePrivateKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripePrivateKey == "" {
		return nil, fmt.Errorf("STRIPE_SECRET_KEY unset")
	}

	stripeEndpointSecret := os.Getenv("STRIPE_ENDPOINT_SECRET")
	if stripeEndpointSecret == "" {
		return nil, fmt.Errorf("STRIPE_ENDPOINT_SECRET unset")
	}

	return &config{
		stripePrivateKey:     stripePrivateKey,
		stripeEndpointSecret: stripeEndpointSecret,
		port:                 "3000",
	}, nil
}
