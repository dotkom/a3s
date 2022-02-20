package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Attendance holds the schema definition for the Attendance entity.
type Attendance struct {
	ent.Schema
}

// Fields of the Attendance.
func (Attendance) Fields() []ent.Field {
	return []ent.Field{
		field.Time("start_time"),
		field.Time("end_time"),
		field.Time("unattend_deadline"),
	}
}

// Edges of the Attendance.
func (Attendance) Edges() []ent.Edge {
	// TODO: Relation to attendees
	return nil
}
