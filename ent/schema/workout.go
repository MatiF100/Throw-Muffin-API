package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Workout holds the schema definition for the Workout entity.
type Workout struct {
	ent.Schema
}

// Fields of the Workout.
func (Workout) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.Int("Intensity").Positive().Default(1),
	}
}

// Edges of the Workout.
func (Workout) Edges() []ent.Edge {
	return nil
}
