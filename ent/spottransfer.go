// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/spottransfer"
)

// SpotTransfer is the model entity for the SpotTransfer schema.
type SpotTransfer struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// User holds the value of the "user" field.
	User string `json:"user,omitempty"`
	// Destination holds the value of the "destination" field.
	Destination string `json:"destination,omitempty"`
	// Token holds the value of the "token" field.
	Token string `json:"token,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount string `json:"amount,omitempty"`
	// Fee holds the value of the "fee" field.
	Fee string `json:"fee,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time,omitempty"`
	// Address holds the value of the "address" field.
	Address      string `json:"address,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SpotTransfer) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case spottransfer.FieldID, spottransfer.FieldTime:
			values[i] = new(sql.NullInt64)
		case spottransfer.FieldUser, spottransfer.FieldDestination, spottransfer.FieldToken, spottransfer.FieldAmount, spottransfer.FieldFee, spottransfer.FieldAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SpotTransfer fields.
func (st *SpotTransfer) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case spottransfer.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			st.ID = int(value.Int64)
		case spottransfer.FieldUser:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user", values[i])
			} else if value.Valid {
				st.User = value.String
			}
		case spottransfer.FieldDestination:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination", values[i])
			} else if value.Valid {
				st.Destination = value.String
			}
		case spottransfer.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				st.Token = value.String
			}
		case spottransfer.FieldAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				st.Amount = value.String
			}
		case spottransfer.FieldFee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value.Valid {
				st.Fee = value.String
			}
		case spottransfer.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				st.Time = value.Int64
			}
		case spottransfer.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				st.Address = value.String
			}
		default:
			st.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the SpotTransfer.
// This includes values selected through modifiers, order, etc.
func (st *SpotTransfer) Value(name string) (ent.Value, error) {
	return st.selectValues.Get(name)
}

// Update returns a builder for updating this SpotTransfer.
// Note that you need to call SpotTransfer.Unwrap() before calling this method if this SpotTransfer
// was returned from a transaction, and the transaction was committed or rolled back.
func (st *SpotTransfer) Update() *SpotTransferUpdateOne {
	return NewSpotTransferClient(st.config).UpdateOne(st)
}

// Unwrap unwraps the SpotTransfer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (st *SpotTransfer) Unwrap() *SpotTransfer {
	_tx, ok := st.config.driver.(*txDriver)
	if !ok {
		panic("ent: SpotTransfer is not a transactional entity")
	}
	st.config.driver = _tx.drv
	return st
}

// String implements the fmt.Stringer.
func (st *SpotTransfer) String() string {
	var builder strings.Builder
	builder.WriteString("SpotTransfer(")
	builder.WriteString(fmt.Sprintf("id=%v, ", st.ID))
	builder.WriteString("user=")
	builder.WriteString(st.User)
	builder.WriteString(", ")
	builder.WriteString("destination=")
	builder.WriteString(st.Destination)
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(st.Token)
	builder.WriteString(", ")
	builder.WriteString("amount=")
	builder.WriteString(st.Amount)
	builder.WriteString(", ")
	builder.WriteString("fee=")
	builder.WriteString(st.Fee)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", st.Time))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(st.Address)
	builder.WriteByte(')')
	return builder.String()
}

// SpotTransfers is a parsable slice of SpotTransfer.
type SpotTransfers []*SpotTransfer
