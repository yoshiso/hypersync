package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type Funding struct {
    ent.Schema
}

// Fields of the User.
func (Funding) Fields() []ent.Field {
    return []ent.Field{
        field.Int64("time"),
        field.String("coin"),
        field.String("usdc"),
        field.String("szi"),
        field.String("funding_rate"),
        field.String("address"),
    }
}

// Edges of the User.
func (Funding) Edges() []ent.Edge {
    return nil
}

func (Funding) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "coin").
            Unique(),
    }
}