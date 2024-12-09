package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// User holds the schema definition for the User entity.
type VaultWithdrawal struct {
    ent.Schema
}

// Fields of the User.
func (VaultWithdrawal) Fields() []ent.Field {
    return []ent.Field{
        field.String("vault"),
        field.String("user"),
        field.String("requested_usd"),
        field.String("commission"),
        field.String("closing_cost"),
        field.String("basis"),
        field.String("net_withdrawn_usd"),
        field.Int64("time"),
        field.String("address"),
    }
}

// Edges of the User.
func (VaultWithdrawal) Edges() []ent.Edge {
    return nil
}

func (VaultWithdrawal) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("address", "time", "vault").
            Unique(),
    }
}