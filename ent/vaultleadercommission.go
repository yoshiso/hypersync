// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/vaultleadercommission"
)

// VaultLeaderCommission is the model entity for the VaultLeaderCommission schema.
type VaultLeaderCommission struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
	// Usdc holds the value of the "usdc" field.
	Usdc string `json:"usdc,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time,omitempty"`
	// Address holds the value of the "address" field.
	Address      string `json:"address,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*VaultLeaderCommission) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case vaultleadercommission.FieldID, vaultleadercommission.FieldTime:
			values[i] = new(sql.NullInt64)
		case vaultleadercommission.FieldUser, vaultleadercommission.FieldUsdc, vaultleadercommission.FieldAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the VaultLeaderCommission fields.
func (vlc *VaultLeaderCommission) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case vaultleadercommission.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			vlc.ID = int(value.Int64)
		case vaultleadercommission.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				vlc.User = value.String
			}
		case vaultleadercommission.FieldUsdc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field usdc", values[i])
			} else if value.Valid {
				vlc.Usdc = value.String
			}
		case vaultleadercommission.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				vlc.Time = value.Int64
			}
		case vaultleadercommission.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				vlc.Address = value.String
			}
		default:
			vlc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the VaultLeaderCommission.
// This includes values selected through modifiers, order, etc.
func (vlc *VaultLeaderCommission) Value(name string) (ent.Value, error) {
	return vlc.selectValues.Get(name)
}

// Update returns a builder for updating this VaultLeaderCommission.
// Note that you need to call VaultLeaderCommission.Unwrap() before calling this method if this VaultLeaderCommission
// was returned from a transaction, and the transaction was committed or rolled back.
func (vlc *VaultLeaderCommission) Update() *VaultLeaderCommissionUpdateOne {
	return NewVaultLeaderCommissionClient(vlc.config).UpdateOne(vlc)
}

// Unwrap unwraps the VaultLeaderCommission entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (vlc *VaultLeaderCommission) Unwrap() *VaultLeaderCommission {
	_tx, ok := vlc.config.driver.(*txDriver)
	if !ok {
		panic("ent: VaultLeaderCommission is not a transactional entity")
	}
	vlc.config.driver = _tx.drv
	return vlc
}

// String implements the fmt.Stringer.
func (vlc *VaultLeaderCommission) String() string {
	var builder strings.Builder
	builder.WriteString("VaultLeaderCommission(")
	builder.WriteString(fmt.Sprintf("id=%v, ", vlc.ID))
	builder.WriteString("user=")
	builder.WriteString(vlc.User)
	builder.WriteString(", ")
	builder.WriteString("usdc=")
	builder.WriteString(vlc.Usdc)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", vlc.Time))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(vlc.Address)
	builder.WriteByte(')')
	return builder.String()
}

// VaultLeaderCommissions is a parsable slice of VaultLeaderCommission.
type VaultLeaderCommissions []*VaultLeaderCommission
