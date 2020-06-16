package main

import (
	"fmt"
	"os"
)

type config struct {
	stripePrivateKey string
	port             string
}

func loadConfig() (*config, error) {
	stripePrivateKey := os.Getenv("STRIPE_SECRET_KEY")
	if stripePrivateKey == "" {
		return nil, fmt.Errorf("STRIPE_SECRET_KEY unset")
	}

	return &config{
		stripePrivateKey: stripePrivateKey,
		port:             "3000",
	}, nil
}
