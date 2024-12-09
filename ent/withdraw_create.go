// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/withdraw"
)

// WithdrawCreate is the builder for creating a Withdraw entity.
type WithdrawCreate struct {
	config
	mutation *WithdrawMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetUsdc sets the "usdc" field.
func (wc *WithdrawCreate) SetUsdc(s string) *WithdrawCreate {
	wc.mutation.SetUsdc(s)
	return wc
}

// SetNonce sets the "nonce" field.
func (wc *WithdrawCreate) SetNonce(i int64) *WithdrawCreate {
	wc.mutation.SetNonce(i)
	return wc
}

// SetFee sets the "fee" field.
func (wc *WithdrawCreate) SetFee(s string) *WithdrawCreate {
	wc.mutation.SetFee(s)
	return wc
}

// SetTime sets the "time" field.
func (wc *WithdrawCreate) SetTime(i int64) *WithdrawCreate {
	wc.mutation.SetTime(i)
	return wc
}

// SetAddress sets the "address" field.
func (wc *WithdrawCreate) SetAddress(s string) *WithdrawCreate {
	wc.mutation.SetAddress(s)
	return wc
}

// Mutation returns the WithdrawMutation object of the builder.
func (wc *WithdrawCreate) Mutation() *WithdrawMutation {
	return wc.mutation
}

// Save creates the Withdraw in the database.
func (wc *WithdrawCreate) Save(ctx context.Context) (*Withdraw, error) {
	return withHooks(ctx, wc.sqlSave, wc.mutation, wc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wc *WithdrawCreate) SaveX(ctx context.Context) *Withdraw {
	v, err := wc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wc *WithdrawCreate) Exec(ctx context.Context) error {
	_, err := wc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wc *WithdrawCreate) ExecX(ctx context.Context) {
	if err := wc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wc *WithdrawCreate) check() error {
	if _, ok := wc.mutation.Usdc(); !ok {
		return &ValidationError{Name: "usdc", err: errors.New(`ent: missing required field "Withdraw.usdc"`)}
	}
	if _, ok := wc.mutation.Nonce(); !ok {
		return &ValidationError{Name: "nonce", err: errors.New(`ent: missing required field "Withdraw.nonce"`)}
	}
	if _, ok := wc.mutation.Fee(); !ok {
		return &ValidationError{Name: "fee", err: errors.New(`ent: missing required field "Withdraw.fee"`)}
	}
	if _, ok := wc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "Withdraw.time"`)}
	}
	if _, ok := wc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "Withdraw.address"`)}
	}
	return nil
}

func (wc *WithdrawCreate) sqlSave(ctx context.Context) (*Withdraw, error) {
	if err := wc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	wc.mutation.id = &_node.ID
	wc.mutation.done = true
	return _node, nil
}

func (wc *WithdrawCreate) createSpec() (*Withdraw, *sqlgraph.CreateSpec) {
	var (
		_node = &Withdraw{config: wc.config}
		_spec = sqlgraph.NewCreateSpec(withdraw.Table, sqlgraph.NewFieldSpec(withdraw.FieldID, field.TypeInt))
	)
	_spec.OnConflict = wc.conflict
	if value, ok := wc.mutation.Usdc(); ok {
		_spec.SetField(withdraw.FieldUsdc, field.TypeString, value)
		_node.Usdc = value
	}
	if value, ok := wc.mutation.Nonce(); ok {
		_spec.SetField(withdraw.FieldNonce, field.TypeInt64, value)
		_node.Nonce = value
	}
	if value, ok := wc.mutation.Fee(); ok {
		_spec.SetField(withdraw.FieldFee, field.TypeString, value)
		_node.Fee = value
	}
	if value, ok := wc.mutation.Time(); ok {
		_spec.SetField(withdraw.FieldTime, field.TypeInt64, value)
		_node.Time = value
	}
	if value, ok := wc.mutation.Address(); ok {
		_spec.SetField(withdraw.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Withdraw.Create().
//		SetUsdc(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WithdrawUpsert) {
//			SetUsdc(v+v).
//		}).
//		Exec(ctx)
func (wc *WithdrawCreate) OnConflict(opts ...sql.ConflictOption) *WithdrawUpsertOne {
	wc.conflict = opts
	return &WithdrawUpsertOne{
		create: wc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Withdraw.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wc *WithdrawCreate) OnConflictColumns(columns ...string) *WithdrawUpsertOne {
	wc.conflict = append(wc.conflict, sql.ConflictColumns(columns...))
	return &WithdrawUpsertOne{
		create: wc,
	}
}

type (
	// WithdrawUpsertOne is the builder for "upsert"-ing
	//  one Withdraw node.
	WithdrawUpsertOne struct {
		create *WithdrawCreate
	}

	// WithdrawUpsert is the "OnConflict" setter.
	WithdrawUpsert struct {
		*sql.UpdateSet
	}
)

// SetUsdc sets the "usdc" field.
func (u *WithdrawUpsert) SetUsdc(v string) *WithdrawUpsert {
	u.Set(withdraw.FieldUsdc, v)
	return u
}

// UpdateUsdc sets the "usdc" field to the value that was provided on create.
func (u *WithdrawUpsert) UpdateUsdc() *WithdrawUpsert {
	u.SetExcluded(withdraw.FieldUsdc)
	return u
}

// SetNonce sets the "nonce" field.
func (u *WithdrawUpsert) SetNonce(v int64) *WithdrawUpsert {
	u.Set(withdraw.FieldNonce, v)
	return u
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *WithdrawUpsert) UpdateNonce() *WithdrawUpsert {
	u.SetExcluded(withdraw.FieldNonce)
	return u
}

// AddNonce adds v to the "nonce" field.
func (u *WithdrawUpsert) AddNonce(v int64) *WithdrawUpsert {
	u.Add(withdraw.FieldNonce, v)
	return u
}

// SetFee sets the "fee" field.
func (u *WithdrawUpsert) SetFee(v string) *WithdrawUpsert {
	u.Set(withdraw.FieldFee, v)
	return u
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *WithdrawUpsert) UpdateFee() *WithdrawUpsert {
	u.SetExcluded(withdraw.FieldFee)
	return u
}

// SetTime sets the "time" field.
func (u *WithdrawUpsert) SetTime(v int64) *WithdrawUpsert {
	u.Set(withdraw.FieldTime, v)
	return u
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *WithdrawUpsert) UpdateTime() *WithdrawUpsert {
	u.SetExcluded(withdraw.FieldTime)
	return u
}

// AddTime adds v to the "time" field.
func (u *WithdrawUpsert) AddTime(v int64) *WithdrawUpsert {
	u.Add(withdraw.FieldTime, v)
	return u
}

// SetAddress sets the "address" field.
func (u *WithdrawUpsert) SetAddress(v string) *WithdrawUpsert {
	u.Set(withdraw.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *WithdrawUpsert) UpdateAddress() *WithdrawUpsert {
	u.SetExcluded(withdraw.FieldAddress)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Withdraw.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *WithdrawUpsertOne) UpdateNewValues() *WithdrawUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Withdraw.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *WithdrawUpsertOne) Ignore() *WithdrawUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WithdrawUpsertOne) DoNothing() *WithdrawUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WithdrawCreate.OnConflict
// documentation for more info.
func (u *WithdrawUpsertOne) Update(set func(*WithdrawUpsert)) *WithdrawUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WithdrawUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsdc sets the "usdc" field.
func (u *WithdrawUpsertOne) SetUsdc(v string) *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetUsdc(v)
	})
}

// UpdateUsdc sets the "usdc" field to the value that was provided on create.
func (u *WithdrawUpsertOne) UpdateUsdc() *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateUsdc()
	})
}

// SetNonce sets the "nonce" field.
func (u *WithdrawUpsertOne) SetNonce(v int64) *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetNonce(v)
	})
}

// AddNonce adds v to the "nonce" field.
func (u *WithdrawUpsertOne) AddNonce(v int64) *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.AddNonce(v)
	})
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *WithdrawUpsertOne) UpdateNonce() *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateNonce()
	})
}

// SetFee sets the "fee" field.
func (u *WithdrawUpsertOne) SetFee(v string) *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetFee(v)
	})
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *WithdrawUpsertOne) UpdateFee() *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateFee()
	})
}

// SetTime sets the "time" field.
func (u *WithdrawUpsertOne) SetTime(v int64) *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetTime(v)
	})
}

// AddTime adds v to the "time" field.
func (u *WithdrawUpsertOne) AddTime(v int64) *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.AddTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *WithdrawUpsertOne) UpdateTime() *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateTime()
	})
}

// SetAddress sets the "address" field.
func (u *WithdrawUpsertOne) SetAddress(v string) *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *WithdrawUpsertOne) UpdateAddress() *WithdrawUpsertOne {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateAddress()
	})
}

// Exec executes the query.
func (u *WithdrawUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WithdrawCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WithdrawUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *WithdrawUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *WithdrawUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// WithdrawCreateBulk is the builder for creating many Withdraw entities in bulk.
type WithdrawCreateBulk struct {
	config
	err      error
	builders []*WithdrawCreate
	conflict []sql.ConflictOption
}

// Save creates the Withdraw entities in the database.
func (wcb *WithdrawCreateBulk) Save(ctx context.Context) ([]*Withdraw, error) {
	if wcb.err != nil {
		return nil, wcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(wcb.builders))
	nodes := make([]*Withdraw, len(wcb.builders))
	mutators := make([]Mutator, len(wcb.builders))
	for i := range wcb.builders {
		func(i int, root context.Context) {
			builder := wcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WithdrawMutation)
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
					_, err = mutators[i+1].Mutate(root, wcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = wcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wcb *WithdrawCreateBulk) SaveX(ctx context.Context) []*Withdraw {
	v, err := wcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wcb *WithdrawCreateBulk) Exec(ctx context.Context) error {
	_, err := wcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wcb *WithdrawCreateBulk) ExecX(ctx context.Context) {
	if err := wcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Withdraw.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.WithdrawUpsert) {
//			SetUsdc(v+v).
//		}).
//		Exec(ctx)
func (wcb *WithdrawCreateBulk) OnConflict(opts ...sql.ConflictOption) *WithdrawUpsertBulk {
	wcb.conflict = opts
	return &WithdrawUpsertBulk{
		create: wcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Withdraw.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (wcb *WithdrawCreateBulk) OnConflictColumns(columns ...string) *WithdrawUpsertBulk {
	wcb.conflict = append(wcb.conflict, sql.ConflictColumns(columns...))
	return &WithdrawUpsertBulk{
		create: wcb,
	}
}

// WithdrawUpsertBulk is the builder for "upsert"-ing
// a bulk of Withdraw nodes.
type WithdrawUpsertBulk struct {
	create *WithdrawCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Withdraw.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *WithdrawUpsertBulk) UpdateNewValues() *WithdrawUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Withdraw.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *WithdrawUpsertBulk) Ignore() *WithdrawUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *WithdrawUpsertBulk) DoNothing() *WithdrawUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the WithdrawCreateBulk.OnConflict
// documentation for more info.
func (u *WithdrawUpsertBulk) Update(set func(*WithdrawUpsert)) *WithdrawUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&WithdrawUpsert{UpdateSet: update})
	}))
	return u
}

// SetUsdc sets the "usdc" field.
func (u *WithdrawUpsertBulk) SetUsdc(v string) *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetUsdc(v)
	})
}

// UpdateUsdc sets the "usdc" field to the value that was provided on create.
func (u *WithdrawUpsertBulk) UpdateUsdc() *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateUsdc()
	})
}

// SetNonce sets the "nonce" field.
func (u *WithdrawUpsertBulk) SetNonce(v int64) *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetNonce(v)
	})
}

// AddNonce adds v to the "nonce" field.
func (u *WithdrawUpsertBulk) AddNonce(v int64) *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.AddNonce(v)
	})
}

// UpdateNonce sets the "nonce" field to the value that was provided on create.
func (u *WithdrawUpsertBulk) UpdateNonce() *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateNonce()
	})
}

// SetFee sets the "fee" field.
func (u *WithdrawUpsertBulk) SetFee(v string) *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetFee(v)
	})
}

// UpdateFee sets the "fee" field to the value that was provided on create.
func (u *WithdrawUpsertBulk) UpdateFee() *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateFee()
	})
}

// SetTime sets the "time" field.
func (u *WithdrawUpsertBulk) SetTime(v int64) *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetTime(v)
	})
}

// AddTime adds v to the "time" field.
func (u *WithdrawUpsertBulk) AddTime(v int64) *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.AddTime(v)
	})
}

// UpdateTime sets the "time" field to the value that was provided on create.
func (u *WithdrawUpsertBulk) UpdateTime() *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateTime()
	})
}

// SetAddress sets the "address" field.
func (u *WithdrawUpsertBulk) SetAddress(v string) *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *WithdrawUpsertBulk) UpdateAddress() *WithdrawUpsertBulk {
	return u.Update(func(s *WithdrawUpsert) {
		s.UpdateAddress()
	})
}

// Exec executes the query.
func (u *WithdrawUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the WithdrawCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for WithdrawCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *WithdrawUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}