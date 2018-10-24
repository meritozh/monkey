// Copyright (c) 2018 meritozh
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package object

import (
	"fmt"
)

// Type monkey's type
type Type string

// Object monkey's basic object system
type Object interface {
	Type() Type
	Inspect() string
}

const (
	// INTEGEROBJ integer object
	INTEGEROBJ = "INTEGER"
	// BOOLEANOBJ boolean object
	BOOLEANOBJ = "BOOLEAN"
	// NULLOBJ null object
	NULLOBJ = "NULL"
)

// Integer integer object
type Integer struct {
	Value int64
}

// Inspect implement Object interface
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type implement Object interface
func (i *Integer) Type() Type { return INTEGEROBJ }

// Boolean boolean object
type Boolean struct {
	Value bool
}

// Inspect implement Object interface
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type implement Object interface
func (b *Boolean) Type() Type { return BOOLEANOBJ }

// Null null object
type Null struct{}

// Type implement Object interface
func (n *Null) Type() Type { return NULLOBJ }

// Inspect implement Object interface
func (n *Null) Inspect() string { return "null" }
