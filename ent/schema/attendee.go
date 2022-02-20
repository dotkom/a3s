package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Attendee holds the schema definition for the Attendee entity.
type Attendee struct {
	ent.Schema
}

// Fields of the Attendee.
func (Attendee) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"),
		field.Bool("has_paid"),
		field.Bool("is_attending"),
	}
}

// Edges of the Attendee.
func (Attendee) Edges() []ent.Edge {
	return nil
}
