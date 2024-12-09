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
	"github.com/yoshiso/hypersync/ent/vaultleadercommission"
)

// VaultLeaderCommissionUpdate is the builder for updating VaultLeaderCommission entities.
type VaultLeaderCommissionUpdate struct {
	config
	hooks    []Hook
	mutation *VaultLeaderCommissionMutation
}

// Where appends a list predicates to the VaultLeaderCommissionUpdate builder.
func (vlcu *VaultLeaderCommissionUpdate) Where(ps ...predicate.VaultLeaderCommission) *VaultLeaderCommissionUpdate {
	vlcu.mutation.Where(ps...)
	return vlcu
}

// SetUser sets the "user" field.
func (vlcu *VaultLeaderCommissionUpdate) SetUser(s string) *VaultLeaderCommissionUpdate {
	vlcu.mutation.SetUser(s)
	return vlcu
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (vlcu *VaultLeaderCommissionUpdate) SetNillableUser(s *string) *VaultLeaderCommissionUpdate {
	if s != nil {
		vlcu.SetUser(*s)
	}
	return vlcu
}

// SetUsdc sets the "usdc" field.
func (vlcu *VaultLeaderCommissionUpdate) SetUsdc(s string) *VaultLeaderCommissionUpdate {
	vlcu.mutation.SetUsdc(s)
	return vlcu
}

// SetNillableUsdc sets the "usdc" field if the given value is not nil.
func (vlcu *VaultLeaderCommissionUpdate) SetNillableUsdc(s *string) *VaultLeaderCommissionUpdate {
	if s != nil {
		vlcu.SetUsdc(*s)
	}
	return vlcu
}

// SetTime sets the "time" field.
func (vlcu *VaultLeaderCommissionUpdate) SetTime(i int64) *VaultLeaderCommissionUpdate {
	vlcu.mutation.ResetTime()
	vlcu.mutation.SetTime(i)
	return vlcu
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (vlcu *VaultLeaderCommissionUpdate) SetNillableTime(i *int64) *VaultLeaderCommissionUpdate {
	if i != nil {
		vlcu.SetTime(*i)
	}
	return vlcu
}

// AddTime adds i to the "time" field.
func (vlcu *VaultLeaderCommissionUpdate) AddTime(i int64) *VaultLeaderCommissionUpdate {
	vlcu.mutation.AddTime(i)
	return vlcu
}

// SetAddress sets the "address" field.
func (vlcu *VaultLeaderCommissionUpdate) SetAddress(s string) *VaultLeaderCommissionUpdate {
	vlcu.mutation.SetAddress(s)
	return vlcu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (vlcu *VaultLeaderCommissionUpdate) SetNillableAddress(s *string) *VaultLeaderCommissionUpdate {
	if s != nil {
		vlcu.SetAddress(*s)
	}
	return vlcu
}

// Mutation returns the VaultLeaderCommissionMutation object of the builder.
func (vlcu *VaultLeaderCommissionUpdate) Mutation() *VaultLeaderCommissionMutation {
	return vlcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vlcu *VaultLeaderCommissionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, vlcu.sqlSave, vlcu.mutation, vlcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vlcu *VaultLeaderCommissionUpdate) SaveX(ctx context.Context) int {
	affected, err := vlcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vlcu *VaultLeaderCommissionUpdate) Exec(ctx context.Context) error {
	_, err := vlcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vlcu *VaultLeaderCommissionUpdate) ExecX(ctx context.Context) {
	if err := vlcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vlcu *VaultLeaderCommissionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(vaultleadercommission.Table, vaultleadercommission.Columns, sqlgraph.NewFieldSpec(vaultleadercommission.FieldID, field.TypeInt))
	if ps := vlcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vlcu.mutation.User(); ok {
		_spec.SetField(vaultleadercommission.FieldUser, field.TypeString, value)
	}
	if value, ok := vlcu.mutation.Usdc(); ok {
		_spec.SetField(vaultleadercommission.FieldUsdc, field.TypeString, value)
	}
	if value, ok := vlcu.mutation.Time(); ok {
		_spec.SetField(vaultleadercommission.FieldTime, field.TypeInt64, value)
	}
	if value, ok := vlcu.mutation.AddedTime(); ok {
		_spec.AddField(vaultleadercommission.FieldTime, field.TypeInt64, value)
	}
	if value, ok := vlcu.mutation.Address(); ok {
		_spec.SetField(vaultleadercommission.FieldAddress, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vlcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vaultleadercommission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vlcu.mutation.done = true
	return n, nil
}

// VaultLeaderCommissionUpdateOne is the builder for updating a single VaultLeaderCommission entity.
type VaultLeaderCommissionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VaultLeaderCommissionMutation
}

// SetUser sets the "user" field.
func (vlcuo *VaultLeaderCommissionUpdateOne) SetUser(s string) *VaultLeaderCommissionUpdateOne {
	vlcuo.mutation.SetUser(s)
	return vlcuo
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (vlcuo *VaultLeaderCommissionUpdateOne) SetNillableUser(s *string) *VaultLeaderCommissionUpdateOne {
	if s != nil {
		vlcuo.SetUser(*s)
	}
	return vlcuo
}

// SetUsdc sets the "usdc" field.
func (vlcuo *VaultLeaderCommissionUpdateOne) SetUsdc(s string) *VaultLeaderCommissionUpdateOne {
	vlcuo.mutation.SetUsdc(s)
	return vlcuo
}

// SetNillableUsdc sets the "usdc" field if the given value is not nil.
func (vlcuo *VaultLeaderCommissionUpdateOne) SetNillableUsdc(s *string) *VaultLeaderCommissionUpdateOne {
	if s != nil {
		vlcuo.SetUsdc(*s)
	}
	return vlcuo
}

// SetTime sets the "time" field.
func (vlcuo *VaultLeaderCommissionUpdateOne) SetTime(i int64) *VaultLeaderCommissionUpdateOne {
	vlcuo.mutation.ResetTime()
	vlcuo.mutation.SetTime(i)
	return vlcuo
}

// SetNillableTime sets the "time" field if the given value is not nil.
func (vlcuo *VaultLeaderCommissionUpdateOne) SetNillableTime(i *int64) *VaultLeaderCommissionUpdateOne {
	if i != nil {
		vlcuo.SetTime(*i)
	}
	return vlcuo
}

// AddTime adds i to the "time" field.
func (vlcuo *VaultLeaderCommissionUpdateOne) AddTime(i int64) *VaultLeaderCommissionUpdateOne {
	vlcuo.mutation.AddTime(i)
	return vlcuo
}

// SetAddress sets the "address" field.
func (vlcuo *VaultLeaderCommissionUpdateOne) SetAddress(s string) *VaultLeaderCommissionUpdateOne {
	vlcuo.mutation.SetAddress(s)
	return vlcuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (vlcuo *VaultLeaderCommissionUpdateOne) SetNillableAddress(s *string) *VaultLeaderCommissionUpdateOne {
	if s != nil {
		vlcuo.SetAddress(*s)
	}
	return vlcuo
}

// Mutation returns the VaultLeaderCommissionMutation object of the builder.
func (vlcuo *VaultLeaderCommissionUpdateOne) Mutation() *VaultLeaderCommissionMutation {
	return vlcuo.mutation
}

// Where appends a list predicates to the VaultLeaderCommissionUpdate builder.
func (vlcuo *VaultLeaderCommissionUpdateOne) Where(ps ...predicate.VaultLeaderCommission) *VaultLeaderCommissionUpdateOne {
	vlcuo.mutation.Where(ps...)
	return vlcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vlcuo *VaultLeaderCommissionUpdateOne) Select(field string, fields ...string) *VaultLeaderCommissionUpdateOne {
	vlcuo.fields = append([]string{field}, fields...)
	return vlcuo
}

// Save executes the query and returns the updated VaultLeaderCommission entity.
func (vlcuo *VaultLeaderCommissionUpdateOne) Save(ctx context.Context) (*VaultLeaderCommission, error) {
	return withHooks(ctx, vlcuo.sqlSave, vlcuo.mutation, vlcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vlcuo *VaultLeaderCommissionUpdateOne) SaveX(ctx context.Context) *VaultLeaderCommission {
	node, err := vlcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vlcuo *VaultLeaderCommissionUpdateOne) Exec(ctx context.Context) error {
	_, err := vlcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vlcuo *VaultLeaderCommissionUpdateOne) ExecX(ctx context.Context) {
	if err := vlcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vlcuo *VaultLeaderCommissionUpdateOne) sqlSave(ctx context.Context) (_node *VaultLeaderCommission, err error) {
	_spec := sqlgraph.NewUpdateSpec(vaultleadercommission.Table, vaultleadercommission.Columns, sqlgraph.NewFieldSpec(vaultleadercommission.FieldID, field.TypeInt))
	id, ok := vlcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "VaultLeaderCommission.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vlcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vaultleadercommission.FieldID)
		for _, f := range fields {
			if !vaultleadercommission.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vaultleadercommission.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vlcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vlcuo.mutation.User(); ok {
		_spec.SetField(vaultleadercommission.FieldUser, field.TypeString, value)
	}
	if value, ok := vlcuo.mutation.Usdc(); ok {
		_spec.SetField(vaultleadercommission.FieldUsdc, field.TypeString, value)
	}
	if value, ok := vlcuo.mutation.Time(); ok {
		_spec.SetField(vaultleadercommission.FieldTime, field.TypeInt64, value)
	}
	if value, ok := vlcuo.mutation.AddedTime(); ok {
		_spec.AddField(vaultleadercommission.FieldTime, field.TypeInt64, value)
	}
	if value, ok := vlcuo.mutation.Address(); ok {
		_spec.SetField(vaultleadercommission.FieldAddress, field.TypeString, value)
	}
	_node = &VaultLeaderCommission{config: vlcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vlcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vaultleadercommission.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vlcuo.mutation.done = true
	return _node, nil
}
