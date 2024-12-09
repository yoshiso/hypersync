package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type SpotTransfer struct {
    ent.Schema
}

// Fields of the User.
func (SpotTransfer) Fields() []ent.Field {
    return []ent.Field{
        field.String("user"),
        field.String("destination"),
        field.String("token"),
        field.String("amount"),
        field.String("fee"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (SpotTransfer) Edges() []ent.Edge {
    return nil
}

func (SpotTransfer) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "token", "user", "destination").
            Unique(),
    }
}