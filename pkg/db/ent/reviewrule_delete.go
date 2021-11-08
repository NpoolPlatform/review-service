// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/reviewrule"
)

// ReviewRuleDelete is the builder for deleting a ReviewRule entity.
type ReviewRuleDelete struct {
	config
	hooks    []Hook
	mutation *ReviewRuleMutation
}

// Where appends a list predicates to the ReviewRuleDelete builder.
func (rrd *ReviewRuleDelete) Where(ps ...predicate.ReviewRule) *ReviewRuleDelete {
	rrd.mutation.Where(ps...)
	return rrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rrd *ReviewRuleDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rrd.hooks) == 0 {
		affected, err = rrd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rrd.mutation = mutation
			affected, err = rrd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rrd.hooks) - 1; i >= 0; i-- {
			if rrd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rrd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rrd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rrd *ReviewRuleDelete) ExecX(ctx context.Context) int {
	n, err := rrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rrd *ReviewRuleDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: reviewrule.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: reviewrule.FieldID,
			},
		},
	}
	if ps := rrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, rrd.driver, _spec)
}

// ReviewRuleDeleteOne is the builder for deleting a single ReviewRule entity.
type ReviewRuleDeleteOne struct {
	rrd *ReviewRuleDelete
}

// Exec executes the deletion query.
func (rrdo *ReviewRuleDeleteOne) Exec(ctx context.Context) error {
	n, err := rrdo.rrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{reviewrule.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rrdo *ReviewRuleDeleteOne) ExecX(ctx context.Context) {
	rrdo.rrd.ExecX(ctx)
}
