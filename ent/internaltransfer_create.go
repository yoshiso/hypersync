// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/internaltransfer"
)

// InternalTransferCreate is the builder for creating a InternalTransfer entity.
type InternalTransferCreate struct {
	config
	mutation *InternalTransferMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUser sets the "user" field.
func (itc *InternalTransferCreate) SetUser(s string) *InternalTransferCreate {
	itc.mutation.SetUser(s)
	return itc
}

// SetDestination sets the "destination" field.
func (itc *InternalTransferCreate) SetDestination(s string) *InternalTransferCreate {
	itc.mutation.SetDestination(s)
	return itc
}

// SetUsdc sets the "usdc" field.
func (itc *InternalTransferCreate) SetUsdc(s string) *InternalTransferCreate {
	itc.mutation.SetUsdc(s)
	return itc
}

// SetFee sets the "fee" field.
func (itc *InternalTransferCreate) SetFee(s string) *InternalTransferCreate {
	itc.mutation.SetFee(s)
	return itc
}

// SetTime sets the "time" field.
func (itc *InternalTransferCreate) SetTime(i int64) *InternalTransferCreate {
	itc.mutation.SetTime(i)
	return itc
}

// SetAddress sets the "address" field.
func (itc *InternalTransferCreate) SetAddress(s string) *InternalTransferCreate {
	itc.mutation.SetAddress(s)
	return itc
}

// Mutation returns the InternalTransferMutation object of the builder.
func (itc *InternalTransferCreate) Mutation() *InternalTransferMutation {
	return itc.mutation
}

// Save creates the InternalTransfer in the database.
func (itc *InternalTransferCreate) Save(ctx context.Context) (*InternalTransfer, error) {
	return withHooks(ctx, itc.sqlSave, itc.mutation, itc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (itc *InternalTransferCreate) SaveX(ctx context.Context) *InternalTransfer {
	v, err := itc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (itc *InternalTransferCreate) Exec(ctx context.Context) error {
	_, err := itc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itc *InternalTransferCreate) ExecX(ctx context.Context) {
	if err := itc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (itc *InternalTransferCreate) check() error {
	if _, ok := itc.mutation.User(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required field "InternalTransfer.user"`)}
	}
	if _, ok := itc.mutation.Destination(); !ok {
		return &ValidationError{Name: "destination", err: errors.New(`ent: missing required field "InternalTransfer.destination"`)}
	}
	if _, ok := itc.mutation.Usdc(); !ok {
		return &ValidationError{Name: "usdc", err: errors.New(`ent: missing required field "InternalTransfer.usdc"`)}
	}
	if _, ok := itc.mutation.Fee(); !ok {
		return &ValidationError{Name: "fee", err: errors.New(`ent: missing required field "InternalTransfer.fee"`)}
	}
	if _, ok := itc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "InternalTransfer.time"`)}
	}
	if _, ok := itc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "InternalTransfer.address"`)}
	}
	return nil
}

func (itc *InternalTransferCreate) sqlSave(ctx context.Context) (*InternalTransfer, error) {
	if err := itc.check(); err != nil {
		return nil, err
	}
	_node, _spec := itc.createSpec()
	if err := sqlgraph.CreateNode(ctx, itc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	itc.mutation.id = &_node.ID
	itc.mutation.done = true
	return _node, nil
}

func (itc *InternalTransferCreate) createSpec() (*InternalTransfer, *sqlgraph.CreateSpec) {
	var (
		_node = &InternalTransfer{config: itc.config}
		_spec = sqlgraph.NewCreateSpec(internaltransfer.Table, sqlgraph.NewFieldSpec(internaltransfer.FieldID, field.TypeInt))
	)
	_spec.OnConflict = itc.conflict
	if value, ok := itc.mutation.User(); ok {
		_spec.SetField(internaltransfer.FieldUser, field.TypeString, value)
		_node.User = value
	}
	if value, ok := itc.mutation.Destination(); ok {
		_spec.SetField(internaltransfer.FieldDestination, field.TypeString, value)
		_node.Destination = value
	}
	if value, ok := itc.mutation.Usdc(); ok {
		_spec.SetField(internaltransfer.FieldUsdc, field.TypeString, value)
		_node.Usdc = value
	}
	if value, ok := itc.mutation.Fee(); ok {
		_spec.SetField(internaltransfer.FieldFee, field.TypeString, value)
		_node.Fee = value
	}
	if value, ok := itc.mutation.Time(); ok {
		_spec.SetField(internaltransfer.FieldTime, field.TypeInt64, value)
		_node.Time = value
	}
	if value, ok := itc.mutation.Address(); ok {
		_spec.SetField(internaltransfer.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.InternalTransfer.Create().
//		SetUser(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InternalTransferUpsert) {
//			SetUser(v+v).
//		}).
//		Exec(ctx)
func (itc *InternalTransferCreate) OnConflict(opts ...sql.ConflictOption) *InternalTransferUpsertOne {
	itc.conflict = opts
	return &InternalTransferUpsertOne{
		create: itc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.InternalTransfer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (itc *InternalTransferCreate) OnConflictColumns(columns ...string) *InternalTransferUpsertOne {
	itc.conflict = append(itc.conflict, sql.ConflictColumns(columns...))
	return &InternalTransferUpsertOne{
		create: itc,
	}
}

type (
	// InternalTransferUpsertOne is the builder for "upsert"-ing
	//  one InternalTransfer node.
	InternalTransferUpsertOne struct {
		create *InternalTransferCreate
	}

	// InternalTransferUpsert is the "OnConflict" setter.
	InternalTransferUpsert struct {
		*sql.UpdateSet
	}
)

// SetUser sets the "user" field.
func (u *InternalTransferUpsert) SetUser(v string) *InternalTransferUpsert {
	u.Set(internaltransfer.FieldUser, v)
	return u
}

// UpdateUser sets the "user" field to the value that was provided on create.
func (u *InternalTransferUpsert) UpdateUser() *InternalTransferUpsert {
	u.SetExcluded(internaltransfer.FieldUser)
	return u
}

// SetDestination sets the "destination" field.
func (u *InternalTransferUpsert) SetDestination(v string) *InternalTransferUpsert {
	u.Set(internaltransfer.FieldDestination, v)
	return u
}

// UpdateDestination sets the "destination" field to the value that was provided on create.
func (u *InternalTransferUpsert) UpdateDestination() *InternalTransferUpsert {
	u.SetExcluded(internaltransfer.FieldDestination)
	return u
}

// SetUsdc sets the "usdc" field.
func (u *InternalTransferUpsert) SetUsdc(v string) *InternalTransferUpsert {
	u.Set(internaltransfer.FieldUsdc, v)
	return u
}

// UpdateUsdc sets the "usdc" field to the value that was provided on create.
func (u *InternalTransferUpsert) UpdateUsdc() *InternalTransferUpsert {
	u.SetExcluded(internaltransfer.FieldUsdc)
	return u
}

// SetFee sets the "fee" field.
func (u *InternalTransferUpsert) SetFee(v string) *InternalTransferUpsert {
	u.Set(internaltransfer.FieldFee, v)
	return u
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *InternalTransferUpsert) UpdateFee() *InternalTransferUpsert {
	u.SetExcluded(internaltransfer.FieldFee)
	return u
}

// SetTime sets the "time" field.
func (u *InternalTransferUpsert) SetTime(v int64) *InternalTransferUpsert {
	u.Set(internaltransfer.FieldTime, v)
	return u
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *InternalTransferUpsert) UpdateTime() *InternalTransferUpsert {
	u.SetExcluded(internaltransfer.FieldTime)
	return u
}

// AddTime adds v to the "time" field.
func (u *InternalTransferUpsert) AddTime(v int64) *InternalTransferUpsert {
	u.Add(internaltransfer.FieldTime, v)
	return u
}

// SetAddress sets the "address" field.
func (u *InternalTransferUpsert) SetAddress(v string) *InternalTransferUpsert {
	u.Set(internaltransfer.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *InternalTransferUpsert) UpdateAddress() *InternalTransferUpsert {
	u.SetExcluded(internaltransfer.FieldAddress)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.InternalTransfer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *InternalTransferUpsertOne) UpdateNewValues() *InternalTransferUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.InternalTransfer.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *InternalTransferUpsertOne) Ignore() *InternalTransferUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InternalTransferUpsertOne) DoNothing() *InternalTransferUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InternalTransferCreate.OnConflict
// documentation for more info.
func (u *InternalTransferUpsertOne) Update(set func(*InternalTransferUpsert)) *InternalTransferUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InternalTransferUpsert{UpdateSet: update})
	}))
	return u
}

// SetUser sets the "user" field.
func (u *InternalTransferUpsertOne) SetUser(v string) *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetUser(v)
	})
}

// UpdateUser sets the "user" field to the value that was provided on create.
func (u *InternalTransferUpsertOne) UpdateUser() *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateUser()
	})
}

// SetDestination sets the "destination" field.
func (u *InternalTransferUpsertOne) SetDestination(v string) *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetDestination(v)
	})
}

// UpdateDestination sets the "destination" field to the value that was provided on create.
func (u *InternalTransferUpsertOne) UpdateDestination() *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateDestination()
	})
}

// SetUsdc sets the "usdc" field.
func (u *InternalTransferUpsertOne) SetUsdc(v string) *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetUsdc(v)
	})
}

// UpdateUsdc sets the "usdc" field to the value that was provided on create.
func (u *InternalTransferUpsertOne) UpdateUsdc() *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateUsdc()
	})
}

// SetFee sets the "fee" field.
func (u *InternalTransferUpsertOne) SetFee(v string) *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetFee(v)
	})
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *InternalTransferUpsertOne) UpdateFee() *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateFee()
	})
}

// SetTime sets the "time" field.
func (u *InternalTransferUpsertOne) SetTime(v int64) *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetTime(v)
	})
}

// AddTime adds v to the "time" field.
func (u *InternalTransferUpsertOne) AddTime(v int64) *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.AddTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *InternalTransferUpsertOne) UpdateTime() *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateTime()
	})
}

// SetAddress sets the "address" field.
func (u *InternalTransferUpsertOne) SetAddress(v string) *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *InternalTransferUpsertOne) UpdateAddress() *InternalTransferUpsertOne {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateAddress()
	})
}

// Exec executes the query.
func (u *InternalTransferUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for InternalTransferCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InternalTransferUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *InternalTransferUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *InternalTransferUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// InternalTransferCreateBulk is the builder for creating many InternalTransfer entities in bulk.
type InternalTransferCreateBulk struct {
	config
	err      error
	builders []*InternalTransferCreate
	conflict []sql.ConflictOption
}

// Save creates the InternalTransfer entities in the database.
func (itcb *InternalTransferCreateBulk) Save(ctx context.Context) ([]*InternalTransfer, error) {
	if itcb.err != nil {
		return nil, itcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(itcb.builders))
	nodes := make([]*InternalTransfer, len(itcb.builders))
	mutators := make([]Mutator, len(itcb.builders))
	for i := range itcb.builders {
		func(i int, root context.Context) {
			builder := itcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InternalTransferMutation)
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
					_, err = mutators[i+1].Mutate(root, itcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = itcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, itcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, itcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (itcb *InternalTransferCreateBulk) SaveX(ctx context.Context) []*InternalTransfer {
	v, err := itcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (itcb *InternalTransferCreateBulk) Exec(ctx context.Context) error {
	_, err := itcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (itcb *InternalTransferCreateBulk) ExecX(ctx context.Context) {
	if err := itcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.InternalTransfer.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.InternalTransferUpsert) {
//			SetUser(v+v).
//		}).
//		Exec(ctx)
func (itcb *InternalTransferCreateBulk) OnConflict(opts ...sql.ConflictOption) *InternalTransferUpsertBulk {
	itcb.conflict = opts
	return &InternalTransferUpsertBulk{
		create: itcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.InternalTransfer.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (itcb *InternalTransferCreateBulk) OnConflictColumns(columns ...string) *InternalTransferUpsertBulk {
	itcb.conflict = append(itcb.conflict, sql.ConflictColumns(columns...))
	return &InternalTransferUpsertBulk{
		create: itcb,
	}
}

// InternalTransferUpsertBulk is the builder for "upsert"-ing
// a bulk of InternalTransfer nodes.
type InternalTransferUpsertBulk struct {
	create *InternalTransferCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.InternalTransfer.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *InternalTransferUpsertBulk) UpdateNewValues() *InternalTransferUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.InternalTransfer.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *InternalTransferUpsertBulk) Ignore() *InternalTransferUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *InternalTransferUpsertBulk) DoNothing() *InternalTransferUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the InternalTransferCreateBulk.OnConflict
// documentation for more info.
func (u *InternalTransferUpsertBulk) Update(set func(*InternalTransferUpsert)) *InternalTransferUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&InternalTransferUpsert{UpdateSet: update})
	}))
	return u
}

// SetUser sets the "user" field.
func (u *InternalTransferUpsertBulk) SetUser(v string) *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetUser(v)
	})
}

// UpdateUser sets the "user" field to the value that was provided on create.
func (u *InternalTransferUpsertBulk) UpdateUser() *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateUser()
	})
}

// SetDestination sets the "destination" field.
func (u *InternalTransferUpsertBulk) SetDestination(v string) *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetDestination(v)
	})
}

// UpdateDestination sets the "destination" field to the value that was provided on create.
func (u *InternalTransferUpsertBulk) UpdateDestination() *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateDestination()
	})
}

// SetUsdc sets the "usdc" field.
func (u *InternalTransferUpsertBulk) SetUsdc(v string) *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetUsdc(v)
	})
}

// UpdateUsdc sets the "usdc" field to the value that was provided on create.
func (u *InternalTransferUpsertBulk) UpdateUsdc() *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateUsdc()
	})
}

// SetFee sets the "fee" field.
func (u *InternalTransferUpsertBulk) SetFee(v string) *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetFee(v)
	})
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *InternalTransferUpsertBulk) UpdateFee() *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateFee()
	})
}

// SetTime sets the "time" field.
func (u *InternalTransferUpsertBulk) SetTime(v int64) *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetTime(v)
	})
}

// AddTime adds v to the "time" field.
func (u *InternalTransferUpsertBulk) AddTime(v int64) *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.AddTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *InternalTransferUpsertBulk) UpdateTime() *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateTime()
	})
}

// SetAddress sets the "address" field.
func (u *InternalTransferUpsertBulk) SetAddress(v string) *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *InternalTransferUpsertBulk) UpdateAddress() *InternalTransferUpsertBulk {
	return u.Update(func(s *InternalTransferUpsert) {
		s.UpdateAddress()
	})
}

// Exec executes the query.
func (u *InternalTransferUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the InternalTransferCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for InternalTransferCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *InternalTransferUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}