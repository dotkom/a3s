package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/dotkom/a3s/graph/generated"
	"github.com/dotkom/a3s/graph/model"
)

func (r *queryResolver) Events(ctx context.Context, limit int, offset int) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Event(ctx context.Context, id int) (*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UpcomingEvents(ctx context.Context, eventType *model.EventType) ([]*model.Event, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
