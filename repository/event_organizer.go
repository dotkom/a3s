package repository

import (
	"context"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/graph/model"
)

type EventOrganizerRepository struct {
	Client *ent.Client
}

func (r *EventOrganizerRepository) Create(name string, email string) (*model.EventOrganizer, error) {
	ctx := context.Background()
	eventOrganizer, err := r.Client.EventOrganizer.Create().
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return &model.EventOrganizer{
		ID:    eventOrganizer.ID,
		Name:  eventOrganizer.Name,
		Email: eventOrganizer.Email,
	}, nil
}
