package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dotkom/a3s/graph"
	"github.com/dotkom/a3s/graph/generated"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

func main() {
	server := chi.NewRouter()
	config := generated.Config{Resolvers: &graph.Resolver{}}
	schema := generated.NewExecutableSchema(config)
	srv := handler.NewDefaultServer(schema)

	server.Post("/", srv.ServeHTTP)
	server.Get("/playground", playground.Handler("GraphQL Playground", "/"))

	log.Println("started GraphQL server at http://localhost:8000/playground")
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		log.Fatalf("failed to start server %s\n", err)
	}
}
