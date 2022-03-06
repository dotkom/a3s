package repository

import (
	"github.com/dotkom/a3s/graph/model"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEventRepository_Create(t *testing.T) {
	cleanup, eventOrganizerRepository, eventRepository := SetupTest(t)
	defer cleanup(t)
	count, err := eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count events")
	assert.Equal(t, count, 0)
	eventOrganizer, err := eventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	expected, err := eventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	actual, err := eventRepository.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventRepository_Get(t *testing.T) {
	cleanup, eventOrganizerRepository, eventRepository := SetupTest(t)
	defer cleanup(t)
	count, err := eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count events")
	assert.Equal(t, count, 0)
	eventOrganizer, err := eventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	expected, err := eventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	actual, err := eventRepository.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventRepository_Count(t *testing.T) {
	cleanup, eventOrganizerRepository, eventRepository := SetupTest(t)
	defer cleanup(t)
	count, err := eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	eventOrganizer, err := eventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	_, err = eventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	count, err = eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
}

func TestEventRepository_All(t *testing.T) {
	cleanup, eventOrganizerRepository, eventRepository := SetupTest(t)
	defer cleanup(t)
	count, err := eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	eventOrganizer, err := eventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	_, err = eventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	count, err = eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	organizers, err := eventRepository.All()
	assert.NoErrorf(t, err, "failed to get all event organizers")
	assert.Equal(t, len(organizers), 1)
}

func TestEventRepository_Delete(t *testing.T) {
	cleanup, eventOrganizerRepository, eventRepository := SetupTest(t)
	defer cleanup(t)
	count, err := eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	err = eventRepository.Delete(0)
	assert.Errorf(t, err, "error deleting event")
	eventOrganizer, err := eventOrganizerRepository.Create("name", "email@example.com")
	assert.NoErrorf(t, err, "error creating event organizer")
	deletable, err := eventRepository.Create("title", time.Now(), time.Now(), eventOrganizer.ID, 0, model.EventTypeSocial, nil, nil, nil, nil, []int{})
	assert.NoErrorf(t, err, "error creating event")
	count, err = eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	err = eventRepository.Delete(deletable.ID)
	assert.NoErrorf(t, err, "error deleting event organizer")
	count, err = eventRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
}
