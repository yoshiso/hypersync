package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type Fill struct {
    ent.Schema
}

// Fields of the User.
func (Fill) Fields() []ent.Field {
    return []ent.Field{
        field.String("coin"),
        field.String("address"),
        field.String("px"),
        field.String("sz"),
        field.String("side"),
        field.Int64("time"),
        field.String("start_position"),
        field.String("dir"),
        field.String("hash"),
        field.Bool("crossed"),
        field.String("fee"),
        field.Int64("oid"),
        field.Int64("tid").Unique(),
        field.String("fee_token"),
        field.String("builder_fee").Optional(),
    }
}

// Edges of the User.
func (Fill) Edges() []ent.Edge {
    return nil
}