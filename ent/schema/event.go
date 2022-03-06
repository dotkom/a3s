package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
		field.Time("start_time"),
		field.Time("end_time"),
		field.String("location").Optional(),
		field.String("short_description").Optional(),
		field.String("long_description").Optional(),
		field.String("image_url").Optional(),
		field.Int("price"),
		field.String("event_type"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("organizer", EventOrganizer.Type).
			Unique(),
		edge.To("companies", Company.Type),
		edge.To("registration", Registration.Type).
			Unique(),
	}
}
