package resolvers

import (
	"context"
	"github.com/dotkom/a3s/graph/model"
)

func CreateEvent(r *mutationResolver, ctx context.Context, data model.CreateEventInput) (*model.Event, error) {
	event, err := r.EventRepository.Create(data.Title, data.StartTime, data.EndTime, data.Organizer, data.Price, data.EventType, data.Location, data.LongDescription, data.ShortDescription, data.ImageURL, data.Companies)
	if err != nil {
		return nil, err
	}
	return r.Query().Event(ctx, event.ID)
}
