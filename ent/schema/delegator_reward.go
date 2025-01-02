package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type DelegatorReward struct {
    ent.Schema
}

// Fields of the User.
func (DelegatorReward) Fields() []ent.Field {
    return []ent.Field{
        field.String("source"),
        field.String("total_amount"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (DelegatorReward) Edges() []ent.Edge {
    return nil
}

func (DelegatorReward) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "source").
            Unique(),
    }
}