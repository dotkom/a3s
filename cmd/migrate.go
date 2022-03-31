package main

import (
	"context"
	"github.com/dotkom/a3s"
	"github.com/dotkom/a3s/ent/migrate"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	env := a3s.GetEnvironment()
	client := a3s.CreateEntClient(env)
	defer client.Close()

	ctx := context.Background()
	// Dump migration changes to stdout.
	if err := client.Debug().Schema.Create(ctx, migrate.WithDropColumn(true), migrate.WithDropIndex(true)); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
