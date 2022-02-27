package resolvers

import (
	"context"
	"github.com/dotkom/a3s/graph/model"
)

func CreateEventOrganizer(r *mutationResolver, ctx context.Context, data model.CreateEventOrganizerInput) (*model.EventOrganizer, error) {
	return r.EventOrganizerRepository.Create(data.Name, data.Email)
}
