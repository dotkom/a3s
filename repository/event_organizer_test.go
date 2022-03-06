package repository

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventOrganizerRepository_Create(t *testing.T) {
	cleanup, eventOrganizerRepository, _ := SetupTest(t)
	defer cleanup(t)
	count, err := eventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	expected, err := eventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	actual, err := eventOrganizerRepository.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventOrganizerRepository_Get(t *testing.T) {
	cleanup, eventOrganizerRepository, _ := SetupTest(t)
	defer cleanup(t)
	_, err := eventOrganizerRepository.Get(0)
	assert.Errorf(t, err, "non-existent organizer should be error")
	expected, err := eventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	actual, err := eventOrganizerRepository.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventOrganizerRepository_Count(t *testing.T) {
	cleanup, eventOrganizerRepository, _ := SetupTest(t)
	defer cleanup(t)
	count, err := eventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	_, err = eventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = eventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
}

func TestEventOrganizerRepository_All(t *testing.T) {
	cleanup, eventOrganizerRepository, _ := SetupTest(t)
	defer cleanup(t)
	count, err := eventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	_, err = eventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = eventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	organizers, err := eventOrganizerRepository.All()
	assert.NoErrorf(t, err, "failed to get all event organizers")
	assert.Equal(t, len(organizers), 1)
}

func TestEventOrganizerRepository_Delete(t *testing.T) {
	cleanup, eventOrganizerRepository, _ := SetupTest(t)
	defer cleanup(t)
	count, err := eventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	deletable, err := eventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = eventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	err = eventOrganizerRepository.Delete(deletable.ID)
	assert.NoErrorf(t, err, "error deleting event organizer")
	count, err = eventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
}
