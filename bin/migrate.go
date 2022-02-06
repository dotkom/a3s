package main

import (
	"context"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/ent/migrate"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=admin dbname=a3s password=admin sslmode=disable")
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
