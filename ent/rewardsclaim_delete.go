// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/yoshiso/hypersync/ent/predicate"
	"github.com/yoshiso/hypersync/ent/rewardsclaim"
)

// RewardsClaimDelete is the builder for deleting a RewardsClaim entity.
type RewardsClaimDelete struct {
	config
	hooks    []Hook
	mutation *RewardsClaimMutation
}

// Where appends a list predicates to the RewardsClaimDelete builder.
func (rcd *RewardsClaimDelete) Where(ps ...predicate.RewardsClaim) *RewardsClaimDelete {
	rcd.mutation.Where(ps...)
	return rcd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rcd *RewardsClaimDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rcd.sqlExec, rcd.mutation, rcd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rcd *RewardsClaimDelete) ExecX(ctx context.Context) int {
	n, err := rcd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rcd *RewardsClaimDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(rewardsclaim.Table, sqlgraph.NewFieldSpec(rewardsclaim.FieldID, field.TypeInt))
	if ps := rcd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rcd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rcd.mutation.done = true
	return affected, err
}

// RewardsClaimDeleteOne is the builder for deleting a single RewardsClaim entity.
type RewardsClaimDeleteOne struct {
	rcd *RewardsClaimDelete
}

// Where appends a list predicates to the RewardsClaimDelete builder.
func (rcdo *RewardsClaimDeleteOne) Where(ps ...predicate.RewardsClaim) *RewardsClaimDeleteOne {
	rcdo.rcd.mutation.Where(ps...)
	return rcdo
}

// Exec executes the deletion query.
func (rcdo *RewardsClaimDeleteOne) Exec(ctx context.Context) error {
	n, err := rcdo.rcd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{rewardsclaim.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rcdo *RewardsClaimDeleteOne) ExecX(ctx context.Context) {
	if err := rcdo.Exec(ctx); err != nil {
		panic(err)
	}
}
