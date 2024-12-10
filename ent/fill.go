// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/fill"
)

// Fill is the model entity for the Fill schema.
type Fill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Coin holds the value of the "coin" field.
	Coin string `json:"coin,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Px holds the value of the "px" field.
	Px string `json:"px,omitempty"`
	// Sz holds the value of the "sz" field.
	Sz string `json:"sz,omitempty"`
	// Side holds the value of the "side" field.
	Side string `json:"side,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time,omitempty"`
	// StartPosition holds the value of the "start_position" field.
	StartPosition string `json:"start_position,omitempty"`
	// ClosedPnl holds the value of the "closed_pnl" field.
	ClosedPnl string `json:"closed_pnl,omitempty"`
	// Dir holds the value of the "dir" field.
	Dir string `json:"dir,omitempty"`
	// Hash holds the value of the "hash" field.
	Hash string `json:"hash,omitempty"`
	// Crossed holds the value of the "crossed" field.
	Crossed bool `json:"crossed,omitempty"`
	// Fee holds the value of the "fee" field.
	Fee string `json:"fee,omitempty"`
	// Oid holds the value of the "oid" field.
	Oid int64 `json:"oid,omitempty"`
	// Tid holds the value of the "tid" field.
	Tid int64 `json:"tid,omitempty"`
	// FeeToken holds the value of the "fee_token" field.
	FeeToken string `json:"fee_token,omitempty"`
	// BuilderFee holds the value of the "builder_fee" field.
	BuilderFee   string `json:"builder_fee,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Fill) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case fill.FieldCrossed:
			values[i] = new(sql.NullBool)
		case fill.FieldID, fill.FieldTime, fill.FieldOid, fill.FieldTid:
			values[i] = new(sql.NullInt64)
		case fill.FieldCoin, fill.FieldAddress, fill.FieldPx, fill.FieldSz, fill.FieldSide, fill.FieldStartPosition, fill.FieldClosedPnl, fill.FieldDir, fill.FieldHash, fill.FieldFee, fill.FieldFeeToken, fill.FieldBuilderFee:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Fill fields.
func (f *Fill) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fill.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			f.ID = int(value.Int64)
		case fill.FieldCoin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field coin", values[i])
			} else if value.Valid {
				f.Coin = value.String
			}
		case fill.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				f.Address = value.String
			}
		case fill.FieldPx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field px", values[i])
			} else if value.Valid {
				f.Px = value.String
			}
		case fill.FieldSz:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sz", values[i])
			} else if value.Valid {
				f.Sz = value.String
			}
		case fill.FieldSide:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field side", values[i])
			} else if value.Valid {
				f.Side = value.String
			}
		case fill.FieldTime:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field time", values[i])
			} else if value.Valid {
				f.Time = value.Int64
			}
		case fill.FieldStartPosition:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field start_position", values[i])
			} else if value.Valid {
				f.StartPosition = value.String
			}
		case fill.FieldClosedPnl:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field closed_pnl", values[i])
			} else if value.Valid {
				f.ClosedPnl = value.String
			}
		case fill.FieldDir:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dir", values[i])
			} else if value.Valid {
				f.Dir = value.String
			}
		case fill.FieldHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field hash", values[i])
			} else if value.Valid {
				f.Hash = value.String
			}
		case fill.FieldCrossed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field crossed", values[i])
			} else if value.Valid {
				f.Crossed = value.Bool
			}
		case fill.FieldFee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fee", values[i])
			} else if value.Valid {
				f.Fee = value.String
			}
		case fill.FieldOid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field oid", values[i])
			} else if value.Valid {
				f.Oid = value.Int64
			}
		case fill.FieldTid:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tid", values[i])
			} else if value.Valid {
				f.Tid = value.Int64
			}
		case fill.FieldFeeToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fee_token", values[i])
			} else if value.Valid {
				f.FeeToken = value.String
			}
		case fill.FieldBuilderFee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field builder_fee", values[i])
			} else if value.Valid {
				f.BuilderFee = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Fill.
// This includes values selected through modifiers, order, etc.
func (f *Fill) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this Fill.
// Note that you need to call Fill.Unwrap() before calling this method if this Fill
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Fill) Update() *FillUpdateOne {
	return NewFillClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Fill entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Fill) Unwrap() *Fill {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Fill is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Fill) String() string {
	var builder strings.Builder
	builder.WriteString("Fill(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("coin=")
	builder.WriteString(f.Coin)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(f.Address)
	builder.WriteString(", ")
	builder.WriteString("px=")
	builder.WriteString(f.Px)
	builder.WriteString(", ")
	builder.WriteString("sz=")
	builder.WriteString(f.Sz)
	builder.WriteString(", ")
	builder.WriteString("side=")
	builder.WriteString(f.Side)
	builder.WriteString(", ")
	builder.WriteString("time=")
	builder.WriteString(fmt.Sprintf("%v", f.Time))
	builder.WriteString(", ")
	builder.WriteString("start_position=")
	builder.WriteString(f.StartPosition)
	builder.WriteString(", ")
	builder.WriteString("closed_pnl=")
	builder.WriteString(f.ClosedPnl)
	builder.WriteString(", ")
	builder.WriteString("dir=")
	builder.WriteString(f.Dir)
	builder.WriteString(", ")
	builder.WriteString("hash=")
	builder.WriteString(f.Hash)
	builder.WriteString(", ")
	builder.WriteString("crossed=")
	builder.WriteString(fmt.Sprintf("%v", f.Crossed))
	builder.WriteString(", ")
	builder.WriteString("fee=")
	builder.WriteString(f.Fee)
	builder.WriteString(", ")
	builder.WriteString("oid=")
	builder.WriteString(fmt.Sprintf("%v", f.Oid))
	builder.WriteString(", ")
	builder.WriteString("tid=")
	builder.WriteString(fmt.Sprintf("%v", f.Tid))
	builder.WriteString(", ")
	builder.WriteString("fee_token=")
	builder.WriteString(f.FeeToken)
	builder.WriteString(", ")
	builder.WriteString("builder_fee=")
	builder.WriteString(f.BuilderFee)
	builder.WriteByte(')')
	return builder.String()
}

// Fills is a parsable slice of Fill.
type Fills []*Fill
