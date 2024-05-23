// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/MatiF100/Throw-Muffin-API/ent/workout"
)

// Workout is the model entity for the Workout schema.
type Workout struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Intensity holds the value of the "Intensity" field.
	Intensity    int `json:"Intensity,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Workout) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workout.FieldID, workout.FieldIntensity:
			values[i] = new(sql.NullInt64)
		case workout.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Workout fields.
func (w *Workout) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workout.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case workout.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Name", values[i])
			} else if value.Valid {
				w.Name = value.String
			}
		case workout.FieldIntensity:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Intensity", values[i])
			} else if value.Valid {
				w.Intensity = int(value.Int64)
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Workout.
// This includes values selected through modifiers, order, etc.
func (w *Workout) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// Update returns a builder for updating this Workout.
// Note that you need to call Workout.Unwrap() before calling this method if this Workout
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Workout) Update() *WorkoutUpdateOne {
	return NewWorkoutClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Workout entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Workout) Unwrap() *Workout {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Workout is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Workout) String() string {
	var builder strings.Builder
	builder.WriteString("Workout(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("Name=")
	builder.WriteString(w.Name)
	builder.WriteString(", ")
	builder.WriteString("Intensity=")
	builder.WriteString(fmt.Sprintf("%v", w.Intensity))
	builder.WriteByte(')')
	return builder.String()
}

// Workouts is a parsable slice of Workout.
type Workouts []*Workout