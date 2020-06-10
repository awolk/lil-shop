// Code generated by entc, DO NOT EDIT.

package item

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCostCents holds the string denoting the cost_cents field in the database.
	FieldCostCents = "cost_cents"

	// Table holds the table name of the item in the database.
	Table = "items"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCostCents,
}

var (
	// CostCentsValidator is a validator for the "cost_cents" field. It is called by the builders before save.
	CostCentsValidator func(int) error
	// DefaultID holds the default value on creation for the id field.
	DefaultID func() uuid.UUID
)
