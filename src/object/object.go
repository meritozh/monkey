// Copyright (c) 2018 meritozh
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package object

import (
	"ast"
	"bytes"
	"fmt"
	"strings"
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
	// RETURNVALUEOBJ return value object
	RETURNVALUEOBJ = "RETURN_VALUE"
	// ERROROBJ error object, for error handling
	ERROROBJ = "ERROR"
	// FUNCTIONOBJ function object
	FUNCTIONOBJ = "FUNCTION"
	// STRINGOBJ string object
	STRINGOBJ = "STRING"
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

// ReturnValue return value object
type ReturnValue struct {
	Value Object
}

// Inspect implement Object interface
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Type implement Object interface
func (rv *ReturnValue) Type() Type { return RETURNVALUEOBJ }

// Error error object
type Error struct {
	Message string
}

// Inspect implement Object interface
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Type implement Object interface
func (e *Error) Type() Type { return ERROROBJ }

// Function function object
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Inspect implement Object interface
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n")

	return out.String()
}

// Type implement Object interface
func (f *Function) Type() Type { return FUNCTIONOBJ }

// String string object
type String struct {
	Value string
}

// Inspect implement Object interface
func (s *String) Inspect() string { return s.Value }

// Type implement Object interface
func (s *String) Type() Type { return STRINGOBJ }
