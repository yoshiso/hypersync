// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/predicate"
	"github.com/yoshiso/hypersync/ent/vaultdelta"
)

// VaultDeltaDelete is the builder for deleting a VaultDelta entity.
type VaultDeltaDelete struct {
	config
	hooks    []Hook
	mutation *VaultDeltaMutation
}

// Where appends a list predicates to the VaultDeltaDelete builder.
func (vdd *VaultDeltaDelete) Where(ps ...predicate.VaultDelta) *VaultDeltaDelete {
	vdd.mutation.Where(ps...)
	return vdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vdd *VaultDeltaDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, vdd.sqlExec, vdd.mutation, vdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (vdd *VaultDeltaDelete) ExecX(ctx context.Context) int {
	n, err := vdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vdd *VaultDeltaDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(vaultdelta.Table, sqlgraph.NewFieldSpec(vaultdelta.FieldID, field.TypeInt))
	if ps := vdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, vdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	vdd.mutation.done = true
	return affected, err
}

// VaultDeltaDeleteOne is the builder for deleting a single VaultDelta entity.
type VaultDeltaDeleteOne struct {
	vdd *VaultDeltaDelete
}

// Where appends a list predicates to the VaultDeltaDelete builder.
func (vddo *VaultDeltaDeleteOne) Where(ps ...predicate.VaultDelta) *VaultDeltaDeleteOne {
	vddo.vdd.mutation.Where(ps...)
	return vddo
}

// Exec executes the deletion query.
func (vddo *VaultDeltaDeleteOne) Exec(ctx context.Context) error {
	n, err := vddo.vdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vaultdelta.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vddo *VaultDeltaDeleteOne) ExecX(ctx context.Context) {
	if err := vddo.Exec(ctx); err != nil {
		panic(err)
	}
}