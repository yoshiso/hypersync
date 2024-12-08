// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/rewardsclaim"
)

// RewardsClaim is the model entity for the RewardsClaim schema.
type RewardsClaim struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount string `json:"amount,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time,omitempty"`
	// Address holds the value of the "address" field.
	Address      string `json:"address,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RewardsClaim) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case rewardsclaim.FieldID, rewardsclaim.FieldTime:
			values[i] = new(sql.NullInt64)
		case rewardsclaim.FieldAmount, rewardsclaim.FieldAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RewardsClaim fields.
func (rc *RewardsClaim) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case rewardsclaim.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rc.ID = int(value.Int64)
		case rewardsclaim.FieldAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				rc.Amount = value.String
			}
		case rewardsclaim.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				rc.Time = value.Int64
			}
		case rewardsclaim.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				rc.Address = value.String
			}
		default:
			rc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RewardsClaim.
// This includes values selected through modifiers, order, etc.
func (rc *RewardsClaim) Value(name string) (ent.Value, error) {
	return rc.selectValues.Get(name)
}

// Update returns a builder for updating this RewardsClaim.
// Note that you need to call RewardsClaim.Unwrap() before calling this method if this RewardsClaim
// was returned from a transaction, and the transaction was committed or rolled back.
func (rc *RewardsClaim) Update() *RewardsClaimUpdateOne {
	return NewRewardsClaimClient(rc.config).UpdateOne(rc)
}

// Unwrap unwraps the RewardsClaim entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rc *RewardsClaim) Unwrap() *RewardsClaim {
	_tx, ok := rc.config.driver.(*txDriver)
	if !ok {
		panic("ent: RewardsClaim is not a transactional entity")
	}
	rc.config.driver = _tx.drv
	return rc
}

// String implements the fmt.Stringer.
func (rc *RewardsClaim) String() string {
	var builder strings.Builder
	builder.WriteString("RewardsClaim(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rc.ID))
	builder.WriteString("amount=")
	builder.WriteString(rc.Amount)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", rc.Time))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(rc.Address)
	builder.WriteByte(')')
	return builder.String()
}

// RewardsClaims is a parsable slice of RewardsClaim.
type RewardsClaims []*RewardsClaim
