package resolvers

import (
	"context"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/graph/model"
)

func transform(eo *ent.EventOrganizer) *model.EventOrganizer {
	return &model.EventOrganizer{
		ID:    eo.ID,
		Name:  eo.Name,
		Email: eo.Email,
	}
}

func EventOrganizer(r *queryResolver, ctx context.Context, id int) (*model.EventOrganizer, error) {
	eventOrganizer, err := r.EventOrganizerRepository.Get(id)
	if err != nil {
		return nil, err
	}
	return transform(eventOrganizer), nil
}
