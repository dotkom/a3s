package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RuleBundle holds the schema definition for the RuleBundle entity.
type RuleBundle struct {
	ent.Schema
}

// Fields of the RuleBundle.
func (RuleBundle) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the RuleBundle.
func (RuleBundle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("registrations", Registration.Type).Ref("rule_bundles"),
	}
}
