package resolvers

import (
	"context"
	"github.com/dotkom/a3s/graph/model"
)

func CreateEventOrganizer(r *mutationResolver, ctx context.Context, data model.CreateEventOrganizerInput) (*model.EventOrganizer, error) {
	eventOrganizer, err := r.EventOrganizerRepository.Create(data.Name, data.Email)
	if err != nil {
		return nil, err
	}
	return r.Query().EventOrganizer(ctx, eventOrganizer.ID)
}
