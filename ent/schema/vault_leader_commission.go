package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type VaultLeaderCommission struct {
    ent.Schema
}

// Fields of the User.
func (VaultLeaderCommission) Fields() []ent.Field {
    return []ent.Field{
        field.String("user"),
        field.String("usdc"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (VaultLeaderCommission) Edges() []ent.Edge {
    return nil
}


func (VaultLeaderCommission) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "user").
            Unique(),
    }
}