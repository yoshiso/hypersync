package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type Withdraw struct {
    ent.Schema
}

// Fields of the User.
func (Withdraw) Fields() []ent.Field {
    return []ent.Field{
        field.String("usdc"),
        field.Int64("nonce"),
        field.String("fee"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (Withdraw) Edges() []ent.Edge {
    return nil
}

func (Withdraw) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "nonce").
            Unique(),
    }
}