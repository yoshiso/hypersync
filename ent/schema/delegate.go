package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type Delegate struct {
    ent.Schema
}

// Fields of the User.
func (Delegate) Fields() []ent.Field {
    return []ent.Field{
        field.String("validator"),
        field.String("amount"),
        field.Bool("is_undelegate"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (Delegate) Edges() []ent.Edge {
    return nil
}

func (Delegate) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "validator").
            Unique(),
    }
}