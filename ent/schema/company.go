package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("short_description"),
		field.String("long_description"),
		field.String("image_url"),
		field.String("site").Optional(),
		field.String("contact_email").Optional(),
		field.String("contact_tel").Optional(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("companies", Event.Type),
	}
}
