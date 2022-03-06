package repository

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEventOrganizerRepository_Create(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	expected, err := context.EventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	actual, err := context.EventOrganizerRepository.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventOrganizerRepository_Get(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	_, err := context.EventOrganizerRepository.Get(0)
	assert.Errorf(t, err, "non-existent organizer should be error")
	expected, err := context.EventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	actual, err := context.EventOrganizerRepository.Get(expected.ID)
	assert.NoErrorf(t, err, "error getting event organizer")
	assert.Equal(t, expected.ID, actual.ID)
}

func TestEventOrganizerRepository_Count(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	_, err = context.EventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = context.EventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
}

func TestEventOrganizerRepository_All(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	_, err = context.EventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = context.EventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	organizers, err := context.EventOrganizerRepository.All()
	assert.NoErrorf(t, err, "failed to get all event organizers")
	assert.Equal(t, len(organizers), 1)
}

func TestEventOrganizerRepository_Delete(t *testing.T) {
	cleanup, context := SetupTest(t)
	defer cleanup(t)
	count, err := context.EventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
	err = context.EventOrganizerRepository.Delete(0)
	assert.Errorf(t, err, "error deleting event organizer")
	deletable, err := context.EventOrganizerRepository.Create("Fagkom", "fagkom@online.ntnu.no")
	assert.NoErrorf(t, err, "error creating event organizer")
	count, err = context.EventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 1)
	err = context.EventOrganizerRepository.Delete(deletable.ID)
	assert.NoErrorf(t, err, "error deleting event organizer")
	count, err = context.EventOrganizerRepository.Count()
	assert.NoErrorf(t, err, "failed to count event organizers")
	assert.Equal(t, count, 0)
}
