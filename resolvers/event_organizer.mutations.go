package resolvers

import (
	"context"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/graph/model"
)

func CreateEventOrganizer(r *mutationResolver, ctx context.Context, data model.CreateEventOrganizerInput) (*model.EventOrganizer, error) {
	eventOrganizer, err := r.EventOrganizerRepository.Create(data.Name, data.Email)
	if err != nil {
		return nil, err
	}
	return r.Query().EventOrganizer(ctx, eventOrganizer.ID)
}

func UpdateEventOrganizerName(r *mutationResolver, ctx context.Context, id int, name string) (*model.EventOrganizer, error) {
	eventOrganizer, err := r.EventOrganizerRepository.Update(id, func(eventOrganizer *ent.EventOrganizerUpdateOne) {
		eventOrganizer.SetName(name)
	})
	if err != nil {
		return nil, err
	}
	return r.Query().EventOrganizer(ctx, eventOrganizer.ID)
}

func UpdateEventOrganizerEmail(r *mutationResolver, ctx context.Context, id int, email string) (*model.EventOrganizer, error) {
	eventOrganizer, err := r.EventOrganizerRepository.Update(id, func(eventOrganizer *ent.EventOrganizerUpdateOne) {
		eventOrganizer.SetEmail(email)
	})
	if err != nil {
		return nil, err
	}
	return r.Query().EventOrganizer(ctx, eventOrganizer.ID)
}
