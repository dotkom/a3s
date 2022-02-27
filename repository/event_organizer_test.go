package repository

import (
	"context"
	"fmt"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/ent/eventorganizer"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"testing"
)

func setupTest(t *testing.T) (func(*testing.T), EventOrganizerRepository) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	query := fmt.Sprintf("host=localhost port=5432 sslmode=disable user=%s password=%s dbname=%s", user, password, db)
	client, err := ent.Open("postgres", query)
	if err != nil {
		log.Fatalf("failed connecting to postgres: %v", err)
	}
	r := EventOrganizerRepository{Client: client}
	return func(t *testing.T) {
		ctx := context.Background()
		client.EventOrganizer.Delete().ExecX(ctx)
		err := client.Close()
		if err != nil {
			log.Fatalf("failed closing postgres: %v", err)
		}
	}, r
}

func TestEventOrganizerRepository_Create(t *testing.T) {
	cleanup, r := setupTest(t)
	defer cleanup(t)
	ctx := context.Background()
	count, _ := r.Client.EventOrganizer.Query().Count(ctx)
	assert.Equal(t, count, 0)
	expected, err := r.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	actual, err := r.Client.EventOrganizer.Query().Where(eventorganizer.Name("Fagkom")).Only(ctx)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}
