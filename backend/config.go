package main

import (
	"fmt"
	"os"
)

type config struct {
	stripePrivateKey     string
	stripeEndpointSecret string
	dbURI                string
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

	dbURI := os.Getenv("DB_URI")
	if dbURI == "" {
		return nil, fmt.Errorf("DB_URI unset")
	}

	return &config{
		stripePrivateKey:     stripePrivateKey,
		stripeEndpointSecret: stripeEndpointSecret,
		dbURI:                dbURI,
		port:                 "3000",
	}, nil
}
