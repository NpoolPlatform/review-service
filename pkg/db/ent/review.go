// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"
	"github.com/google/uuid"
)

// Review is the model entity for the Review schema.
type Review struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// EntityType holds the value of the "entity_type" field.
	EntityType string `json:"entity_type,omitempty"`
	// Domain holds the value of the "domain" field.
	Domain string `json:"domain,omitempty"`
	// ObjectID holds the value of the "object_id" field.
	ObjectID uuid.UUID `json:"object_id,omitempty"`
	// ReviewerID holds the value of the "reviewer_id" field.
	ReviewerID uuid.UUID `json:"reviewer_id,omitempty"`
	// State holds the value of the "state" field.
	State review.State `json:"state,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt uint32 `json:"create_at,omitempty"`
	// UpdateAt holds the value of the "update_at" field.
	UpdateAt uint32 `json:"update_at,omitempty"`
	// DeleteAt holds the value of the "delete_at" field.
	DeleteAt uint32 `json:"delete_at,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Review) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case review.FieldCreateAt, review.FieldUpdateAt, review.FieldDeleteAt:
			values[i] = new(sql.NullInt64)
		case review.FieldEntityType, review.FieldDomain, review.FieldState, review.FieldMessage:
			values[i] = new(sql.NullString)
		case review.FieldID, review.FieldObjectID, review.FieldReviewerID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Review", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Review fields.
func (r *Review) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case review.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				r.ID = *value
			}
		case review.FieldEntityType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field entity_type", values[i])
			} else if value.Valid {
				r.EntityType = value.String
			}
		case review.FieldDomain:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field domain", values[i])
			} else if value.Valid {
				r.Domain = value.String
			}
		case review.FieldObjectID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field object_id", values[i])
			} else if value != nil {
				r.ObjectID = *value
			}
		case review.FieldReviewerID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field reviewer_id", values[i])
			} else if value != nil {
				r.ReviewerID = *value
			}
		case review.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				r.State = review.State(value.String)
			}
		case review.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				r.Message = value.String
			}
		case review.FieldCreateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				r.CreateAt = uint32(value.Int64)
			}
		case review.FieldUpdateAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_at", values[i])
			} else if value.Valid {
				r.UpdateAt = uint32(value.Int64)
			}
		case review.FieldDeleteAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field delete_at", values[i])
			} else if value.Valid {
				r.DeleteAt = uint32(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Review.
// Note that you need to call Review.Unwrap() before calling this method if this Review
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Review) Update() *ReviewUpdateOne {
	return (&ReviewClient{config: r.config}).UpdateOne(r)
}

// Unwrap unwraps the Review entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Review) Unwrap() *Review {
	tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Review is not a transactional entity")
	}
	r.config.driver = tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Review) String() string {
	var builder strings.Builder
	builder.WriteString("Review(")
	builder.WriteString(fmt.Sprintf("id=%v", r.ID))
	builder.WriteString(", entity_type=")
	builder.WriteString(r.EntityType)
	builder.WriteString(", domain=")
	builder.WriteString(r.Domain)
	builder.WriteString(", object_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ObjectID))
	builder.WriteString(", reviewer_id=")
	builder.WriteString(fmt.Sprintf("%v", r.ReviewerID))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", r.State))
	builder.WriteString(", message=")
	builder.WriteString(r.Message)
	builder.WriteString(", create_at=")
	builder.WriteString(fmt.Sprintf("%v", r.CreateAt))
	builder.WriteString(", update_at=")
	builder.WriteString(fmt.Sprintf("%v", r.UpdateAt))
	builder.WriteString(", delete_at=")
	builder.WriteString(fmt.Sprintf("%v", r.DeleteAt))
	builder.WriteByte(')')
	return builder.String()
}

// Reviews is a parsable slice of Review.
type Reviews []*Review

func (r Reviews) config(cfg config) {
	for _i := range r {
		r[_i].config = cfg
	}
}
