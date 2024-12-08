// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/predicate"
	"github.com/yoshiso/hypersync/ent/spotgenesis"
)

// SpotGenesisDelete is the builder for deleting a SpotGenesis entity.
type SpotGenesisDelete struct {
	config
	hooks    []Hook
	mutation *SpotGenesisMutation
}

// Where appends a list predicates to the SpotGenesisDelete builder.
func (sgd *SpotGenesisDelete) Where(ps ...predicate.SpotGenesis) *SpotGenesisDelete {
	sgd.mutation.Where(ps...)
	return sgd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sgd *SpotGenesisDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sgd.sqlExec, sgd.mutation, sgd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sgd *SpotGenesisDelete) ExecX(ctx context.Context) int {
	n, err := sgd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sgd *SpotGenesisDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(spotgenesis.Table, sqlgraph.NewFieldSpec(spotgenesis.FieldID, field.TypeInt))
	if ps := sgd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sgd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sgd.mutation.done = true
	return affected, err
}

// SpotGenesisDeleteOne is the builder for deleting a single SpotGenesis entity.
type SpotGenesisDeleteOne struct {
	sgd *SpotGenesisDelete
}

// Where appends a list predicates to the SpotGenesisDelete builder.
func (sgdo *SpotGenesisDeleteOne) Where(ps ...predicate.SpotGenesis) *SpotGenesisDeleteOne {
	sgdo.sgd.mutation.Where(ps...)
	return sgdo
}

// Exec executes the deletion query.
func (sgdo *SpotGenesisDeleteOne) Exec(ctx context.Context) error {
	n, err := sgdo.sgd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{spotgenesis.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sgdo *SpotGenesisDeleteOne) ExecX(ctx context.Context) {
	if err := sgdo.Exec(ctx); err != nil {
		panic(err)
	}
}
