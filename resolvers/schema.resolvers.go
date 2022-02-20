package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/dotkom/a3s/graph/generated"
	"github.com/dotkom/a3s/graph/model"
)

func (r *queryResolver) Companies(ctx context.Context) ([]*model.Company, error) {
	return Companies(r, ctx)
}

func (r *queryResolver) Company(ctx context.Context, name string) (*model.Company, error) {
	return Company(r, ctx, name)
}

func (r *queryResolver) EventOrganizers(ctx context.Context) ([]*model.EventOrganizer, error) {
	return EventOrganizers(r, ctx)
}

func (r *queryResolver) EventOrganizer(ctx context.Context, name string) (*model.EventOrganizer, error) {
	return EventOrganizer(r, ctx, name)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
