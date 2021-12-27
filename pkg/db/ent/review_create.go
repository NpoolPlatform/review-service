// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"
	"github.com/google/uuid"
)

// ReviewCreate is the builder for creating a Review entity.
type ReviewCreate struct {
	config
	mutation *ReviewMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetObjectType sets the "object_type" field.
func (rc *ReviewCreate) SetObjectType(s string) *ReviewCreate {
	rc.mutation.SetObjectType(s)
	return rc
}

// SetDomain sets the "domain" field.
func (rc *ReviewCreate) SetDomain(s string) *ReviewCreate {
	rc.mutation.SetDomain(s)
	return rc
}

// SetObjectID sets the "object_id" field.
func (rc *ReviewCreate) SetObjectID(u uuid.UUID) *ReviewCreate {
	rc.mutation.SetObjectID(u)
	return rc
}

// SetReviewerID sets the "reviewer_id" field.
func (rc *ReviewCreate) SetReviewerID(u uuid.UUID) *ReviewCreate {
	rc.mutation.SetReviewerID(u)
	return rc
}

// SetState sets the "state" field.
func (rc *ReviewCreate) SetState(r review.State) *ReviewCreate {
	rc.mutation.SetState(r)
	return rc
}

// SetMessage sets the "message" field.
func (rc *ReviewCreate) SetMessage(s string) *ReviewCreate {
	rc.mutation.SetMessage(s)
	return rc
}

// SetCreateAt sets the "create_at" field.
func (rc *ReviewCreate) SetCreateAt(u uint32) *ReviewCreate {
	rc.mutation.SetCreateAt(u)
	return rc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableCreateAt(u *uint32) *ReviewCreate {
	if u != nil {
		rc.SetCreateAt(*u)
	}
	return rc
}

// SetUpdateAt sets the "update_at" field.
func (rc *ReviewCreate) SetUpdateAt(u uint32) *ReviewCreate {
	rc.mutation.SetUpdateAt(u)
	return rc
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableUpdateAt(u *uint32) *ReviewCreate {
	if u != nil {
		rc.SetUpdateAt(*u)
	}
	return rc
}

// SetDeleteAt sets the "delete_at" field.
func (rc *ReviewCreate) SetDeleteAt(u uint32) *ReviewCreate {
	rc.mutation.SetDeleteAt(u)
	return rc
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (rc *ReviewCreate) SetNillableDeleteAt(u *uint32) *ReviewCreate {
	if u != nil {
		rc.SetDeleteAt(*u)
	}
	return rc
}

// SetID sets the "id" field.
func (rc *ReviewCreate) SetID(u uuid.UUID) *ReviewCreate {
	rc.mutation.SetID(u)
	return rc
}

// Mutation returns the ReviewMutation object of the builder.
func (rc *ReviewCreate) Mutation() *ReviewMutation {
	return rc.mutation
}

// Save creates the Review in the database.
func (rc *ReviewCreate) Save(ctx context.Context) (*Review, error) {
	var (
		err  error
		node *Review
	)
	rc.defaults()
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReviewMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReviewCreate) SaveX(ctx context.Context) *Review {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReviewCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReviewCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReviewCreate) defaults() {
	if _, ok := rc.mutation.CreateAt(); !ok {
		v := review.DefaultCreateAt()
		rc.mutation.SetCreateAt(v)
	}
	if _, ok := rc.mutation.UpdateAt(); !ok {
		v := review.DefaultUpdateAt()
		rc.mutation.SetUpdateAt(v)
	}
	if _, ok := rc.mutation.DeleteAt(); !ok {
		v := review.DefaultDeleteAt()
		rc.mutation.SetDeleteAt(v)
	}
	if _, ok := rc.mutation.ID(); !ok {
		v := review.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReviewCreate) check() error {
	if _, ok := rc.mutation.ObjectType(); !ok {
		return &ValidationError{Name: "object_type", err: errors.New(`ent: missing required field "object_type"`)}
	}
	if _, ok := rc.mutation.Domain(); !ok {
		return &ValidationError{Name: "domain", err: errors.New(`ent: missing required field "domain"`)}
	}
	if _, ok := rc.mutation.ObjectID(); !ok {
		return &ValidationError{Name: "object_id", err: errors.New(`ent: missing required field "object_id"`)}
	}
	if _, ok := rc.mutation.ReviewerID(); !ok {
		return &ValidationError{Name: "reviewer_id", err: errors.New(`ent: missing required field "reviewer_id"`)}
	}
	if _, ok := rc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "state"`)}
	}
	if v, ok := rc.mutation.State(); ok {
		if err := review.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "state": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Message(); !ok {
		return &ValidationError{Name: "message", err: errors.New(`ent: missing required field "message"`)}
	}
	if _, ok := rc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "create_at"`)}
	}
	if _, ok := rc.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "update_at"`)}
	}
	if _, ok := rc.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "delete_at"`)}
	}
	return nil
}

func (rc *ReviewCreate) sqlSave(ctx context.Context) (*Review, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (rc *ReviewCreate) createSpec() (*Review, *sqlgraph.CreateSpec) {
	var (
		_node = &Review{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: review.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: review.FieldID,
			},
		}
	)
	_spec.OnConflict = rc.conflict
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.ObjectType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldObjectType,
		})
		_node.ObjectType = value
	}
	if value, ok := rc.mutation.Domain(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldDomain,
		})
		_node.Domain = value
	}
	if value, ok := rc.mutation.ObjectID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: review.FieldObjectID,
		})
		_node.ObjectID = value
	}
	if value, ok := rc.mutation.ReviewerID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: review.FieldReviewerID,
		})
		_node.ReviewerID = value
	}
	if value, ok := rc.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: review.FieldState,
		})
		_node.State = value
	}
	if value, ok := rc.mutation.Message(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: review.FieldMessage,
		})
		_node.Message = value
	}
	if value, ok := rc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := rc.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := rc.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: review.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Review.Create().
//		SetObjectType(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReviewUpsert) {
//			SetObjectType(v+v).
//		}).
//		Exec(ctx)
//
func (rc *ReviewCreate) OnConflict(opts ...sql.ConflictOption) *ReviewUpsertOne {
	rc.conflict = opts
	return &ReviewUpsertOne{
		create: rc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Review.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rc *ReviewCreate) OnConflictColumns(columns ...string) *ReviewUpsertOne {
	rc.conflict = append(rc.conflict, sql.ConflictColumns(columns...))
	return &ReviewUpsertOne{
		create: rc,
	}
}

type (
	// ReviewUpsertOne is the builder for "upsert"-ing
	//  one Review node.
	ReviewUpsertOne struct {
		create *ReviewCreate
	}

	// ReviewUpsert is the "OnConflict" setter.
	ReviewUpsert struct {
		*sql.UpdateSet
	}
)

// SetObjectType sets the "object_type" field.
func (u *ReviewUpsert) SetObjectType(v string) *ReviewUpsert {
	u.Set(review.FieldObjectType, v)
	return u
}

// UpdateObjectType sets the "object_type" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateObjectType() *ReviewUpsert {
	u.SetExcluded(review.FieldObjectType)
	return u
}

// SetDomain sets the "domain" field.
func (u *ReviewUpsert) SetDomain(v string) *ReviewUpsert {
	u.Set(review.FieldDomain, v)
	return u
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateDomain() *ReviewUpsert {
	u.SetExcluded(review.FieldDomain)
	return u
}

// SetObjectID sets the "object_id" field.
func (u *ReviewUpsert) SetObjectID(v uuid.UUID) *ReviewUpsert {
	u.Set(review.FieldObjectID, v)
	return u
}

// UpdateObjectID sets the "object_id" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateObjectID() *ReviewUpsert {
	u.SetExcluded(review.FieldObjectID)
	return u
}

// SetReviewerID sets the "reviewer_id" field.
func (u *ReviewUpsert) SetReviewerID(v uuid.UUID) *ReviewUpsert {
	u.Set(review.FieldReviewerID, v)
	return u
}

// UpdateReviewerID sets the "reviewer_id" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateReviewerID() *ReviewUpsert {
	u.SetExcluded(review.FieldReviewerID)
	return u
}

// SetState sets the "state" field.
func (u *ReviewUpsert) SetState(v review.State) *ReviewUpsert {
	u.Set(review.FieldState, v)
	return u
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateState() *ReviewUpsert {
	u.SetExcluded(review.FieldState)
	return u
}

// SetMessage sets the "message" field.
func (u *ReviewUpsert) SetMessage(v string) *ReviewUpsert {
	u.Set(review.FieldMessage, v)
	return u
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateMessage() *ReviewUpsert {
	u.SetExcluded(review.FieldMessage)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *ReviewUpsert) SetCreateAt(v uint32) *ReviewUpsert {
	u.Set(review.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateCreateAt() *ReviewUpsert {
	u.SetExcluded(review.FieldCreateAt)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *ReviewUpsert) SetUpdateAt(v uint32) *ReviewUpsert {
	u.Set(review.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateUpdateAt() *ReviewUpsert {
	u.SetExcluded(review.FieldUpdateAt)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *ReviewUpsert) SetDeleteAt(v uint32) *ReviewUpsert {
	u.Set(review.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ReviewUpsert) UpdateDeleteAt() *ReviewUpsert {
	u.SetExcluded(review.FieldDeleteAt)
	return u
}

// UpdateNewValues updates the fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Review.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(review.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ReviewUpsertOne) UpdateNewValues() *ReviewUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(review.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Review.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ReviewUpsertOne) Ignore() *ReviewUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReviewUpsertOne) DoNothing() *ReviewUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReviewCreate.OnConflict
// documentation for more info.
func (u *ReviewUpsertOne) Update(set func(*ReviewUpsert)) *ReviewUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReviewUpsert{UpdateSet: update})
	}))
	return u
}

// SetObjectType sets the "object_type" field.
func (u *ReviewUpsertOne) SetObjectType(v string) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetObjectType(v)
	})
}

// UpdateObjectType sets the "object_type" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateObjectType() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateObjectType()
	})
}

// SetDomain sets the "domain" field.
func (u *ReviewUpsertOne) SetDomain(v string) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetDomain(v)
	})
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateDomain() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateDomain()
	})
}

// SetObjectID sets the "object_id" field.
func (u *ReviewUpsertOne) SetObjectID(v uuid.UUID) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetObjectID(v)
	})
}

// UpdateObjectID sets the "object_id" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateObjectID() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateObjectID()
	})
}

// SetReviewerID sets the "reviewer_id" field.
func (u *ReviewUpsertOne) SetReviewerID(v uuid.UUID) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetReviewerID(v)
	})
}

// UpdateReviewerID sets the "reviewer_id" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateReviewerID() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateReviewerID()
	})
}

// SetState sets the "state" field.
func (u *ReviewUpsertOne) SetState(v review.State) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateState() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateState()
	})
}

// SetMessage sets the "message" field.
func (u *ReviewUpsertOne) SetMessage(v string) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateMessage() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateMessage()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *ReviewUpsertOne) SetCreateAt(v uint32) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateCreateAt() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *ReviewUpsertOne) SetUpdateAt(v uint32) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateUpdateAt() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *ReviewUpsertOne) SetDeleteAt(v uint32) *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ReviewUpsertOne) UpdateDeleteAt() *ReviewUpsertOne {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *ReviewUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReviewCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReviewUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ReviewUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: ReviewUpsertOne.ID is not supported by MySQL driver. Use ReviewUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ReviewUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ReviewCreateBulk is the builder for creating many Review entities in bulk.
type ReviewCreateBulk struct {
	config
	builders []*ReviewCreate
	conflict []sql.ConflictOption
}

// Save creates the Review entities in the database.
func (rcb *ReviewCreateBulk) Save(ctx context.Context) ([]*Review, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Review, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReviewMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReviewCreateBulk) SaveX(ctx context.Context) []*Review {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReviewCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReviewCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Review.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ReviewUpsert) {
//			SetObjectType(v+v).
//		}).
//		Exec(ctx)
//
func (rcb *ReviewCreateBulk) OnConflict(opts ...sql.ConflictOption) *ReviewUpsertBulk {
	rcb.conflict = opts
	return &ReviewUpsertBulk{
		create: rcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Review.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (rcb *ReviewCreateBulk) OnConflictColumns(columns ...string) *ReviewUpsertBulk {
	rcb.conflict = append(rcb.conflict, sql.ConflictColumns(columns...))
	return &ReviewUpsertBulk{
		create: rcb,
	}
}

// ReviewUpsertBulk is the builder for "upsert"-ing
// a bulk of Review nodes.
type ReviewUpsertBulk struct {
	create *ReviewCreateBulk
}

// UpdateNewValues updates the fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Review.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(review.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *ReviewUpsertBulk) UpdateNewValues() *ReviewUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(review.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Review.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ReviewUpsertBulk) Ignore() *ReviewUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ReviewUpsertBulk) DoNothing() *ReviewUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ReviewCreateBulk.OnConflict
// documentation for more info.
func (u *ReviewUpsertBulk) Update(set func(*ReviewUpsert)) *ReviewUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ReviewUpsert{UpdateSet: update})
	}))
	return u
}

// SetObjectType sets the "object_type" field.
func (u *ReviewUpsertBulk) SetObjectType(v string) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetObjectType(v)
	})
}

// UpdateObjectType sets the "object_type" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateObjectType() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateObjectType()
	})
}

// SetDomain sets the "domain" field.
func (u *ReviewUpsertBulk) SetDomain(v string) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetDomain(v)
	})
}

// UpdateDomain sets the "domain" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateDomain() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateDomain()
	})
}

// SetObjectID sets the "object_id" field.
func (u *ReviewUpsertBulk) SetObjectID(v uuid.UUID) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetObjectID(v)
	})
}

// UpdateObjectID sets the "object_id" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateObjectID() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateObjectID()
	})
}

// SetReviewerID sets the "reviewer_id" field.
func (u *ReviewUpsertBulk) SetReviewerID(v uuid.UUID) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetReviewerID(v)
	})
}

// UpdateReviewerID sets the "reviewer_id" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateReviewerID() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateReviewerID()
	})
}

// SetState sets the "state" field.
func (u *ReviewUpsertBulk) SetState(v review.State) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetState(v)
	})
}

// UpdateState sets the "state" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateState() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateState()
	})
}

// SetMessage sets the "message" field.
func (u *ReviewUpsertBulk) SetMessage(v string) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetMessage(v)
	})
}

// UpdateMessage sets the "message" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateMessage() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateMessage()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *ReviewUpsertBulk) SetCreateAt(v uint32) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateCreateAt() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *ReviewUpsertBulk) SetUpdateAt(v uint32) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateUpdateAt() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *ReviewUpsertBulk) SetDeleteAt(v uint32) *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.SetDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *ReviewUpsertBulk) UpdateDeleteAt() *ReviewUpsertBulk {
	return u.Update(func(s *ReviewUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *ReviewUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ReviewCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ReviewCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ReviewUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
