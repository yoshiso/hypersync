package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type VaultDelta struct {
    ent.Schema
}

// Fields of the User.
func (VaultDelta) Fields() []ent.Field {
    return []ent.Field{
        field.String("type"),
        field.String("vault"),
        field.String("usdc"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (VaultDelta) Edges() []ent.Edge {
    return nil
}

func (VaultDelta) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "vault").
            Unique(),
    }
}