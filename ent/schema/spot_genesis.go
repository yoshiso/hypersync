package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type SpotGenesis struct {
    ent.Schema
}

// Fields of the User.
func (SpotGenesis) Fields() []ent.Field {
    return []ent.Field{
        field.String("coin"),
        field.String("amount"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (SpotGenesis) Edges() []ent.Edge {
    return nil
}

func (SpotGenesis) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "coin").
            Unique(),
    }
}