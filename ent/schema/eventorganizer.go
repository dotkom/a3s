package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// EventOrganizer holds the schema definition for the EventOrganizer entity.
type EventOrganizer struct {
	ent.Schema
}

// Fields of the EventOrganizer.
func (EventOrganizer) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
	}
}

// Edges of the EventOrganizer.
func (EventOrganizer) Edges() []ent.Edge {
	return []ent.Edge{
		//edge.From("event", Event.Type).
		//	Ref("organizer").
		//	Unique(),

	}
}
