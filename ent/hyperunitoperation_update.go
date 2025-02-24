// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/hyperunitoperation"
	"github.com/yoshiso/hypersync/ent/predicate"
)

// HyperunitOperationUpdate is the builder for updating HyperunitOperation entities.
type HyperunitOperationUpdate struct {
	config
	hooks    []Hook
	mutation *HyperunitOperationMutation
}

// Where appends a list predicates to the HyperunitOperationUpdate builder.
func (hou *HyperunitOperationUpdate) Where(ps ...predicate.HyperunitOperation) *HyperunitOperationUpdate {
	hou.mutation.Where(ps...)
	return hou
}

// SetAddress sets the "address" field.
func (hou *HyperunitOperationUpdate) SetAddress(s string) *HyperunitOperationUpdate {
	hou.mutation.SetAddress(s)
	return hou
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableAddress(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetAddress(*s)
	}
	return hou
}

// SetOperationID sets the "operation_id" field.
func (hou *HyperunitOperationUpdate) SetOperationID(s string) *HyperunitOperationUpdate {
	hou.mutation.SetOperationID(s)
	return hou
}

// SetNillableOperationID sets the "operation_id" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableOperationID(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetOperationID(*s)
	}
	return hou
}

// SetSourceChain sets the "source_chain" field.
func (hou *HyperunitOperationUpdate) SetSourceChain(s string) *HyperunitOperationUpdate {
	hou.mutation.SetSourceChain(s)
	return hou
}

// SetNillableSourceChain sets the "source_chain" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableSourceChain(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetSourceChain(*s)
	}
	return hou
}

// SetSourceAmount sets the "source_amount" field.
func (hou *HyperunitOperationUpdate) SetSourceAmount(s string) *HyperunitOperationUpdate {
	hou.mutation.SetSourceAmount(s)
	return hou
}

// SetNillableSourceAmount sets the "source_amount" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableSourceAmount(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetSourceAmount(*s)
	}
	return hou
}

// SetSourceAddress sets the "source_address" field.
func (hou *HyperunitOperationUpdate) SetSourceAddress(s string) *HyperunitOperationUpdate {
	hou.mutation.SetSourceAddress(s)
	return hou
}

// SetNillableSourceAddress sets the "source_address" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableSourceAddress(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetSourceAddress(*s)
	}
	return hou
}

// SetSourceTxHash sets the "source_tx_hash" field.
func (hou *HyperunitOperationUpdate) SetSourceTxHash(s string) *HyperunitOperationUpdate {
	hou.mutation.SetSourceTxHash(s)
	return hou
}

// SetNillableSourceTxHash sets the "source_tx_hash" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableSourceTxHash(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetSourceTxHash(*s)
	}
	return hou
}

// SetDestinationTxHash sets the "destination_tx_hash" field.
func (hou *HyperunitOperationUpdate) SetDestinationTxHash(s string) *HyperunitOperationUpdate {
	hou.mutation.SetDestinationTxHash(s)
	return hou
}

// SetNillableDestinationTxHash sets the "destination_tx_hash" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableDestinationTxHash(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetDestinationTxHash(*s)
	}
	return hou
}

// SetDestinationFeeAmount sets the "destination_fee_amount" field.
func (hou *HyperunitOperationUpdate) SetDestinationFeeAmount(s string) *HyperunitOperationUpdate {
	hou.mutation.SetDestinationFeeAmount(s)
	return hou
}

// SetNillableDestinationFeeAmount sets the "destination_fee_amount" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableDestinationFeeAmount(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetDestinationFeeAmount(*s)
	}
	return hou
}

// SetDestinationChain sets the "destination_chain" field.
func (hou *HyperunitOperationUpdate) SetDestinationChain(s string) *HyperunitOperationUpdate {
	hou.mutation.SetDestinationChain(s)
	return hou
}

// SetNillableDestinationChain sets the "destination_chain" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableDestinationChain(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetDestinationChain(*s)
	}
	return hou
}

// SetDestinationAddress sets the "destination_address" field.
func (hou *HyperunitOperationUpdate) SetDestinationAddress(s string) *HyperunitOperationUpdate {
	hou.mutation.SetDestinationAddress(s)
	return hou
}

// SetNillableDestinationAddress sets the "destination_address" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableDestinationAddress(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetDestinationAddress(*s)
	}
	return hou
}

// SetSweepFeeAmount sets the "sweep_fee_amount" field.
func (hou *HyperunitOperationUpdate) SetSweepFeeAmount(s string) *HyperunitOperationUpdate {
	hou.mutation.SetSweepFeeAmount(s)
	return hou
}

// SetNillableSweepFeeAmount sets the "sweep_fee_amount" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableSweepFeeAmount(s *string) *HyperunitOperationUpdate {
	if s != nil {
		hou.SetSweepFeeAmount(*s)
	}
	return hou
}

// SetOpCreatedAt sets the "op_created_at" field.
func (hou *HyperunitOperationUpdate) SetOpCreatedAt(t time.Time) *HyperunitOperationUpdate {
	hou.mutation.SetOpCreatedAt(t)
	return hou
}

// SetNillableOpCreatedAt sets the "op_created_at" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableOpCreatedAt(t *time.Time) *HyperunitOperationUpdate {
	if t != nil {
		hou.SetOpCreatedAt(*t)
	}
	return hou
}

// SetBroadcastAt sets the "broadcast_at" field.
func (hou *HyperunitOperationUpdate) SetBroadcastAt(t time.Time) *HyperunitOperationUpdate {
	hou.mutation.SetBroadcastAt(t)
	return hou
}

// SetNillableBroadcastAt sets the "broadcast_at" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableBroadcastAt(t *time.Time) *HyperunitOperationUpdate {
	if t != nil {
		hou.SetBroadcastAt(*t)
	}
	return hou
}

// SetStateUpdatedAt sets the "state_updated_at" field.
func (hou *HyperunitOperationUpdate) SetStateUpdatedAt(t time.Time) *HyperunitOperationUpdate {
	hou.mutation.SetStateUpdatedAt(t)
	return hou
}

// SetNillableStateUpdatedAt sets the "state_updated_at" field if the given value is not nil.
func (hou *HyperunitOperationUpdate) SetNillableStateUpdatedAt(t *time.Time) *HyperunitOperationUpdate {
	if t != nil {
		hou.SetStateUpdatedAt(*t)
	}
	return hou
}

// Mutation returns the HyperunitOperationMutation object of the builder.
func (hou *HyperunitOperationUpdate) Mutation() *HyperunitOperationMutation {
	return hou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (hou *HyperunitOperationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, hou.sqlSave, hou.mutation, hou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (hou *HyperunitOperationUpdate) SaveX(ctx context.Context) int {
	affected, err := hou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (hou *HyperunitOperationUpdate) Exec(ctx context.Context) error {
	_, err := hou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (hou *HyperunitOperationUpdate) ExecX(ctx context.Context) {
	if err := hou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (hou *HyperunitOperationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(hyperunitoperation.Table, hyperunitoperation.Columns, sqlgraph.NewFieldSpec(hyperunitoperation.FieldID, field.TypeInt))
	if ps := hou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := hou.mutation.Address(); ok {
		_spec.SetField(hyperunitoperation.FieldAddress, field.TypeString, value)
	}
	if value, ok := hou.mutation.OperationID(); ok {
		_spec.SetField(hyperunitoperation.FieldOperationID, field.TypeString, value)
	}
	if value, ok := hou.mutation.SourceChain(); ok {
		_spec.SetField(hyperunitoperation.FieldSourceChain, field.TypeString, value)
	}
	if value, ok := hou.mutation.SourceAmount(); ok {
		_spec.SetField(hyperunitoperation.FieldSourceAmount, field.TypeString, value)
	}
	if value, ok := hou.mutation.SourceAddress(); ok {
		_spec.SetField(hyperunitoperation.FieldSourceAddress, field.TypeString, value)
	}
	if value, ok := hou.mutation.SourceTxHash(); ok {
		_spec.SetField(hyperunitoperation.FieldSourceTxHash, field.TypeString, value)
	}
	if value, ok := hou.mutation.DestinationTxHash(); ok {
		_spec.SetField(hyperunitoperation.FieldDestinationTxHash, field.TypeString, value)
	}
	if value, ok := hou.mutation.DestinationFeeAmount(); ok {
		_spec.SetField(hyperunitoperation.FieldDestinationFeeAmount, field.TypeString, value)
	}
	if value, ok := hou.mutation.DestinationChain(); ok {
		_spec.SetField(hyperunitoperation.FieldDestinationChain, field.TypeString, value)
	}
	if value, ok := hou.mutation.DestinationAddress(); ok {
		_spec.SetField(hyperunitoperation.FieldDestinationAddress, field.TypeString, value)
	}
	if value, ok := hou.mutation.SweepFeeAmount(); ok {
		_spec.SetField(hyperunitoperation.FieldSweepFeeAmount, field.TypeString, value)
	}
	if value, ok := hou.mutation.OpCreatedAt(); ok {
		_spec.SetField(hyperunitoperation.FieldOpCreatedAt, field.TypeTime, value)
	}
	if value, ok := hou.mutation.BroadcastAt(); ok {
		_spec.SetField(hyperunitoperation.FieldBroadcastAt, field.TypeTime, value)
	}
	if value, ok := hou.mutation.StateUpdatedAt(); ok {
		_spec.SetField(hyperunitoperation.FieldStateUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, hou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hyperunitoperation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	hou.mutation.done = true
	return n, nil
}

// HyperunitOperationUpdateOne is the builder for updating a single HyperunitOperation entity.
type HyperunitOperationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *HyperunitOperationMutation
}

// SetAddress sets the "address" field.
func (houo *HyperunitOperationUpdateOne) SetAddress(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetAddress(s)
	return houo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableAddress(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetAddress(*s)
	}
	return houo
}

// SetOperationID sets the "operation_id" field.
func (houo *HyperunitOperationUpdateOne) SetOperationID(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetOperationID(s)
	return houo
}

// SetNillableOperationID sets the "operation_id" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableOperationID(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetOperationID(*s)
	}
	return houo
}

// SetSourceChain sets the "source_chain" field.
func (houo *HyperunitOperationUpdateOne) SetSourceChain(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetSourceChain(s)
	return houo
}

// SetNillableSourceChain sets the "source_chain" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableSourceChain(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetSourceChain(*s)
	}
	return houo
}

// SetSourceAmount sets the "source_amount" field.
func (houo *HyperunitOperationUpdateOne) SetSourceAmount(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetSourceAmount(s)
	return houo
}

// SetNillableSourceAmount sets the "source_amount" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableSourceAmount(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetSourceAmount(*s)
	}
	return houo
}

// SetSourceAddress sets the "source_address" field.
func (houo *HyperunitOperationUpdateOne) SetSourceAddress(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetSourceAddress(s)
	return houo
}

// SetNillableSourceAddress sets the "source_address" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableSourceAddress(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetSourceAddress(*s)
	}
	return houo
}

// SetSourceTxHash sets the "source_tx_hash" field.
func (houo *HyperunitOperationUpdateOne) SetSourceTxHash(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetSourceTxHash(s)
	return houo
}

// SetNillableSourceTxHash sets the "source_tx_hash" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableSourceTxHash(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetSourceTxHash(*s)
	}
	return houo
}

// SetDestinationTxHash sets the "destination_tx_hash" field.
func (houo *HyperunitOperationUpdateOne) SetDestinationTxHash(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetDestinationTxHash(s)
	return houo
}

// SetNillableDestinationTxHash sets the "destination_tx_hash" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableDestinationTxHash(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetDestinationTxHash(*s)
	}
	return houo
}

// SetDestinationFeeAmount sets the "destination_fee_amount" field.
func (houo *HyperunitOperationUpdateOne) SetDestinationFeeAmount(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetDestinationFeeAmount(s)
	return houo
}

// SetNillableDestinationFeeAmount sets the "destination_fee_amount" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableDestinationFeeAmount(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetDestinationFeeAmount(*s)
	}
	return houo
}

// SetDestinationChain sets the "destination_chain" field.
func (houo *HyperunitOperationUpdateOne) SetDestinationChain(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetDestinationChain(s)
	return houo
}

// SetNillableDestinationChain sets the "destination_chain" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableDestinationChain(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetDestinationChain(*s)
	}
	return houo
}

// SetDestinationAddress sets the "destination_address" field.
func (houo *HyperunitOperationUpdateOne) SetDestinationAddress(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetDestinationAddress(s)
	return houo
}

// SetNillableDestinationAddress sets the "destination_address" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableDestinationAddress(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetDestinationAddress(*s)
	}
	return houo
}

// SetSweepFeeAmount sets the "sweep_fee_amount" field.
func (houo *HyperunitOperationUpdateOne) SetSweepFeeAmount(s string) *HyperunitOperationUpdateOne {
	houo.mutation.SetSweepFeeAmount(s)
	return houo
}

// SetNillableSweepFeeAmount sets the "sweep_fee_amount" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableSweepFeeAmount(s *string) *HyperunitOperationUpdateOne {
	if s != nil {
		houo.SetSweepFeeAmount(*s)
	}
	return houo
}

// SetOpCreatedAt sets the "op_created_at" field.
func (houo *HyperunitOperationUpdateOne) SetOpCreatedAt(t time.Time) *HyperunitOperationUpdateOne {
	houo.mutation.SetOpCreatedAt(t)
	return houo
}

// SetNillableOpCreatedAt sets the "op_created_at" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableOpCreatedAt(t *time.Time) *HyperunitOperationUpdateOne {
	if t != nil {
		houo.SetOpCreatedAt(*t)
	}
	return houo
}

// SetBroadcastAt sets the "broadcast_at" field.
func (houo *HyperunitOperationUpdateOne) SetBroadcastAt(t time.Time) *HyperunitOperationUpdateOne {
	houo.mutation.SetBroadcastAt(t)
	return houo
}

// SetNillableBroadcastAt sets the "broadcast_at" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableBroadcastAt(t *time.Time) *HyperunitOperationUpdateOne {
	if t != nil {
		houo.SetBroadcastAt(*t)
	}
	return houo
}

// SetStateUpdatedAt sets the "state_updated_at" field.
func (houo *HyperunitOperationUpdateOne) SetStateUpdatedAt(t time.Time) *HyperunitOperationUpdateOne {
	houo.mutation.SetStateUpdatedAt(t)
	return houo
}

// SetNillableStateUpdatedAt sets the "state_updated_at" field if the given value is not nil.
func (houo *HyperunitOperationUpdateOne) SetNillableStateUpdatedAt(t *time.Time) *HyperunitOperationUpdateOne {
	if t != nil {
		houo.SetStateUpdatedAt(*t)
	}
	return houo
}

// Mutation returns the HyperunitOperationMutation object of the builder.
func (houo *HyperunitOperationUpdateOne) Mutation() *HyperunitOperationMutation {
	return houo.mutation
}

// Where appends a list predicates to the HyperunitOperationUpdate builder.
func (houo *HyperunitOperationUpdateOne) Where(ps ...predicate.HyperunitOperation) *HyperunitOperationUpdateOne {
	houo.mutation.Where(ps...)
	return houo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (houo *HyperunitOperationUpdateOne) Select(field string, fields ...string) *HyperunitOperationUpdateOne {
	houo.fields = append([]string{field}, fields...)
	return houo
}

// Save executes the query and returns the updated HyperunitOperation entity.
func (houo *HyperunitOperationUpdateOne) Save(ctx context.Context) (*HyperunitOperation, error) {
	return withHooks(ctx, houo.sqlSave, houo.mutation, houo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (houo *HyperunitOperationUpdateOne) SaveX(ctx context.Context) *HyperunitOperation {
	node, err := houo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (houo *HyperunitOperationUpdateOne) Exec(ctx context.Context) error {
	_, err := houo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (houo *HyperunitOperationUpdateOne) ExecX(ctx context.Context) {
	if err := houo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (houo *HyperunitOperationUpdateOne) sqlSave(ctx context.Context) (_node *HyperunitOperation, err error) {
	_spec := sqlgraph.NewUpdateSpec(hyperunitoperation.Table, hyperunitoperation.Columns, sqlgraph.NewFieldSpec(hyperunitoperation.FieldID, field.TypeInt))
	id, ok := houo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "HyperunitOperation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := houo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, hyperunitoperation.FieldID)
		for _, f := range fields {
			if !hyperunitoperation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != hyperunitoperation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := houo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := houo.mutation.Address(); ok {
		_spec.SetField(hyperunitoperation.FieldAddress, field.TypeString, value)
	}
	if value, ok := houo.mutation.OperationID(); ok {
		_spec.SetField(hyperunitoperation.FieldOperationID, field.TypeString, value)
	}
	if value, ok := houo.mutation.SourceChain(); ok {
		_spec.SetField(hyperunitoperation.FieldSourceChain, field.TypeString, value)
	}
	if value, ok := houo.mutation.SourceAmount(); ok {
		_spec.SetField(hyperunitoperation.FieldSourceAmount, field.TypeString, value)
	}
	if value, ok := houo.mutation.SourceAddress(); ok {
		_spec.SetField(hyperunitoperation.FieldSourceAddress, field.TypeString, value)
	}
	if value, ok := houo.mutation.SourceTxHash(); ok {
		_spec.SetField(hyperunitoperation.FieldSourceTxHash, field.TypeString, value)
	}
	if value, ok := houo.mutation.DestinationTxHash(); ok {
		_spec.SetField(hyperunitoperation.FieldDestinationTxHash, field.TypeString, value)
	}
	if value, ok := houo.mutation.DestinationFeeAmount(); ok {
		_spec.SetField(hyperunitoperation.FieldDestinationFeeAmount, field.TypeString, value)
	}
	if value, ok := houo.mutation.DestinationChain(); ok {
		_spec.SetField(hyperunitoperation.FieldDestinationChain, field.TypeString, value)
	}
	if value, ok := houo.mutation.DestinationAddress(); ok {
		_spec.SetField(hyperunitoperation.FieldDestinationAddress, field.TypeString, value)
	}
	if value, ok := houo.mutation.SweepFeeAmount(); ok {
		_spec.SetField(hyperunitoperation.FieldSweepFeeAmount, field.TypeString, value)
	}
	if value, ok := houo.mutation.OpCreatedAt(); ok {
		_spec.SetField(hyperunitoperation.FieldOpCreatedAt, field.TypeTime, value)
	}
	if value, ok := houo.mutation.BroadcastAt(); ok {
		_spec.SetField(hyperunitoperation.FieldBroadcastAt, field.TypeTime, value)
	}
	if value, ok := houo.mutation.StateUpdatedAt(); ok {
		_spec.SetField(hyperunitoperation.FieldStateUpdatedAt, field.TypeTime, value)
	}
	_node = &HyperunitOperation{config: houo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, houo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{hyperunitoperation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	houo.mutation.done = true
	return _node, nil
}
