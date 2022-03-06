package repository

import (
	"context"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/graph/model"
)

type EventOrganizerRepository struct {
	Client *ent.Client
}

func transform(eo *ent.EventOrganizer) *model.EventOrganizer {
	return &model.EventOrganizer{
		ID:    eo.ID,
		Name:  eo.Name,
		Email: eo.Email,
	}
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
	return transform(eventOrganizer), nil
}

func (r *EventOrganizerRepository) Count() (int, error) {
	ctx := context.Background()
	count, err := r.Client.EventOrganizer.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *EventOrganizerRepository) All() ([]*model.EventOrganizer, error) {
	ctx := context.Background()
	eventOrganizers, err := r.Client.EventOrganizer.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]*model.EventOrganizer, len(eventOrganizers))
	for i, eo := range eventOrganizers {
		result[i] = transform(eo)
	}
	return result, nil
}

func (r *EventOrganizerRepository) Get(id int) (*model.EventOrganizer, error) {
	ctx := context.Background()
	eventOrganizer, err := r.Client.EventOrganizer.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return transform(eventOrganizer), nil
}

func (r *EventOrganizerRepository) Delete(id int) error {
	ctx := context.Background()
	err := r.Client.EventOrganizer.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
