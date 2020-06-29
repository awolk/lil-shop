package graph

import "github.com/awolk/lil-shop/backend/service"

//go:generate go run github.com/99designs/gqlgen

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver holds dependencies for the GraphQL resolvers
type Resolver struct {
	Service *service.Service
}
