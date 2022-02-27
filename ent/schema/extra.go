package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Extra holds the schema definition for the Extra entity.
type Extra struct {
	ent.Schema
}

// Fields of the Extra.
func (Extra) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Extra.
func (Extra) Edges() []ent.Edge {
	return nil
}
