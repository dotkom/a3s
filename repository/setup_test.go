package repository

import (
	"context"
	"fmt"
	"github.com/dotkom/a3s/ent"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

type TestContext struct {
	EventOrganizerRepository EventOrganizerRepository
	EventRepository          EventRepository
}

func SetupTest(t *testing.T) (func(*testing.T), TestContext) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	query := fmt.Sprintf("host=localhost port=5432 sslmode=disable user=%s password=%s dbname=%s", user, password, db)
	client, err := ent.Open("postgres", query)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	eventOrganizerRepository := EventOrganizerRepository{Client: client}
	eventRepository := EventRepository{Client: client}

	testContext := TestContext{
		EventRepository:          eventRepository,
		EventOrganizerRepository: eventOrganizerRepository,
	}

	return func(t *testing.T) {
		ctx := context.Background()
		client.Event.Delete().ExecX(ctx)
		client.EventOrganizer.Delete().ExecX(ctx)
		err := client.Close()
		if err != nil {
			log.Fatalf("failed closing postgres: %v", err)
		}
	}, testContext
}
