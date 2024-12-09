// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/withdraw"
)

// Withdraw is the model entity for the Withdraw schema.
type Withdraw struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Usdc holds the value of the "usdc" field.
	Usdc string `json:"usdc,omitempty"`
	// Nonce holds the value of the "nonce" field.
	Nonce int64 `json:"nonce,omitempty"`
	// Fee holds the value of the "fee" field.
	Fee string `json:"fee,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time,omitempty"`
	// Address holds the value of the "address" field.
	Address      string `json:"address,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Withdraw) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case withdraw.FieldID, withdraw.FieldNonce, withdraw.FieldTime:
			values[i] = new(sql.NullInt64)
		case withdraw.FieldUsdc, withdraw.FieldFee, withdraw.FieldAddress:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Withdraw fields.
func (w *Withdraw) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case withdraw.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			w.ID = int(value.Int64)
		case withdraw.FieldUsdc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field usdc", values[i])
			} else if value.Valid {
				w.Usdc = value.String
			}
		case withdraw.FieldNonce:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nonce", values[i])
			} else if value.Valid {
				w.Nonce = value.Int64
			}
		case withdraw.FieldFee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value.Valid {
				w.Fee = value.String
			}
		case withdraw.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				w.Time = value.Int64
			}
		case withdraw.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				w.Address = value.String
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Withdraw.
// This includes values selected through modifiers, order, etc.
func (w *Withdraw) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// Update returns a builder for updating this Withdraw.
// Note that you need to call Withdraw.Unwrap() before calling this method if this Withdraw
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Withdraw) Update() *WithdrawUpdateOne {
	return NewWithdrawClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Withdraw entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Withdraw) Unwrap() *Withdraw {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("ent: Withdraw is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Withdraw) String() string {
	var builder strings.Builder
	builder.WriteString("Withdraw(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("usdc=")
	builder.WriteString(w.Usdc)
	builder.WriteString(", ")
	builder.WriteString("nonce=")
	builder.WriteString(fmt.Sprintf("%v", w.Nonce))
	builder.WriteString(", ")
	builder.WriteString("fee=")
	builder.WriteString(w.Fee)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", w.Time))
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(w.Address)
	builder.WriteByte(')')
	return builder.String()
}

// Withdraws is a parsable slice of Withdraw.
type Withdraws []*Withdraw