package main

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/graph/generated"
	"github.com/dotkom/a3s/resolvers"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
)

func main() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	query := fmt.Sprintf("host=localhost port=5432 sslmode=disable user=%s password=%s dbname=%s", user, password, db)
	client, err := ent.Open("postgres", query)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer client.Close()

	server := chi.NewRouter()
	config := generated.Config{Resolvers: &resolvers.Resolver{Client: client}}
	schema := generated.NewExecutableSchema(config)
	srv := handler.NewDefaultServer(schema)

	server.Post("/", srv.ServeHTTP)
	server.Get("/playground", playground.Handler("GraphQL Playground", "/"))

	log.Println("started GraphQL server at http://localhost:8000/playground")
	err = http.ListenAndServe(":8000", server)
	if err != nil {
		log.Fatalf("failed to start server %s\n", err)
	}
}
