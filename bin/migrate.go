package main

import (
	"context"
	"fmt"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/ent/migrate"
	_ "github.com/lib/pq"
	"log"
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
	ctx := context.Background()
	// Dump migration changes to stdout.
	if err := client.Debug().Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
