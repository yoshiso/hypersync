// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/predicate"
	"github.com/yoshiso/hypersync/ent/spottransfer"
)

// SpotTransferDelete is the builder for deleting a SpotTransfer entity.
type SpotTransferDelete struct {
	config
	hooks    []Hook
	mutation *SpotTransferMutation
}

// Where appends a list predicates to the SpotTransferDelete builder.
func (std *SpotTransferDelete) Where(ps ...predicate.SpotTransfer) *SpotTransferDelete {
	std.mutation.Where(ps...)
	return std
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (std *SpotTransferDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, std.sqlExec, std.mutation, std.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (std *SpotTransferDelete) ExecX(ctx context.Context) int {
	n, err := std.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (std *SpotTransferDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(spottransfer.Table, sqlgraph.NewFieldSpec(spottransfer.FieldID, field.TypeInt))
	if ps := std.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, std.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	std.mutation.done = true
	return affected, err
}

// SpotTransferDeleteOne is the builder for deleting a single SpotTransfer entity.
type SpotTransferDeleteOne struct {
	std *SpotTransferDelete
}

// Where appends a list predicates to the SpotTransferDelete builder.
func (stdo *SpotTransferDeleteOne) Where(ps ...predicate.SpotTransfer) *SpotTransferDeleteOne {
	stdo.std.mutation.Where(ps...)
	return stdo
}

// Exec executes the deletion query.
func (stdo *SpotTransferDeleteOne) Exec(ctx context.Context) error {
	n, err := stdo.std.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{spottransfer.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (stdo *SpotTransferDeleteOne) ExecX(ctx context.Context) {
	if err := stdo.Exec(ctx); err != nil {
		panic(err)
	}
}
