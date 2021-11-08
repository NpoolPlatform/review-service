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

// ReviewRuleUpdate is the builder for updating ReviewRule entities.
type ReviewRuleUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewRuleMutation
}

// Where appends a list predicates to the ReviewRuleUpdate builder.
func (rru *ReviewRuleUpdate) Where(ps ...predicate.ReviewRule) *ReviewRuleUpdate {
	rru.mutation.Where(ps...)
	return rru
}

// SetEntityType sets the "entity_type" field.
func (rru *ReviewRuleUpdate) SetEntityType(s string) *ReviewRuleUpdate {
	rru.mutation.SetEntityType(s)
	return rru
}

// SetDomain sets the "domain" field.
func (rru *ReviewRuleUpdate) SetDomain(s string) *ReviewRuleUpdate {
	rru.mutation.SetDomain(s)
	return rru
}

// SetRules sets the "rules" field.
func (rru *ReviewRuleUpdate) SetRules(s string) *ReviewRuleUpdate {
	rru.mutation.SetRules(s)
	return rru
}

// SetNillableRules sets the "rules" field if the given value is not nil.
func (rru *ReviewRuleUpdate) SetNillableRules(s *string) *ReviewRuleUpdate {
	if s != nil {
		rru.SetRules(*s)
	}
	return rru
}

// SetCreateAt sets the "create_at" field.
func (rru *ReviewRuleUpdate) SetCreateAt(u uint32) *ReviewRuleUpdate {
	rru.mutation.ResetCreateAt()
	rru.mutation.SetCreateAt(u)
	return rru
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (rru *ReviewRuleUpdate) SetNillableCreateAt(u *uint32) *ReviewRuleUpdate {
	if u != nil {
		rru.SetCreateAt(*u)
	}
	return rru
}

// AddCreateAt adds u to the "create_at" field.
func (rru *ReviewRuleUpdate) AddCreateAt(u uint32) *ReviewRuleUpdate {
	rru.mutation.AddCreateAt(u)
	return rru
}

// SetUpdateAt sets the "update_at" field.
func (rru *ReviewRuleUpdate) SetUpdateAt(u uint32) *ReviewRuleUpdate {
	rru.mutation.ResetUpdateAt()
	rru.mutation.SetUpdateAt(u)
	return rru
}

// AddUpdateAt adds u to the "update_at" field.
func (rru *ReviewRuleUpdate) AddUpdateAt(u uint32) *ReviewRuleUpdate {
	rru.mutation.AddUpdateAt(u)
	return rru
}

// SetDeleteAt sets the "delete_at" field.
func (rru *ReviewRuleUpdate) SetDeleteAt(u uint32) *ReviewRuleUpdate {
	rru.mutation.ResetDeleteAt()
	rru.mutation.SetDeleteAt(u)
	return rru
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (rru *ReviewRuleUpdate) SetNillableDeleteAt(u *uint32) *ReviewRuleUpdate {
	if u != nil {
		rru.SetDeleteAt(*u)
	}
	return rru
}

// AddDeleteAt adds u to the "delete_at" field.
func (rru *ReviewRuleUpdate) AddDeleteAt(u uint32) *ReviewRuleUpdate {
	rru.mutation.AddDeleteAt(u)
	return rru
}

// Mutation returns the ReviewRuleMutation object of the builder.
func (rru *ReviewRuleUpdate) Mutation() *ReviewRuleMutation {
	return rru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rru *ReviewRuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	rru.defaults()
	if len(rru.hooks) == 0 {
		affected, err = rru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rru.mutation = mutation
			affected, err = rru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rru.hooks) - 1; i >= 0; i-- {
			if rru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rru *ReviewRuleUpdate) SaveX(ctx context.Context) int {
	affected, err := rru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rru *ReviewRuleUpdate) Exec(ctx context.Context) error {
	_, err := rru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rru *ReviewRuleUpdate) ExecX(ctx context.Context) {
	if err := rru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rru *ReviewRuleUpdate) defaults() {
	if _, ok := rru.mutation.UpdateAt(); !ok {
		v := reviewrule.UpdateDefaultUpdateAt()
		rru.mutation.SetUpdateAt(v)
	}
}

func (rru *ReviewRuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reviewrule.Table,
			Columns: reviewrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: reviewrule.FieldID,
			},
		},
	}
	if ps := rru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rru.mutation.EntityType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewrule.FieldEntityType,
		})
	}
	if value, ok := rru.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewrule.FieldDomain,
		})
	}
	if value, ok := rru.mutation.Rules(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewrule.FieldRules,
		})
	}
	if value, ok := rru.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldCreateAt,
		})
	}
	if value, ok := rru.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldCreateAt,
		})
	}
	if value, ok := rru.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldUpdateAt,
		})
	}
	if value, ok := rru.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldUpdateAt,
		})
	}
	if value, ok := rru.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldDeleteAt,
		})
	}
	if value, ok := rru.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reviewrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReviewRuleUpdateOne is the builder for updating a single ReviewRule entity.
type ReviewRuleUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewRuleMutation
}

// SetEntityType sets the "entity_type" field.
func (rruo *ReviewRuleUpdateOne) SetEntityType(s string) *ReviewRuleUpdateOne {
	rruo.mutation.SetEntityType(s)
	return rruo
}

// SetDomain sets the "domain" field.
func (rruo *ReviewRuleUpdateOne) SetDomain(s string) *ReviewRuleUpdateOne {
	rruo.mutation.SetDomain(s)
	return rruo
}

// SetRules sets the "rules" field.
func (rruo *ReviewRuleUpdateOne) SetRules(s string) *ReviewRuleUpdateOne {
	rruo.mutation.SetRules(s)
	return rruo
}

// SetNillableRules sets the "rules" field if the given value is not nil.
func (rruo *ReviewRuleUpdateOne) SetNillableRules(s *string) *ReviewRuleUpdateOne {
	if s != nil {
		rruo.SetRules(*s)
	}
	return rruo
}

// SetCreateAt sets the "create_at" field.
func (rruo *ReviewRuleUpdateOne) SetCreateAt(u uint32) *ReviewRuleUpdateOne {
	rruo.mutation.ResetCreateAt()
	rruo.mutation.SetCreateAt(u)
	return rruo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (rruo *ReviewRuleUpdateOne) SetNillableCreateAt(u *uint32) *ReviewRuleUpdateOne {
	if u != nil {
		rruo.SetCreateAt(*u)
	}
	return rruo
}

// AddCreateAt adds u to the "create_at" field.
func (rruo *ReviewRuleUpdateOne) AddCreateAt(u uint32) *ReviewRuleUpdateOne {
	rruo.mutation.AddCreateAt(u)
	return rruo
}

// SetUpdateAt sets the "update_at" field.
func (rruo *ReviewRuleUpdateOne) SetUpdateAt(u uint32) *ReviewRuleUpdateOne {
	rruo.mutation.ResetUpdateAt()
	rruo.mutation.SetUpdateAt(u)
	return rruo
}

// AddUpdateAt adds u to the "update_at" field.
func (rruo *ReviewRuleUpdateOne) AddUpdateAt(u uint32) *ReviewRuleUpdateOne {
	rruo.mutation.AddUpdateAt(u)
	return rruo
}

// SetDeleteAt sets the "delete_at" field.
func (rruo *ReviewRuleUpdateOne) SetDeleteAt(u uint32) *ReviewRuleUpdateOne {
	rruo.mutation.ResetDeleteAt()
	rruo.mutation.SetDeleteAt(u)
	return rruo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (rruo *ReviewRuleUpdateOne) SetNillableDeleteAt(u *uint32) *ReviewRuleUpdateOne {
	if u != nil {
		rruo.SetDeleteAt(*u)
	}
	return rruo
}

// AddDeleteAt adds u to the "delete_at" field.
func (rruo *ReviewRuleUpdateOne) AddDeleteAt(u uint32) *ReviewRuleUpdateOne {
	rruo.mutation.AddDeleteAt(u)
	return rruo
}

// Mutation returns the ReviewRuleMutation object of the builder.
func (rruo *ReviewRuleUpdateOne) Mutation() *ReviewRuleMutation {
	return rruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rruo *ReviewRuleUpdateOne) Select(field string, fields ...string) *ReviewRuleUpdateOne {
	rruo.fields = append([]string{field}, fields...)
	return rruo
}

// Save executes the query and returns the updated ReviewRule entity.
func (rruo *ReviewRuleUpdateOne) Save(ctx context.Context) (*ReviewRule, error) {
	var (
		err  error
		node *ReviewRule
	)
	rruo.defaults()
	if len(rruo.hooks) == 0 {
		node, err = rruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewRuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rruo.mutation = mutation
			node, err = rruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rruo.hooks) - 1; i >= 0; i-- {
			if rruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rruo *ReviewRuleUpdateOne) SaveX(ctx context.Context) *ReviewRule {
	node, err := rruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rruo *ReviewRuleUpdateOne) Exec(ctx context.Context) error {
	_, err := rruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rruo *ReviewRuleUpdateOne) ExecX(ctx context.Context) {
	if err := rruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rruo *ReviewRuleUpdateOne) defaults() {
	if _, ok := rruo.mutation.UpdateAt(); !ok {
		v := reviewrule.UpdateDefaultUpdateAt()
		rruo.mutation.SetUpdateAt(v)
	}
}

func (rruo *ReviewRuleUpdateOne) sqlSave(ctx context.Context) (_node *ReviewRule, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reviewrule.Table,
			Columns: reviewrule.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: reviewrule.FieldID,
			},
		},
	}
	id, ok := rruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing ReviewRule.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := rruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reviewrule.FieldID)
		for _, f := range fields {
			if !reviewrule.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reviewrule.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rruo.mutation.EntityType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewrule.FieldEntityType,
		})
	}
	if value, ok := rruo.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewrule.FieldDomain,
		})
	}
	if value, ok := rruo.mutation.Rules(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reviewrule.FieldRules,
		})
	}
	if value, ok := rruo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldCreateAt,
		})
	}
	if value, ok := rruo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldCreateAt,
		})
	}
	if value, ok := rruo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldUpdateAt,
		})
	}
	if value, ok := rruo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldUpdateAt,
		})
	}
	if value, ok := rruo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldDeleteAt,
		})
	}
	if value, ok := rruo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: reviewrule.FieldDeleteAt,
		})
	}
	_node = &ReviewRule{config: rruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reviewrule.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}