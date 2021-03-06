// Code generated by entc, DO NOT EDIT.

package predicate

import (
	"github.com/facebookincubator/ent/dialect/sql"
)

// Cart is the predicate function for cart builders.
type Cart func(*sql.Selector)

// Item is the predicate function for item builders.
type Item func(*sql.Selector)

// LineItem is the predicate function for lineitem builders.
type LineItem func(*sql.Selector)

// Order is the predicate function for order builders.
type Order func(*sql.Selector)

// OrderLineItem is the predicate function for orderlineitem builders.
type OrderLineItem func(*sql.Selector)
