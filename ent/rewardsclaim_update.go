// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/predicate"
	"github.com/yoshiso/hypersync/ent/rewardsclaim"
)

// RewardsClaimUpdate is the builder for updating RewardsClaim entities.
type RewardsClaimUpdate struct {
	config
	hooks    []Hook
	mutation *RewardsClaimMutation
}

// Where appends a list predicates to the RewardsClaimUpdate builder.
func (rcu *RewardsClaimUpdate) Where(ps ...predicate.RewardsClaim) *RewardsClaimUpdate {
	rcu.mutation.Where(ps...)
	return rcu
}

// SetAmount sets the "amount" field.
func (rcu *RewardsClaimUpdate) SetAmount(s string) *RewardsClaimUpdate {
	rcu.mutation.SetAmount(s)
	return rcu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (rcu *RewardsClaimUpdate) SetNillableAmount(s *string) *RewardsClaimUpdate {
	if s != nil {
		rcu.SetAmount(*s)
	}
	return rcu
}

// SetTime sets the "time" field.
func (rcu *RewardsClaimUpdate) SetTime(i int64) *RewardsClaimUpdate {
	rcu.mutation.ResetTime()
	rcu.mutation.SetTime(i)
	return rcu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (rcu *RewardsClaimUpdate) SetNillableTime(i *int64) *RewardsClaimUpdate {
	if i != nil {
		rcu.SetTime(*i)
	}
	return rcu
}

// AddTime adds i to the "time" field.
func (rcu *RewardsClaimUpdate) AddTime(i int64) *RewardsClaimUpdate {
	rcu.mutation.AddTime(i)
	return rcu
}

// SetAddress sets the "address" field.
func (rcu *RewardsClaimUpdate) SetAddress(s string) *RewardsClaimUpdate {
	rcu.mutation.SetAddress(s)
	return rcu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (rcu *RewardsClaimUpdate) SetNillableAddress(s *string) *RewardsClaimUpdate {
	if s != nil {
		rcu.SetAddress(*s)
	}
	return rcu
}

// Mutation returns the RewardsClaimMutation object of the builder.
func (rcu *RewardsClaimUpdate) Mutation() *RewardsClaimMutation {
	return rcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rcu *RewardsClaimUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rcu.sqlSave, rcu.mutation, rcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcu *RewardsClaimUpdate) SaveX(ctx context.Context) int {
	affected, err := rcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rcu *RewardsClaimUpdate) Exec(ctx context.Context) error {
	_, err := rcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcu *RewardsClaimUpdate) ExecX(ctx context.Context) {
	if err := rcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rcu *RewardsClaimUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(rewardsclaim.Table, rewardsclaim.Columns, sqlgraph.NewFieldSpec(rewardsclaim.FieldID, field.TypeInt))
	if ps := rcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcu.mutation.Amount(); ok {
		_spec.SetField(rewardsclaim.FieldAmount, field.TypeString, value)
	}
	if value, ok := rcu.mutation.Time(); ok {
		_spec.SetField(rewardsclaim.FieldTime, field.TypeInt64, value)
	}
	if value, ok := rcu.mutation.AddedTime(); ok {
		_spec.AddField(rewardsclaim.FieldTime, field.TypeInt64, value)
	}
	if value, ok := rcu.mutation.Address(); ok {
		_spec.SetField(rewardsclaim.FieldAddress, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rewardsclaim.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rcu.mutation.done = true
	return n, nil
}

// RewardsClaimUpdateOne is the builder for updating a single RewardsClaim entity.
type RewardsClaimUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RewardsClaimMutation
}

// SetAmount sets the "amount" field.
func (rcuo *RewardsClaimUpdateOne) SetAmount(s string) *RewardsClaimUpdateOne {
	rcuo.mutation.SetAmount(s)
	return rcuo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (rcuo *RewardsClaimUpdateOne) SetNillableAmount(s *string) *RewardsClaimUpdateOne {
	if s != nil {
		rcuo.SetAmount(*s)
	}
	return rcuo
}

// SetTime sets the "time" field.
func (rcuo *RewardsClaimUpdateOne) SetTime(i int64) *RewardsClaimUpdateOne {
	rcuo.mutation.ResetTime()
	rcuo.mutation.SetTime(i)
	return rcuo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (rcuo *RewardsClaimUpdateOne) SetNillableTime(i *int64) *RewardsClaimUpdateOne {
	if i != nil {
		rcuo.SetTime(*i)
	}
	return rcuo
}

// AddTime adds i to the "time" field.
func (rcuo *RewardsClaimUpdateOne) AddTime(i int64) *RewardsClaimUpdateOne {
	rcuo.mutation.AddTime(i)
	return rcuo
}

// SetAddress sets the "address" field.
func (rcuo *RewardsClaimUpdateOne) SetAddress(s string) *RewardsClaimUpdateOne {
	rcuo.mutation.SetAddress(s)
	return rcuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (rcuo *RewardsClaimUpdateOne) SetNillableAddress(s *string) *RewardsClaimUpdateOne {
	if s != nil {
		rcuo.SetAddress(*s)
	}
	return rcuo
}

// Mutation returns the RewardsClaimMutation object of the builder.
func (rcuo *RewardsClaimUpdateOne) Mutation() *RewardsClaimMutation {
	return rcuo.mutation
}

// Where appends a list predicates to the RewardsClaimUpdate builder.
func (rcuo *RewardsClaimUpdateOne) Where(ps ...predicate.RewardsClaim) *RewardsClaimUpdateOne {
	rcuo.mutation.Where(ps...)
	return rcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rcuo *RewardsClaimUpdateOne) Select(field string, fields ...string) *RewardsClaimUpdateOne {
	rcuo.fields = append([]string{field}, fields...)
	return rcuo
}

// Save executes the query and returns the updated RewardsClaim entity.
func (rcuo *RewardsClaimUpdateOne) Save(ctx context.Context) (*RewardsClaim, error) {
	return withHooks(ctx, rcuo.sqlSave, rcuo.mutation, rcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rcuo *RewardsClaimUpdateOne) SaveX(ctx context.Context) *RewardsClaim {
	node, err := rcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rcuo *RewardsClaimUpdateOne) Exec(ctx context.Context) error {
	_, err := rcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcuo *RewardsClaimUpdateOne) ExecX(ctx context.Context) {
	if err := rcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rcuo *RewardsClaimUpdateOne) sqlSave(ctx context.Context) (_node *RewardsClaim, err error) {
	_spec := sqlgraph.NewUpdateSpec(rewardsclaim.Table, rewardsclaim.Columns, sqlgraph.NewFieldSpec(rewardsclaim.FieldID, field.TypeInt))
	id, ok := rcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RewardsClaim.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rewardsclaim.FieldID)
		for _, f := range fields {
			if !rewardsclaim.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rewardsclaim.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rcuo.mutation.Amount(); ok {
		_spec.SetField(rewardsclaim.FieldAmount, field.TypeString, value)
	}
	if value, ok := rcuo.mutation.Time(); ok {
		_spec.SetField(rewardsclaim.FieldTime, field.TypeInt64, value)
	}
	if value, ok := rcuo.mutation.AddedTime(); ok {
		_spec.AddField(rewardsclaim.FieldTime, field.TypeInt64, value)
	}
	if value, ok := rcuo.mutation.Address(); ok {
		_spec.SetField(rewardsclaim.FieldAddress, field.TypeString, value)
	}
	_node = &RewardsClaim{config: rcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rewardsclaim.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rcuo.mutation.done = true
	return _node, nil
}
