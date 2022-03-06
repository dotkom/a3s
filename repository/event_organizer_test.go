package repository

import (
	"context"
	"fmt"
	"github.com/dotkom/a3s/ent"
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
	count, err := r.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	expected, err := r.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	actual, err := r.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventOrganizerRepository_Get(t *testing.T) {
	cleanup, r := setupTest(t)
	defer cleanup(t)
	_, err := r.Get(0)
	assert.Errorf(t, err, "non-existent organizer should be error")
	expected, err := r.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	actual, err := r.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventOrganizerRepository_Count(t *testing.T) {
	cleanup, r := setupTest(t)
	defer cleanup(t)
	count, err := r.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	_, err = r.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = r.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
}

func TestEventOrganizerRepository_All(t *testing.T) {
	cleanup, r := setupTest(t)
	defer cleanup(t)
	count, err := r.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	_, err = r.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = r.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	organizers, err := r.All()
	assert.NoErrorf(t, err, "failed to get all event organizers")
	assert.Equal(t, len(organizers), 1)
}

func TestEventOrganizerRepository_Delete(t *testing.T) {
	cleanup, r := setupTest(t)
	defer cleanup(t)
	count, err := r.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	deletable, err := r.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = r.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	err = r.Delete(deletable.ID)
	assert.NoErrorf(t, err, "error deleting event organizer")
	count, err = r.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
}
