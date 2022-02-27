package schema

import (
	"entgo.io/ent"
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
		field.String("location"),
		field.String("short_description"),
		field.String("long_description"),
		field.String("image_url"),
		field.Int("price"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	// TODO: Relation to:
	// - EventOrganizer
	// - list of Company entities
	// - Registration
	// - list of event rules
	// - EventType
	return nil
}
