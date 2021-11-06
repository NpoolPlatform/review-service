// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"

	"github.com/NpoolPlatform/review-service/pkg/db/ent/predicate"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"
	"github.com/google/uuid"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeReview = "Review"
)

// ReviewMutation represents an operation that mutates the Review nodes in the graph.
type ReviewMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	entity_type   *string
	domain        *string
	object_id     *uuid.UUID
	reviewer_id   *uuid.UUID
	state         *review.State
	message       *string
	create_at     *uint32
	addcreate_at  *uint32
	update_at     *uint32
	addupdate_at  *uint32
	delete_at     *uint32
	adddelete_at  *uint32
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Review, error)
	predicates    []predicate.Review
}

var _ ent.Mutation = (*ReviewMutation)(nil)

// reviewOption allows management of the mutation configuration using functional options.
type reviewOption func(*ReviewMutation)

// newReviewMutation creates new mutation for the Review entity.
func newReviewMutation(c config, op Op, opts ...reviewOption) *ReviewMutation {
	m := &ReviewMutation{
		config:        c,
		op:            op,
		typ:           TypeReview,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withReviewID sets the ID field of the mutation.
func withReviewID(id uuid.UUID) reviewOption {
	return func(m *ReviewMutation) {
		var (
			err   error
			once  sync.Once
			value *Review
		)
		m.oldValue = func(ctx context.Context) (*Review, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Review.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withReview sets the old Review of the mutation.
func withReview(node *Review) reviewOption {
	return func(m *ReviewMutation) {
		m.oldValue = func(context.Context) (*Review, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m ReviewMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m ReviewMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of Review entities.
func (m *ReviewMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *ReviewMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetEntityType sets the "entity_type" field.
func (m *ReviewMutation) SetEntityType(s string) {
	m.entity_type = &s
}

// EntityType returns the value of the "entity_type" field in the mutation.
func (m *ReviewMutation) EntityType() (r string, exists bool) {
	v := m.entity_type
	if v == nil {
		return
	}
	return *v, true
}

// OldEntityType returns the old "entity_type" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldEntityType(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldEntityType is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldEntityType requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEntityType: %w", err)
	}
	return oldValue.EntityType, nil
}

// ResetEntityType resets all changes to the "entity_type" field.
func (m *ReviewMutation) ResetEntityType() {
	m.entity_type = nil
}

// SetDomain sets the "domain" field.
func (m *ReviewMutation) SetDomain(s string) {
	m.domain = &s
}

// Domain returns the value of the "domain" field in the mutation.
func (m *ReviewMutation) Domain() (r string, exists bool) {
	v := m.domain
	if v == nil {
		return
	}
	return *v, true
}

// OldDomain returns the old "domain" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldDomain(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDomain is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDomain requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDomain: %w", err)
	}
	return oldValue.Domain, nil
}

// ResetDomain resets all changes to the "domain" field.
func (m *ReviewMutation) ResetDomain() {
	m.domain = nil
}

// SetObjectID sets the "object_id" field.
func (m *ReviewMutation) SetObjectID(u uuid.UUID) {
	m.object_id = &u
}

// ObjectID returns the value of the "object_id" field in the mutation.
func (m *ReviewMutation) ObjectID() (r uuid.UUID, exists bool) {
	v := m.object_id
	if v == nil {
		return
	}
	return *v, true
}

// OldObjectID returns the old "object_id" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldObjectID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldObjectID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldObjectID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldObjectID: %w", err)
	}
	return oldValue.ObjectID, nil
}

// ResetObjectID resets all changes to the "object_id" field.
func (m *ReviewMutation) ResetObjectID() {
	m.object_id = nil
}

// SetReviewerID sets the "reviewer_id" field.
func (m *ReviewMutation) SetReviewerID(u uuid.UUID) {
	m.reviewer_id = &u
}

// ReviewerID returns the value of the "reviewer_id" field in the mutation.
func (m *ReviewMutation) ReviewerID() (r uuid.UUID, exists bool) {
	v := m.reviewer_id
	if v == nil {
		return
	}
	return *v, true
}

// OldReviewerID returns the old "reviewer_id" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldReviewerID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldReviewerID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldReviewerID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldReviewerID: %w", err)
	}
	return oldValue.ReviewerID, nil
}

// ResetReviewerID resets all changes to the "reviewer_id" field.
func (m *ReviewMutation) ResetReviewerID() {
	m.reviewer_id = nil
}

// SetState sets the "state" field.
func (m *ReviewMutation) SetState(r review.State) {
	m.state = &r
}

// State returns the value of the "state" field in the mutation.
func (m *ReviewMutation) State() (r review.State, exists bool) {
	v := m.state
	if v == nil {
		return
	}
	return *v, true
}

// OldState returns the old "state" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldState(ctx context.Context) (v review.State, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldState is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldState requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldState: %w", err)
	}
	return oldValue.State, nil
}

// ResetState resets all changes to the "state" field.
func (m *ReviewMutation) ResetState() {
	m.state = nil
}

// SetMessage sets the "message" field.
func (m *ReviewMutation) SetMessage(s string) {
	m.message = &s
}

// Message returns the value of the "message" field in the mutation.
func (m *ReviewMutation) Message() (r string, exists bool) {
	v := m.message
	if v == nil {
		return
	}
	return *v, true
}

// OldMessage returns the old "message" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldMessage(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldMessage is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldMessage requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldMessage: %w", err)
	}
	return oldValue.Message, nil
}

// ResetMessage resets all changes to the "message" field.
func (m *ReviewMutation) ResetMessage() {
	m.message = nil
}

// SetCreateAt sets the "create_at" field.
func (m *ReviewMutation) SetCreateAt(u uint32) {
	m.create_at = &u
	m.addcreate_at = nil
}

// CreateAt returns the value of the "create_at" field in the mutation.
func (m *ReviewMutation) CreateAt() (r uint32, exists bool) {
	v := m.create_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateAt returns the old "create_at" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldCreateAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateAt: %w", err)
	}
	return oldValue.CreateAt, nil
}

// AddCreateAt adds u to the "create_at" field.
func (m *ReviewMutation) AddCreateAt(u uint32) {
	if m.addcreate_at != nil {
		*m.addcreate_at += u
	} else {
		m.addcreate_at = &u
	}
}

// AddedCreateAt returns the value that was added to the "create_at" field in this mutation.
func (m *ReviewMutation) AddedCreateAt() (r uint32, exists bool) {
	v := m.addcreate_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetCreateAt resets all changes to the "create_at" field.
func (m *ReviewMutation) ResetCreateAt() {
	m.create_at = nil
	m.addcreate_at = nil
}

// SetUpdateAt sets the "update_at" field.
func (m *ReviewMutation) SetUpdateAt(u uint32) {
	m.update_at = &u
	m.addupdate_at = nil
}

// UpdateAt returns the value of the "update_at" field in the mutation.
func (m *ReviewMutation) UpdateAt() (r uint32, exists bool) {
	v := m.update_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdateAt returns the old "update_at" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldUpdateAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUpdateAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUpdateAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdateAt: %w", err)
	}
	return oldValue.UpdateAt, nil
}

// AddUpdateAt adds u to the "update_at" field.
func (m *ReviewMutation) AddUpdateAt(u uint32) {
	if m.addupdate_at != nil {
		*m.addupdate_at += u
	} else {
		m.addupdate_at = &u
	}
}

// AddedUpdateAt returns the value that was added to the "update_at" field in this mutation.
func (m *ReviewMutation) AddedUpdateAt() (r uint32, exists bool) {
	v := m.addupdate_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetUpdateAt resets all changes to the "update_at" field.
func (m *ReviewMutation) ResetUpdateAt() {
	m.update_at = nil
	m.addupdate_at = nil
}

// SetDeleteAt sets the "delete_at" field.
func (m *ReviewMutation) SetDeleteAt(u uint32) {
	m.delete_at = &u
	m.adddelete_at = nil
}

// DeleteAt returns the value of the "delete_at" field in the mutation.
func (m *ReviewMutation) DeleteAt() (r uint32, exists bool) {
	v := m.delete_at
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleteAt returns the old "delete_at" field's value of the Review entity.
// If the Review object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *ReviewMutation) OldDeleteAt(ctx context.Context) (v uint32, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldDeleteAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldDeleteAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleteAt: %w", err)
	}
	return oldValue.DeleteAt, nil
}

// AddDeleteAt adds u to the "delete_at" field.
func (m *ReviewMutation) AddDeleteAt(u uint32) {
	if m.adddelete_at != nil {
		*m.adddelete_at += u
	} else {
		m.adddelete_at = &u
	}
}

// AddedDeleteAt returns the value that was added to the "delete_at" field in this mutation.
func (m *ReviewMutation) AddedDeleteAt() (r uint32, exists bool) {
	v := m.adddelete_at
	if v == nil {
		return
	}
	return *v, true
}

// ResetDeleteAt resets all changes to the "delete_at" field.
func (m *ReviewMutation) ResetDeleteAt() {
	m.delete_at = nil
	m.adddelete_at = nil
}

// Where appends a list predicates to the ReviewMutation builder.
func (m *ReviewMutation) Where(ps ...predicate.Review) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *ReviewMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Review).
func (m *ReviewMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *ReviewMutation) Fields() []string {
	fields := make([]string, 0, 9)
	if m.entity_type != nil {
		fields = append(fields, review.FieldEntityType)
	}
	if m.domain != nil {
		fields = append(fields, review.FieldDomain)
	}
	if m.object_id != nil {
		fields = append(fields, review.FieldObjectID)
	}
	if m.reviewer_id != nil {
		fields = append(fields, review.FieldReviewerID)
	}
	if m.state != nil {
		fields = append(fields, review.FieldState)
	}
	if m.message != nil {
		fields = append(fields, review.FieldMessage)
	}
	if m.create_at != nil {
		fields = append(fields, review.FieldCreateAt)
	}
	if m.update_at != nil {
		fields = append(fields, review.FieldUpdateAt)
	}
	if m.delete_at != nil {
		fields = append(fields, review.FieldDeleteAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *ReviewMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case review.FieldEntityType:
		return m.EntityType()
	case review.FieldDomain:
		return m.Domain()
	case review.FieldObjectID:
		return m.ObjectID()
	case review.FieldReviewerID:
		return m.ReviewerID()
	case review.FieldState:
		return m.State()
	case review.FieldMessage:
		return m.Message()
	case review.FieldCreateAt:
		return m.CreateAt()
	case review.FieldUpdateAt:
		return m.UpdateAt()
	case review.FieldDeleteAt:
		return m.DeleteAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *ReviewMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case review.FieldEntityType:
		return m.OldEntityType(ctx)
	case review.FieldDomain:
		return m.OldDomain(ctx)
	case review.FieldObjectID:
		return m.OldObjectID(ctx)
	case review.FieldReviewerID:
		return m.OldReviewerID(ctx)
	case review.FieldState:
		return m.OldState(ctx)
	case review.FieldMessage:
		return m.OldMessage(ctx)
	case review.FieldCreateAt:
		return m.OldCreateAt(ctx)
	case review.FieldUpdateAt:
		return m.OldUpdateAt(ctx)
	case review.FieldDeleteAt:
		return m.OldDeleteAt(ctx)
	}
	return nil, fmt.Errorf("unknown Review field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ReviewMutation) SetField(name string, value ent.Value) error {
	switch name {
	case review.FieldEntityType:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEntityType(v)
		return nil
	case review.FieldDomain:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDomain(v)
		return nil
	case review.FieldObjectID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetObjectID(v)
		return nil
	case review.FieldReviewerID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetReviewerID(v)
		return nil
	case review.FieldState:
		v, ok := value.(review.State)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetState(v)
		return nil
	case review.FieldMessage:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetMessage(v)
		return nil
	case review.FieldCreateAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateAt(v)
		return nil
	case review.FieldUpdateAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdateAt(v)
		return nil
	case review.FieldDeleteAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleteAt(v)
		return nil
	}
	return fmt.Errorf("unknown Review field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *ReviewMutation) AddedFields() []string {
	var fields []string
	if m.addcreate_at != nil {
		fields = append(fields, review.FieldCreateAt)
	}
	if m.addupdate_at != nil {
		fields = append(fields, review.FieldUpdateAt)
	}
	if m.adddelete_at != nil {
		fields = append(fields, review.FieldDeleteAt)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *ReviewMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case review.FieldCreateAt:
		return m.AddedCreateAt()
	case review.FieldUpdateAt:
		return m.AddedUpdateAt()
	case review.FieldDeleteAt:
		return m.AddedDeleteAt()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *ReviewMutation) AddField(name string, value ent.Value) error {
	switch name {
	case review.FieldCreateAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddCreateAt(v)
		return nil
	case review.FieldUpdateAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUpdateAt(v)
		return nil
	case review.FieldDeleteAt:
		v, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddDeleteAt(v)
		return nil
	}
	return fmt.Errorf("unknown Review numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *ReviewMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *ReviewMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *ReviewMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Review nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *ReviewMutation) ResetField(name string) error {
	switch name {
	case review.FieldEntityType:
		m.ResetEntityType()
		return nil
	case review.FieldDomain:
		m.ResetDomain()
		return nil
	case review.FieldObjectID:
		m.ResetObjectID()
		return nil
	case review.FieldReviewerID:
		m.ResetReviewerID()
		return nil
	case review.FieldState:
		m.ResetState()
		return nil
	case review.FieldMessage:
		m.ResetMessage()
		return nil
	case review.FieldCreateAt:
		m.ResetCreateAt()
		return nil
	case review.FieldUpdateAt:
		m.ResetUpdateAt()
		return nil
	case review.FieldDeleteAt:
		m.ResetDeleteAt()
		return nil
	}
	return fmt.Errorf("unknown Review field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *ReviewMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *ReviewMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *ReviewMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *ReviewMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *ReviewMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *ReviewMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *ReviewMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Review unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *ReviewMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Review edge %s", name)
}
