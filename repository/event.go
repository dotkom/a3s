package repository

import (
	"context"
	"github.com/dotkom/a3s/ent"
	"github.com/dotkom/a3s/graph/model"
	"time"
)

type EventRepository struct {
	Client *ent.Client
}

func (r *EventRepository) Create(
	title string,
	startTime time.Time,
	endTime time.Time,
	organizer int,
	price int,
	eventType model.EventType,
	location *string,
	longDescription *string,
	shortDescription *string,
	imageUrl *string,
	companies []int,
) (*ent.Event, error) {
	ctx := context.Background()
	event, err := r.Client.Event.Create().
		SetTitle(title).
		SetStartTime(startTime).
		SetEndTime(endTime).
		SetOrganizerID(organizer).
		SetPrice(price).
		SetEventType(eventType.String()).
		SetNillableLocation(location).
		SetNillableLongDescription(longDescription).
		SetNillableShortDescription(shortDescription).
		SetNillableImageURL(imageUrl).
		AddCompanyIDs(companies...).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return event, nil
}
