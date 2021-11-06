// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/NpoolPlatform/review-service/pkg/db/ent/review"
	"github.com/NpoolPlatform/review-service/pkg/db/ent/schema"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	reviewFields := schema.Review{}.Fields()
	_ = reviewFields
	// reviewDescID is the schema descriptor for id field.
	reviewDescID := reviewFields[0].Descriptor()
	// review.DefaultID holds the default value on creation for the id field.
	review.DefaultID = reviewDescID.Default.(func() uuid.UUID)
}
