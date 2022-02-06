package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").Positive().Unique(),
		field.String("name"),
		field.String("lead"),
		field.String("description"),
		field.String("image"),
		field.String("site").Optional(),
		field.String("email").Optional(),
		field.String("telephone").Optional(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return nil
}

func (Company) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
