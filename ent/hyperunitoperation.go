// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/yoshiso/hypersync/ent/hyperunitoperation"
)

// HyperunitOperation is the model entity for the HyperunitOperation schema.
type HyperunitOperation struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// OperationID holds the value of the "operation_id" field.
	OperationID string `json:"operation_id,omitempty"`
	// SourceChain holds the value of the "source_chain" field.
	SourceChain string `json:"source_chain,omitempty"`
	// SourceAmount holds the value of the "source_amount" field.
	SourceAmount string `json:"source_amount,omitempty"`
	// SourceAddress holds the value of the "source_address" field.
	SourceAddress string `json:"source_address,omitempty"`
	// SourceTxHash holds the value of the "source_tx_hash" field.
	SourceTxHash string `json:"source_tx_hash,omitempty"`
	// DestinationTxHash holds the value of the "destination_tx_hash" field.
	DestinationTxHash string `json:"destination_tx_hash,omitempty"`
	// DestinationFeeAmount holds the value of the "destination_fee_amount" field.
	DestinationFeeAmount string `json:"destination_fee_amount,omitempty"`
	// DestinationChain holds the value of the "destination_chain" field.
	DestinationChain string `json:"destination_chain,omitempty"`
	// DestinationAddress holds the value of the "destination_address" field.
	DestinationAddress string `json:"destination_address,omitempty"`
	// SweepFeeAmount holds the value of the "sweep_fee_amount" field.
	SweepFeeAmount string `json:"sweep_fee_amount,omitempty"`
	// OpCreatedAt holds the value of the "op_created_at" field.
	OpCreatedAt time.Time `json:"op_created_at,omitempty"`
	// BroadcastAt holds the value of the "broadcast_at" field.
	BroadcastAt time.Time `json:"broadcast_at,omitempty"`
	// StateUpdatedAt holds the value of the "state_updated_at" field.
	StateUpdatedAt time.Time `json:"state_updated_at,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*HyperunitOperation) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case hyperunitoperation.FieldID:
			values[i] = new(sql.NullInt64)
		case hyperunitoperation.FieldAddress, hyperunitoperation.FieldOperationID, hyperunitoperation.FieldSourceChain, hyperunitoperation.FieldSourceAmount, hyperunitoperation.FieldSourceAddress, hyperunitoperation.FieldSourceTxHash, hyperunitoperation.FieldDestinationTxHash, hyperunitoperation.FieldDestinationFeeAmount, hyperunitoperation.FieldDestinationChain, hyperunitoperation.FieldDestinationAddress, hyperunitoperation.FieldSweepFeeAmount:
			values[i] = new(sql.NullString)
		case hyperunitoperation.FieldOpCreatedAt, hyperunitoperation.FieldBroadcastAt, hyperunitoperation.FieldStateUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the HyperunitOperation fields.
func (ho *HyperunitOperation) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case hyperunitoperation.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ho.ID = int(value.Int64)
		case hyperunitoperation.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				ho.Address = value.String
			}
		case hyperunitoperation.FieldOperationID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field operation_id", values[i])
			} else if value.Valid {
				ho.OperationID = value.String
			}
		case hyperunitoperation.FieldSourceChain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_chain", values[i])
			} else if value.Valid {
				ho.SourceChain = value.String
			}
		case hyperunitoperation.FieldSourceAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_amount", values[i])
			} else if value.Valid {
				ho.SourceAmount = value.String
			}
		case hyperunitoperation.FieldSourceAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_address", values[i])
			} else if value.Valid {
				ho.SourceAddress = value.String
			}
		case hyperunitoperation.FieldSourceTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_tx_hash", values[i])
			} else if value.Valid {
				ho.SourceTxHash = value.String
			}
		case hyperunitoperation.FieldDestinationTxHash:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination_tx_hash", values[i])
			} else if value.Valid {
				ho.DestinationTxHash = value.String
			}
		case hyperunitoperation.FieldDestinationFeeAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination_fee_amount", values[i])
			} else if value.Valid {
				ho.DestinationFeeAmount = value.String
			}
		case hyperunitoperation.FieldDestinationChain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination_chain", values[i])
			} else if value.Valid {
				ho.DestinationChain = value.String
			}
		case hyperunitoperation.FieldDestinationAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field destination_address", values[i])
			} else if value.Valid {
				ho.DestinationAddress = value.String
			}
		case hyperunitoperation.FieldSweepFeeAmount:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field sweep_fee_amount", values[i])
			} else if value.Valid {
				ho.SweepFeeAmount = value.String
			}
		case hyperunitoperation.FieldOpCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field op_created_at", values[i])
			} else if value.Valid {
				ho.OpCreatedAt = value.Time
			}
		case hyperunitoperation.FieldBroadcastAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field broadcast_at", values[i])
			} else if value.Valid {
				ho.BroadcastAt = value.Time
			}
		case hyperunitoperation.FieldStateUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field state_updated_at", values[i])
			} else if value.Valid {
				ho.StateUpdatedAt = value.Time
			}
		default:
			ho.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the HyperunitOperation.
// This includes values selected through modifiers, order, etc.
func (ho *HyperunitOperation) Value(name string) (ent.Value, error) {
	return ho.selectValues.Get(name)
}

// Update returns a builder for updating this HyperunitOperation.
// Note that you need to call HyperunitOperation.Unwrap() before calling this method if this HyperunitOperation
// was returned from a transaction, and the transaction was committed or rolled back.
func (ho *HyperunitOperation) Update() *HyperunitOperationUpdateOne {
	return NewHyperunitOperationClient(ho.config).UpdateOne(ho)
}

// Unwrap unwraps the HyperunitOperation entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ho *HyperunitOperation) Unwrap() *HyperunitOperation {
	_tx, ok := ho.config.driver.(*txDriver)
	if !ok {
		panic("ent: HyperunitOperation is not a transactional entity")
	}
	ho.config.driver = _tx.drv
	return ho
}

// String implements the fmt.Stringer.
func (ho *HyperunitOperation) String() string {
	var builder strings.Builder
	builder.WriteString("HyperunitOperation(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ho.ID))
	builder.WriteString("address=")
	builder.WriteString(ho.Address)
	builder.WriteString(", ")
	builder.WriteString("operation_id=")
	builder.WriteString(ho.OperationID)
	builder.WriteString(", ")
	builder.WriteString("source_chain=")
	builder.WriteString(ho.SourceChain)
	builder.WriteString(", ")
	builder.WriteString("source_amount=")
	builder.WriteString(ho.SourceAmount)
	builder.WriteString(", ")
	builder.WriteString("source_address=")
	builder.WriteString(ho.SourceAddress)
	builder.WriteString(", ")
	builder.WriteString("source_tx_hash=")
	builder.WriteString(ho.SourceTxHash)
	builder.WriteString(", ")
	builder.WriteString("destination_tx_hash=")
	builder.WriteString(ho.DestinationTxHash)
	builder.WriteString(", ")
	builder.WriteString("destination_fee_amount=")
	builder.WriteString(ho.DestinationFeeAmount)
	builder.WriteString(", ")
	builder.WriteString("destination_chain=")
	builder.WriteString(ho.DestinationChain)
	builder.WriteString(", ")
	builder.WriteString("destination_address=")
	builder.WriteString(ho.DestinationAddress)
	builder.WriteString(", ")
	builder.WriteString("sweep_fee_amount=")
	builder.WriteString(ho.SweepFeeAmount)
	builder.WriteString(", ")
	builder.WriteString("op_created_at=")
	builder.WriteString(ho.OpCreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("broadcast_at=")
	builder.WriteString(ho.BroadcastAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("state_updated_at=")
	builder.WriteString(ho.StateUpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// HyperunitOperations is a parsable slice of HyperunitOperation.
type HyperunitOperations []*HyperunitOperation
