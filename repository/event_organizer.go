package repository

import (
	"context"
	"github.com/dotkom/a3s/ent"
)

type EventOrganizerRepository struct {
	Client *ent.Client
}

func (r *EventOrganizerRepository) Create(name string, email string) (*ent.EventOrganizer, error) {
	ctx := context.Background()
	eventOrganizer, err := r.Client.EventOrganizer.Create().
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return eventOrganizer, nil
}

func (r *EventOrganizerRepository) Count() (int, error) {
	ctx := context.Background()
	count, err := r.Client.EventOrganizer.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *EventOrganizerRepository) All() ([]*ent.EventOrganizer, error) {
	ctx := context.Background()
	eventOrganizers, err := r.Client.EventOrganizer.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return eventOrganizers, nil
}

func (r *EventOrganizerRepository) Get(id int) (*ent.EventOrganizer, error) {
	ctx := context.Background()
	eventOrganizer, err := r.Client.EventOrganizer.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return eventOrganizer, nil
}

func (r *EventOrganizerRepository) Delete(id int) error {
	ctx := context.Background()
	err := r.Client.EventOrganizer.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
