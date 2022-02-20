package resolvers

import (
	"context"
	"github.com/dotkom/a3s/ent/eventorganizer"
	"github.com/dotkom/a3s/graph/model"
)

func EventOrganizers(r *queryResolver, ctx context.Context) ([]*model.EventOrganizer, error) {
	background := context.Background()
	all, err := r.Client.EventOrganizer.Query().All(background)
	if err != nil {
		return nil, err
	}
	var list []*model.EventOrganizer
	for _, result := range all {
		record := model.EventOrganizer{
			ID:    result.ID,
			Name:  result.Name,
			Email: result.Email,
		}
		list = append(list, &record)
	}
	return list, nil
}

func EventOrganizer(r *queryResolver, ctx context.Context, name string) (*model.EventOrganizer, error) {
	background := context.Background()
	result, err := r.Client.EventOrganizer.Query().Where(eventorganizer.NameEQ(name)).Only(background)
	if err != nil {
		return nil, err
	}
	return &model.EventOrganizer{
		ID:    result.ID,
		Name:  result.Name,
		Email: result.Email,
	}, nil
}
