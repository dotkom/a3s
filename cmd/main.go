package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dotkom/a3s"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	env := a3s.GetEnvironment()
	client := a3s.CreateEntClient(env)
	defer client.Close()
	schema := a3s.CreateGraphQLSchema(client)
	srv := handler.NewDefaultServer(schema)

	server := chi.NewRouter()
	server.Post("/", srv.ServeHTTP)
	server.Get("/playground", playground.Handler("GraphQL Playground", "/"))

	log.Println("started GraphQL server at http://localhost:8000/playground")
	err := http.ListenAndServe(":8000", server)
	if err != nil {
		log.Fatalf("failed to start server %s\n", err)
	}
}
