package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Registration holds the schema definition for the Registration entity.
type Registration struct {
	ent.Schema
}

// Fields of the Registration.
func (Registration) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_time"),
		field.Time("end_time"),
		field.Time("unattend_deadline"),
	}
}

// Edges of the Registration.
func (Registration) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("attendees", Attendee.Type),
		edge.To("extras", Extra.Type),
		edge.To("rule_bundles", RuleBundle.Type),
	}
}
