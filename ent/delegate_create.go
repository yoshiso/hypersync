// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/delegate"
)

// DelegateCreate is the builder for creating a Delegate entity.
type DelegateCreate struct {
	config
	mutation *DelegateMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetValidator sets the "validator" field.
func (dc *DelegateCreate) SetValidator(s string) *DelegateCreate {
	dc.mutation.SetValidator(s)
	return dc
}

// SetAmount sets the "amount" field.
func (dc *DelegateCreate) SetAmount(s string) *DelegateCreate {
	dc.mutation.SetAmount(s)
	return dc
}

// SetIsUndelegate sets the "is_undelegate" field.
func (dc *DelegateCreate) SetIsUndelegate(b bool) *DelegateCreate {
	dc.mutation.SetIsUndelegate(b)
	return dc
}

// SetTime sets the "time" field.
func (dc *DelegateCreate) SetTime(i int64) *DelegateCreate {
	dc.mutation.SetTime(i)
	return dc
}

// SetAddress sets the "address" field.
func (dc *DelegateCreate) SetAddress(s string) *DelegateCreate {
	dc.mutation.SetAddress(s)
	return dc
}

// Mutation returns the DelegateMutation object of the builder.
func (dc *DelegateCreate) Mutation() *DelegateMutation {
	return dc.mutation
}

// Save creates the Delegate in the database.
func (dc *DelegateCreate) Save(ctx context.Context) (*Delegate, error) {
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DelegateCreate) SaveX(ctx context.Context) *Delegate {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DelegateCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DelegateCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DelegateCreate) check() error {
	if _, ok := dc.mutation.Validator(); !ok {
		return &ValidationError{Name: "validator", err: errors.New(`ent: missing required field "Delegate.validator"`)}
	}
	if _, ok := dc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Delegate.amount"`)}
	}
	if _, ok := dc.mutation.IsUndelegate(); !ok {
		return &ValidationError{Name: "is_undelegate", err: errors.New(`ent: missing required field "Delegate.is_undelegate"`)}
	}
	if _, ok := dc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "Delegate.time"`)}
	}
	if _, ok := dc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Delegate.address"`)}
	}
	return nil
}

func (dc *DelegateCreate) sqlSave(ctx context.Context) (*Delegate, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DelegateCreate) createSpec() (*Delegate, *sqlgraph.CreateSpec) {
	var (
		_node = &Delegate{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(delegate.Table, sqlgraph.NewFieldSpec(delegate.FieldID, field.TypeInt))
	)
	_spec.OnConflict = dc.conflict
	if value, ok := dc.mutation.Validator(); ok {
		_spec.SetField(delegate.FieldValidator, field.TypeString, value)
		_node.Validator = value
	}
	if value, ok := dc.mutation.Amount(); ok {
		_spec.SetField(delegate.FieldAmount, field.TypeString, value)
		_node.Amount = value
	}
	if value, ok := dc.mutation.IsUndelegate(); ok {
		_spec.SetField(delegate.FieldIsUndelegate, field.TypeBool, value)
		_node.IsUndelegate = value
	}
	if value, ok := dc.mutation.Time(); ok {
		_spec.SetField(delegate.FieldTime, field.TypeInt64, value)
		_node.Time = value
	}
	if value, ok := dc.mutation.Address(); ok {
		_spec.SetField(delegate.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Delegate.Create().
//		SetValidator(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DelegateUpsert) {
//			SetValidator(v+v).
//		}).
//		Exec(ctx)
func (dc *DelegateCreate) OnConflict(opts ...sql.ConflictOption) *DelegateUpsertOne {
	dc.conflict = opts
	return &DelegateUpsertOne{
		create: dc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Delegate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dc *DelegateCreate) OnConflictColumns(columns ...string) *DelegateUpsertOne {
	dc.conflict = append(dc.conflict, sql.ConflictColumns(columns...))
	return &DelegateUpsertOne{
		create: dc,
	}
}

type (
	// DelegateUpsertOne is the builder for "upsert"-ing
	//  one Delegate node.
	DelegateUpsertOne struct {
		create *DelegateCreate
	}

	// DelegateUpsert is the "OnConflict" setter.
	DelegateUpsert struct {
		*sql.UpdateSet
	}
)

// SetValidator sets the "validator" field.
func (u *DelegateUpsert) SetValidator(v string) *DelegateUpsert {
	u.Set(delegate.FieldValidator, v)
	return u
}

// UpdateValidator sets the "validator" field to the value that was provided on create.
func (u *DelegateUpsert) UpdateValidator() *DelegateUpsert {
	u.SetExcluded(delegate.FieldValidator)
	return u
}

// SetAmount sets the "amount" field.
func (u *DelegateUpsert) SetAmount(v string) *DelegateUpsert {
	u.Set(delegate.FieldAmount, v)
	return u
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *DelegateUpsert) UpdateAmount() *DelegateUpsert {
	u.SetExcluded(delegate.FieldAmount)
	return u
}

// SetIsUndelegate sets the "is_undelegate" field.
func (u *DelegateUpsert) SetIsUndelegate(v bool) *DelegateUpsert {
	u.Set(delegate.FieldIsUndelegate, v)
	return u
}

// UpdateIsUndelegate sets the "is_undelegate" field to the value that was provided on create.
func (u *DelegateUpsert) UpdateIsUndelegate() *DelegateUpsert {
	u.SetExcluded(delegate.FieldIsUndelegate)
	return u
}

// SetTime sets the "time" field.
func (u *DelegateUpsert) SetTime(v int64) *DelegateUpsert {
	u.Set(delegate.FieldTime, v)
	return u
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *DelegateUpsert) UpdateTime() *DelegateUpsert {
	u.SetExcluded(delegate.FieldTime)
	return u
}

// AddTime adds v to the "time" field.
func (u *DelegateUpsert) AddTime(v int64) *DelegateUpsert {
	u.Add(delegate.FieldTime, v)
	return u
}

// SetAddress sets the "address" field.
func (u *DelegateUpsert) SetAddress(v string) *DelegateUpsert {
	u.Set(delegate.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *DelegateUpsert) UpdateAddress() *DelegateUpsert {
	u.SetExcluded(delegate.FieldAddress)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Delegate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DelegateUpsertOne) UpdateNewValues() *DelegateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Delegate.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *DelegateUpsertOne) Ignore() *DelegateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DelegateUpsertOne) DoNothing() *DelegateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DelegateCreate.OnConflict
// documentation for more info.
func (u *DelegateUpsertOne) Update(set func(*DelegateUpsert)) *DelegateUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DelegateUpsert{UpdateSet: update})
	}))
	return u
}

// SetValidator sets the "validator" field.
func (u *DelegateUpsertOne) SetValidator(v string) *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.SetValidator(v)
	})
}

// UpdateValidator sets the "validator" field to the value that was provided on create.
func (u *DelegateUpsertOne) UpdateValidator() *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateValidator()
	})
}

// SetAmount sets the "amount" field.
func (u *DelegateUpsertOne) SetAmount(v string) *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *DelegateUpsertOne) UpdateAmount() *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateAmount()
	})
}

// SetIsUndelegate sets the "is_undelegate" field.
func (u *DelegateUpsertOne) SetIsUndelegate(v bool) *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.SetIsUndelegate(v)
	})
}

// UpdateIsUndelegate sets the "is_undelegate" field to the value that was provided on create.
func (u *DelegateUpsertOne) UpdateIsUndelegate() *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateIsUndelegate()
	})
}

// SetTime sets the "time" field.
func (u *DelegateUpsertOne) SetTime(v int64) *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.SetTime(v)
	})
}

// AddTime adds v to the "time" field.
func (u *DelegateUpsertOne) AddTime(v int64) *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.AddTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *DelegateUpsertOne) UpdateTime() *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateTime()
	})
}

// SetAddress sets the "address" field.
func (u *DelegateUpsertOne) SetAddress(v string) *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *DelegateUpsertOne) UpdateAddress() *DelegateUpsertOne {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateAddress()
	})
}

// Exec executes the query.
func (u *DelegateUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DelegateCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DelegateUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *DelegateUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *DelegateUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// DelegateCreateBulk is the builder for creating many Delegate entities in bulk.
type DelegateCreateBulk struct {
	config
	err      error
	builders []*DelegateCreate
	conflict []sql.ConflictOption
}

// Save creates the Delegate entities in the database.
func (dcb *DelegateCreateBulk) Save(ctx context.Context) ([]*Delegate, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Delegate, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DelegateMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = dcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DelegateCreateBulk) SaveX(ctx context.Context) []*Delegate {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DelegateCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DelegateCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Delegate.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.DelegateUpsert) {
//			SetValidator(v+v).
//		}).
//		Exec(ctx)
func (dcb *DelegateCreateBulk) OnConflict(opts ...sql.ConflictOption) *DelegateUpsertBulk {
	dcb.conflict = opts
	return &DelegateUpsertBulk{
		create: dcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Delegate.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (dcb *DelegateCreateBulk) OnConflictColumns(columns ...string) *DelegateUpsertBulk {
	dcb.conflict = append(dcb.conflict, sql.ConflictColumns(columns...))
	return &DelegateUpsertBulk{
		create: dcb,
	}
}

// DelegateUpsertBulk is the builder for "upsert"-ing
// a bulk of Delegate nodes.
type DelegateUpsertBulk struct {
	create *DelegateCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Delegate.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *DelegateUpsertBulk) UpdateNewValues() *DelegateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Delegate.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *DelegateUpsertBulk) Ignore() *DelegateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *DelegateUpsertBulk) DoNothing() *DelegateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the DelegateCreateBulk.OnConflict
// documentation for more info.
func (u *DelegateUpsertBulk) Update(set func(*DelegateUpsert)) *DelegateUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&DelegateUpsert{UpdateSet: update})
	}))
	return u
}

// SetValidator sets the "validator" field.
func (u *DelegateUpsertBulk) SetValidator(v string) *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.SetValidator(v)
	})
}

// UpdateValidator sets the "validator" field to the value that was provided on create.
func (u *DelegateUpsertBulk) UpdateValidator() *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateValidator()
	})
}

// SetAmount sets the "amount" field.
func (u *DelegateUpsertBulk) SetAmount(v string) *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.SetAmount(v)
	})
}

// UpdateAmount sets the "amount" field to the value that was provided on create.
func (u *DelegateUpsertBulk) UpdateAmount() *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateAmount()
	})
}

// SetIsUndelegate sets the "is_undelegate" field.
func (u *DelegateUpsertBulk) SetIsUndelegate(v bool) *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.SetIsUndelegate(v)
	})
}

// UpdateIsUndelegate sets the "is_undelegate" field to the value that was provided on create.
func (u *DelegateUpsertBulk) UpdateIsUndelegate() *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateIsUndelegate()
	})
}

// SetTime sets the "time" field.
func (u *DelegateUpsertBulk) SetTime(v int64) *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.SetTime(v)
	})
}

// AddTime adds v to the "time" field.
func (u *DelegateUpsertBulk) AddTime(v int64) *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.AddTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *DelegateUpsertBulk) UpdateTime() *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateTime()
	})
}

// SetAddress sets the "address" field.
func (u *DelegateUpsertBulk) SetAddress(v string) *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *DelegateUpsertBulk) UpdateAddress() *DelegateUpsertBulk {
	return u.Update(func(s *DelegateUpsert) {
		s.UpdateAddress()
	})
}

// Exec executes the query.
func (u *DelegateUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the DelegateCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for DelegateCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *DelegateUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
