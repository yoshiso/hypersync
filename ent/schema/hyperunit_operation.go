package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type HyperunitOperation struct {
    ent.Schema
}

// Fields of the User.
func (HyperunitOperation) Fields() []ent.Field {
    return []ent.Field{
        field.String("address"),
        field.String("operation_id"),
        field.String("source_chain"),
        field.String("source_amount"),
        field.String("source_address"),
        field.String("source_tx_hash"),
        field.String("destination_tx_hash"),
        field.String("destination_fee_amount"),
        field.String("destination_chain"),
        field.String("destination_address"),
        field.String("sweep_fee_amount"),
        field.Time("op_created_at"),
        field.Time("broadcast_at"),
        field.Time("state_updated_at"),
    }
}

// Edges of the User.
func (HyperunitOperation) Edges() []ent.Edge {
    return nil
}

func (HyperunitOperation) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "operation_id").
            Unique(),
    }
}