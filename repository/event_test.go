package repository

import (
	"github.com/dotkom/a3s/graph/model"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventRepository_Create(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count events")
	assert.Equal(t, count, 0)
	eventOrganizer, err := context.EventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	expected, err := context.EventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	actual, err := context.EventRepository.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventRepository_Get(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count events")
	assert.Equal(t, count, 0)
	eventOrganizer, err := context.EventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	expected, err := context.EventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	actual, err := context.EventRepository.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventRepository_Count(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	eventOrganizer, err := context.EventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	_, err = context.EventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	count, err = context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
}

func TestEventRepository_All(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	eventOrganizer, err := context.EventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	_, err = context.EventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	count, err = context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	organizers, err := context.EventRepository.All()
	assert.NoErrorf(t, err, "failed to get all event organizers")
	assert.Equal(t, len(organizers), 1)
}

func TestEventRepository_Delete(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	err = context.EventRepository.Delete(0)
	assert.Errorf(t, err, "error deleting event")
	eventOrganizer, err := context.EventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	deletable, err := context.EventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	count, err = context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	err = context.EventRepository.Delete(deletable.ID)
	assert.NoErrorf(t, err, "error deleting event organizer")
	count, err = context.EventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
}
