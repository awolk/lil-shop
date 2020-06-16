package main

import (
	"context"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"

	"github.com/awolk/lil-shop/backend/ent"
	"github.com/awolk/lil-shop/backend/graph"
	"github.com/awolk/lil-shop/backend/graph/generated"
	"github.com/awolk/lil-shop/backend/payments"
	"github.com/awolk/lil-shop/backend/service"
)

func main() {
	// load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("failed loading configuration: %v", err)
	}

	// connect to database
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// run migrations
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// construct services
	paymentService := payments.New(config.stripePrivateKey)
	service := service.New(client, paymentService)

	_, err = service.NewItem(context.Background(), "Sunglasses", 1099)
	if err != nil {
		log.Fatal(err)
	}
	_, err = service.NewItem(context.Background(), "Apple", 199)
	if err != nil {
		log.Fatal(err)
	}

	// start server
	resolver := &graph.Resolver{
		Service: service,
	}
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/playground", playground.Handler("GraphQL playground", "/graphql"))
	http.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.port)
	log.Fatal(http.ListenAndServe(":"+config.port, nil))
}
