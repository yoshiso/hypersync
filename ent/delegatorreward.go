// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/delegatorreward"
)

// DelegatorReward is the model entity for the DelegatorReward schema.
type DelegatorReward struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// TotalAmount holds the value of the "total_amount" field.
	TotalAmount string `json:"total_amount,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time,omitempty"`
	// Address holds the value of the "address" field.
	Address      string `json:"address,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DelegatorReward) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case delegatorreward.FieldID, delegatorreward.FieldTime:
			values[i] = new(sql.NullInt64)
		case delegatorreward.FieldSource, delegatorreward.FieldTotalAmount, delegatorreward.FieldAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DelegatorReward fields.
func (dr *DelegatorReward) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case delegatorreward.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			dr.ID = int(value.Int64)
		case delegatorreward.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				dr.Source = value.String
			}
		case delegatorreward.FieldTotalAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field total_amount", values[i])
			} else if value.Valid {
				dr.TotalAmount = value.String
			}
		case delegatorreward.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				dr.Time = value.Int64
			}
		case delegatorreward.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				dr.Address = value.String
			}
		default:
			dr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the DelegatorReward.
// This includes values selected through modifiers, order, etc.
func (dr *DelegatorReward) Value(name string) (ent.Value, error) {
	return dr.selectValues.Get(name)
}

// Update returns a builder for updating this DelegatorReward.
// Note that you need to call DelegatorReward.Unwrap() before calling this method if this DelegatorReward
// was returned from a transaction, and the transaction was committed or rolled back.
func (dr *DelegatorReward) Update() *DelegatorRewardUpdateOne {
	return NewDelegatorRewardClient(dr.config).UpdateOne(dr)
}

// Unwrap unwraps the DelegatorReward entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (dr *DelegatorReward) Unwrap() *DelegatorReward {
	_tx, ok := dr.config.driver.(*txDriver)
	if !ok {
		panic("ent: DelegatorReward is not a transactional entity")
	}
	dr.config.driver = _tx.drv
	return dr
}

// String implements the fmt.Stringer.
func (dr *DelegatorReward) String() string {
	var builder strings.Builder
	builder.WriteString("DelegatorReward(")
	builder.WriteString(fmt.Sprintf("id=%v, ", dr.ID))
	builder.WriteString("source=")
	builder.WriteString(dr.Source)
	builder.WriteString(", ")
	builder.WriteString("total_amount=")
	builder.WriteString(dr.TotalAmount)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", dr.Time))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(dr.Address)
	builder.WriteByte(')')
	return builder.String()
}

// DelegatorRewards is a parsable slice of DelegatorReward.
type DelegatorRewards []*DelegatorReward
