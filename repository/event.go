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

func (r *EventRepository) Count() (int, error) {
	ctx := context.Background()
	count, err := r.Client.Event.Query().Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *EventRepository) All() ([]*ent.Event, error) {
	ctx := context.Background()
	events, err := r.Client.Event.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (r *EventRepository) Get(id int) (*ent.Event, error) {
	ctx := context.Background()
	event, err := r.Client.Event.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (r *EventRepository) Delete(id int) error {
	ctx := context.Background()
	err := r.Client.Event.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}
