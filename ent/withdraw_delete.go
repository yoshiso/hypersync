// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/predicate"
	"github.com/yoshiso/hypersync/ent/withdraw"
)

// WithdrawDelete is the builder for deleting a Withdraw entity.
type WithdrawDelete struct {
	config
	hooks    []Hook
	mutation *WithdrawMutation
}

// Where appends a list predicates to the WithdrawDelete builder.
func (wd *WithdrawDelete) Where(ps ...predicate.Withdraw) *WithdrawDelete {
	wd.mutation.Where(ps...)
	return wd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (wd *WithdrawDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, wd.sqlExec, wd.mutation, wd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (wd *WithdrawDelete) ExecX(ctx context.Context) int {
	n, err := wd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (wd *WithdrawDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(withdraw.Table, sqlgraph.NewFieldSpec(withdraw.FieldID, field.TypeInt))
	if ps := wd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, wd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	wd.mutation.done = true
	return affected, err
}

// WithdrawDeleteOne is the builder for deleting a single Withdraw entity.
type WithdrawDeleteOne struct {
	wd *WithdrawDelete
}

// Where appends a list predicates to the WithdrawDelete builder.
func (wdo *WithdrawDeleteOne) Where(ps ...predicate.Withdraw) *WithdrawDeleteOne {
	wdo.wd.mutation.Where(ps...)
	return wdo
}

// Exec executes the deletion query.
func (wdo *WithdrawDeleteOne) Exec(ctx context.Context) error {
	n, err := wdo.wd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{withdraw.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (wdo *WithdrawDeleteOne) ExecX(ctx context.Context) {
	if err := wdo.Exec(ctx); err != nil {
		panic(err)
	}
}
