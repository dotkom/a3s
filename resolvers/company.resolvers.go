package resolvers

import (
	"context"
	"github.com/dotkom/a3s/ent/company"
	"github.com/dotkom/a3s/graph/model"
)

func Companies(r *queryResolver, ctx context.Context) ([]*model.Company, error) {
	background := context.Background()
	all, err := r.Client.Company.Query().All(background)
	if err != nil {
		return nil, err
	}
	var list []*model.Company
	for _, result := range all {
		record := model.Company{
			ID:               result.ID,
			Name:             result.Name,
			ShortDescription: result.ShortDescription,
			LongDescription:  result.LongDescription,
			ImageURL:         result.ImageURL,
			Site:             &result.Site,
			ContactEmail:     &result.ContactEmail,
			ContactTel:       &result.ContactTel,
		}
		list = append(list, &record)
	}
	return list, nil
}

func Company(r *queryResolver, ctx context.Context, name string) (*model.Company, error) {
	background := context.Background()
	result, err := r.Client.Company.Query().Where(company.NameEQ(name)).Only(background)
	if err != nil {
		return nil, err
	}
	return &model.Company{
		ID:               result.ID,
		Name:             result.Name,
		ShortDescription: result.ShortDescription,
		LongDescription:  result.LongDescription,
		ImageURL:         result.ImageURL,
		Site:             &result.Site,
		ContactEmail:     &result.ContactEmail,
		ContactTel:       &result.ContactTel,
	}, nil
}
