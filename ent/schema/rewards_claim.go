package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type RewardsClaim struct {
    ent.Schema
}

// Fields of the User.
func (RewardsClaim) Fields() []ent.Field {
    return []ent.Field{
        field.String("amount"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (RewardsClaim) Edges() []ent.Edge {
    return nil
}