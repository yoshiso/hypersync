package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type InternalTransfer struct {
    ent.Schema
}

// Fields of the User.
func (InternalTransfer) Fields() []ent.Field {
    return []ent.Field{
        field.String("user"),
        field.String("destination"),
        field.String("usdc"),
        field.String("fee"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (InternalTransfer) Edges() []ent.Edge {
    return nil
}

func (InternalTransfer) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "user", "destination").
            Unique(),
    }
}