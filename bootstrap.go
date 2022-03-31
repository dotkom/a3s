package a3s

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/graph/generated"
	"github.com/dotkom/a3s/repository"
	"github.com/dotkom/a3s/resolvers"
	"log"
	"os"
)

type SystemEnvironment struct {
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

func GetEnvironment() *SystemEnvironment {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	return &SystemEnvironment{
		DatabaseUser:     user,
		DatabasePassword: password,
		DatabaseName:     db,
	}
}

func CreateEntClient(env *SystemEnvironment) *ent.Client {
	query := fmt.Sprintf("host=localhost port=5432 sslmode=disable user=%s password=%s dbname=%s", env.DatabaseUser, env.DatabasePassword, env.DatabaseName)
	client, err := ent.Open("postgres", query)
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	return client
}

func CreateGraphQLSchema(client *ent.Client) graphql.ExecutableSchema {
	resolver := resolvers.Resolver{
		EventOrganizerRepository: &repository.EventOrganizerRepository{Client: client},
		EventRepository:          &repository.EventRepository{Client: client},
	}
	configuration := generated.Config{Resolvers: &resolver}
	return generated.NewExecutableSchema(configuration)
}
