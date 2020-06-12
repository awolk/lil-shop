package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/mattn/go-sqlite3"

	"github.com/awolk/lil-shop/backend/ent"
	"github.com/awolk/lil-shop/backend/graph"
	"github.com/awolk/lil-shop/backend/graph/generated"
	"github.com/awolk/lil-shop/backend/service"
)

const defaultPort = "3000"

func main() {
	// load configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
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

	// construct service
	service := service.New(client)

	_, err = service.NewItem(context.Background(), "Test", 199)
	if err != nil {
		log.Fatal(err)
	}
	_, err = service.NewItem(context.Background(), "Test 2", 599)
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

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
