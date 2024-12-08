// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/funding"
)

// Funding is the model entity for the Funding schema.
type Funding struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time,omitempty"`
	// Coin holds the value of the "coin" field.
	Coin string `json:"coin,omitempty"`
	// Usdc holds the value of the "usdc" field.
	Usdc string `json:"usdc,omitempty"`
	// Szi holds the value of the "szi" field.
	Szi string `json:"szi,omitempty"`
	// FundingRate holds the value of the "funding_rate" field.
	FundingRate string `json:"funding_rate,omitempty"`
	// Address holds the value of the "address" field.
	Address      string `json:"address,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Funding) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case funding.FieldID, funding.FieldTime:
			values[i] = new(sql.NullInt64)
		case funding.FieldCoin, funding.FieldUsdc, funding.FieldSzi, funding.FieldFundingRate, funding.FieldAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Funding fields.
func (f *Funding) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case funding.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case funding.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				f.Time = value.Int64
			}
		case funding.FieldCoin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field coin", values[i])
			} else if value.Valid {
				f.Coin = value.String
			}
		case funding.FieldUsdc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field usdc", values[i])
			} else if value.Valid {
				f.Usdc = value.String
			}
		case funding.FieldSzi:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field szi", values[i])
			} else if value.Valid {
				f.Szi = value.String
			}
		case funding.FieldFundingRate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field funding_rate", values[i])
			} else if value.Valid {
				f.FundingRate = value.String
			}
		case funding.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				f.Address = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Funding.
// This includes values selected through modifiers, order, etc.
func (f *Funding) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this Funding.
// Note that you need to call Funding.Unwrap() before calling this method if this Funding
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Funding) Update() *FundingUpdateOne {
	return NewFundingClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Funding entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Funding) Unwrap() *Funding {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Funding is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Funding) String() string {
	var builder strings.Builder
	builder.WriteString("Funding(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", f.Time))
	builder.WriteString(", ")
	builder.WriteString("coin=")
	builder.WriteString(f.Coin)
	builder.WriteString(", ")
	builder.WriteString("usdc=")
	builder.WriteString(f.Usdc)
	builder.WriteString(", ")
	builder.WriteString("szi=")
	builder.WriteString(f.Szi)
	builder.WriteString(", ")
	builder.WriteString("funding_rate=")
	builder.WriteString(f.FundingRate)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(f.Address)
	builder.WriteByte(')')
	return builder.String()
}

// Fundings is a parsable slice of Funding.
type Fundings []*Funding
