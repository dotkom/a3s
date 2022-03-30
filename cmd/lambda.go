package main

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/graph/generated"
	"github.com/dotkom/a3s/repository"
	"github.com/dotkom/a3s/resolvers"
	"github.com/go-chi/chi/v5"
	"log"
	"os"
)

var chiLambda *chiadapter.ChiLambda

func init() {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	query := fmt.Sprintf("host=localhost port=5432 sslmode=disable user=%s password=%s dbname=%s", user, password, db)
	client, err := ent.Open("postgres", query)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			log.Fatalf("failed closing postgres: %v", err)
		}
	}(client)

	server := chi.NewRouter()
	resolver := resolvers.Resolver{
		EventOrganizerRepository: &repository.EventOrganizerRepository{Client: client},
		EventRepository:          &repository.EventRepository{Client: client},
	}
	config := generated.Config{Resolvers: &resolver}
	schema := generated.NewExecutableSchema(config)
	srv := handler.NewDefaultServer(schema)

	server.Post("/", srv.ServeHTTP)
	chiLambda = chiadapter.New(server)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
