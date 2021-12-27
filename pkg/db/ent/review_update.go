// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"
	"github.com/google/uuid"
)

// ReviewUpdate is the builder for updating Review entities.
type ReviewUpdate struct {
	config
	hooks    []Hook
	mutation *ReviewMutation
}

// Where appends a list predicates to the ReviewUpdate builder.
func (ru *ReviewUpdate) Where(ps ...predicate.Review) *ReviewUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetObjectType sets the "object_type" field.
func (ru *ReviewUpdate) SetObjectType(s string) *ReviewUpdate {
	ru.mutation.SetObjectType(s)
	return ru
}

// SetDomain sets the "domain" field.
func (ru *ReviewUpdate) SetDomain(s string) *ReviewUpdate {
	ru.mutation.SetDomain(s)
	return ru
}

// SetObjectID sets the "object_id" field.
func (ru *ReviewUpdate) SetObjectID(u uuid.UUID) *ReviewUpdate {
	ru.mutation.SetObjectID(u)
	return ru
}

// SetReviewerID sets the "reviewer_id" field.
func (ru *ReviewUpdate) SetReviewerID(u uuid.UUID) *ReviewUpdate {
	ru.mutation.SetReviewerID(u)
	return ru
}

// SetState sets the "state" field.
func (ru *ReviewUpdate) SetState(r review.State) *ReviewUpdate {
	ru.mutation.SetState(r)
	return ru
}

// SetMessage sets the "message" field.
func (ru *ReviewUpdate) SetMessage(s string) *ReviewUpdate {
	ru.mutation.SetMessage(s)
	return ru
}

// SetCreateAt sets the "create_at" field.
func (ru *ReviewUpdate) SetCreateAt(u uint32) *ReviewUpdate {
	ru.mutation.ResetCreateAt()
	ru.mutation.SetCreateAt(u)
	return ru
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableCreateAt(u *uint32) *ReviewUpdate {
	if u != nil {
		ru.SetCreateAt(*u)
	}
	return ru
}

// AddCreateAt adds u to the "create_at" field.
func (ru *ReviewUpdate) AddCreateAt(u uint32) *ReviewUpdate {
	ru.mutation.AddCreateAt(u)
	return ru
}

// SetUpdateAt sets the "update_at" field.
func (ru *ReviewUpdate) SetUpdateAt(u uint32) *ReviewUpdate {
	ru.mutation.ResetUpdateAt()
	ru.mutation.SetUpdateAt(u)
	return ru
}

// AddUpdateAt adds u to the "update_at" field.
func (ru *ReviewUpdate) AddUpdateAt(u uint32) *ReviewUpdate {
	ru.mutation.AddUpdateAt(u)
	return ru
}

// SetDeleteAt sets the "delete_at" field.
func (ru *ReviewUpdate) SetDeleteAt(u uint32) *ReviewUpdate {
	ru.mutation.ResetDeleteAt()
	ru.mutation.SetDeleteAt(u)
	return ru
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ru *ReviewUpdate) SetNillableDeleteAt(u *uint32) *ReviewUpdate {
	if u != nil {
		ru.SetDeleteAt(*u)
	}
	return ru
}

// AddDeleteAt adds u to the "delete_at" field.
func (ru *ReviewUpdate) AddDeleteAt(u uint32) *ReviewUpdate {
	ru.mutation.AddDeleteAt(u)
	return ru
}

// Mutation returns the ReviewMutation object of the builder.
func (ru *ReviewUpdate) Mutation() *ReviewMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReviewUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ru.defaults()
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReviewUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReviewUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReviewUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ru *ReviewUpdate) defaults() {
	if _, ok := ru.mutation.UpdateAt(); !ok {
		v := review.UpdateDefaultUpdateAt()
		ru.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *ReviewUpdate) check() error {
	if v, ok := ru.mutation.State(); ok {
		if err := review.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (ru *ReviewUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: review.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.ObjectType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldObjectType,
		})
	}
	if value, ok := ru.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldDomain,
		})
	}
	if value, ok := ru.mutation.ObjectID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: review.FieldObjectID,
		})
	}
	if value, ok := ru.mutation.ReviewerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: review.FieldReviewerID,
		})
	}
	if value, ok := ru.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: review.FieldState,
		})
	}
	if value, ok := ru.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldMessage,
		})
	}
	if value, ok := ru.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldCreateAt,
		})
	}
	if value, ok := ru.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldCreateAt,
		})
	}
	if value, ok := ru.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldUpdateAt,
		})
	}
	if value, ok := ru.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldUpdateAt,
		})
	}
	if value, ok := ru.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldDeleteAt,
		})
	}
	if value, ok := ru.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldDeleteAt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReviewUpdateOne is the builder for updating a single Review entity.
type ReviewUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReviewMutation
}

// SetObjectType sets the "object_type" field.
func (ruo *ReviewUpdateOne) SetObjectType(s string) *ReviewUpdateOne {
	ruo.mutation.SetObjectType(s)
	return ruo
}

// SetDomain sets the "domain" field.
func (ruo *ReviewUpdateOne) SetDomain(s string) *ReviewUpdateOne {
	ruo.mutation.SetDomain(s)
	return ruo
}

// SetObjectID sets the "object_id" field.
func (ruo *ReviewUpdateOne) SetObjectID(u uuid.UUID) *ReviewUpdateOne {
	ruo.mutation.SetObjectID(u)
	return ruo
}

// SetReviewerID sets the "reviewer_id" field.
func (ruo *ReviewUpdateOne) SetReviewerID(u uuid.UUID) *ReviewUpdateOne {
	ruo.mutation.SetReviewerID(u)
	return ruo
}

// SetState sets the "state" field.
func (ruo *ReviewUpdateOne) SetState(r review.State) *ReviewUpdateOne {
	ruo.mutation.SetState(r)
	return ruo
}

// SetMessage sets the "message" field.
func (ruo *ReviewUpdateOne) SetMessage(s string) *ReviewUpdateOne {
	ruo.mutation.SetMessage(s)
	return ruo
}

// SetCreateAt sets the "create_at" field.
func (ruo *ReviewUpdateOne) SetCreateAt(u uint32) *ReviewUpdateOne {
	ruo.mutation.ResetCreateAt()
	ruo.mutation.SetCreateAt(u)
	return ruo
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableCreateAt(u *uint32) *ReviewUpdateOne {
	if u != nil {
		ruo.SetCreateAt(*u)
	}
	return ruo
}

// AddCreateAt adds u to the "create_at" field.
func (ruo *ReviewUpdateOne) AddCreateAt(u uint32) *ReviewUpdateOne {
	ruo.mutation.AddCreateAt(u)
	return ruo
}

// SetUpdateAt sets the "update_at" field.
func (ruo *ReviewUpdateOne) SetUpdateAt(u uint32) *ReviewUpdateOne {
	ruo.mutation.ResetUpdateAt()
	ruo.mutation.SetUpdateAt(u)
	return ruo
}

// AddUpdateAt adds u to the "update_at" field.
func (ruo *ReviewUpdateOne) AddUpdateAt(u uint32) *ReviewUpdateOne {
	ruo.mutation.AddUpdateAt(u)
	return ruo
}

// SetDeleteAt sets the "delete_at" field.
func (ruo *ReviewUpdateOne) SetDeleteAt(u uint32) *ReviewUpdateOne {
	ruo.mutation.ResetDeleteAt()
	ruo.mutation.SetDeleteAt(u)
	return ruo
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (ruo *ReviewUpdateOne) SetNillableDeleteAt(u *uint32) *ReviewUpdateOne {
	if u != nil {
		ruo.SetDeleteAt(*u)
	}
	return ruo
}

// AddDeleteAt adds u to the "delete_at" field.
func (ruo *ReviewUpdateOne) AddDeleteAt(u uint32) *ReviewUpdateOne {
	ruo.mutation.AddDeleteAt(u)
	return ruo
}

// Mutation returns the ReviewMutation object of the builder.
func (ruo *ReviewUpdateOne) Mutation() *ReviewMutation {
	return ruo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReviewUpdateOne) Select(field string, fields ...string) *ReviewUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Review entity.
func (ruo *ReviewUpdateOne) Save(ctx context.Context) (*Review, error) {
	var (
		err  error
		node *Review
	)
	ruo.defaults()
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReviewUpdateOne) SaveX(ctx context.Context) *Review {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReviewUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReviewUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ruo *ReviewUpdateOne) defaults() {
	if _, ok := ruo.mutation.UpdateAt(); !ok {
		v := review.UpdateDefaultUpdateAt()
		ruo.mutation.SetUpdateAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *ReviewUpdateOne) check() error {
	if v, ok := ruo.mutation.State(); ok {
		if err := review.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf("ent: validator failed for field \"state\": %w", err)}
		}
	}
	return nil
}

func (ruo *ReviewUpdateOne) sqlSave(ctx context.Context) (_node *Review, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   review.Table,
			Columns: review.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: review.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Review.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, review.FieldID)
		for _, f := range fields {
			if !review.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != review.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.ObjectType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldObjectType,
		})
	}
	if value, ok := ruo.mutation.Domain(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldDomain,
		})
	}
	if value, ok := ruo.mutation.ObjectID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: review.FieldObjectID,
		})
	}
	if value, ok := ruo.mutation.ReviewerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: review.FieldReviewerID,
		})
	}
	if value, ok := ruo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: review.FieldState,
		})
	}
	if value, ok := ruo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldMessage,
		})
	}
	if value, ok := ruo.mutation.CreateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldCreateAt,
		})
	}
	if value, ok := ruo.mutation.AddedCreateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldCreateAt,
		})
	}
	if value, ok := ruo.mutation.UpdateAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldUpdateAt,
		})
	}
	if value, ok := ruo.mutation.AddedUpdateAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldUpdateAt,
		})
	}
	if value, ok := ruo.mutation.DeleteAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldDeleteAt,
		})
	}
	if value, ok := ruo.mutation.AddedDeleteAt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldDeleteAt,
		})
	}
	_node = &Review{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{review.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
