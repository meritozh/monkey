// Copyright (c) 2018 meritozh
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package object

import (
	"ast"
	"bytes"
	"fmt"
	"hash/fnv"
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
	// BUILTINOBJ bulti-in function
	BUILTINOBJ = "BUILTIN"
	// ARRAYOBJ array object
	ARRAYOBJ = "ARRAY"
	// HASHOBJ hash object
	HASHOBJ = "HASH"
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

// BuiltinFunction built-in function prototype
type BuiltinFunction func(args ...Object) Object

// Builtin built-in function
type Builtin struct {
	Fn BuiltinFunction
}

// Inspect implement Object interface
func (b *Builtin) Inspect() string { return "builtin function" }

// Type implement Object interface
func (b *Builtin) Type() Type { return BUILTINOBJ }

// Array array object
type Array struct {
	Elements []Object
}

// Inspect implement Object interface
func (a *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// Type implement Object interface
func (a *Array) Type() Type { return ARRAYOBJ }

// HashKey hash key
type HashKey struct {
	Type  Type
	Value uint64
}

// HashKey boolean object hashable
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey integer object hashable
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey string object hashable
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair hash object base structure
type HashPair struct {
	Key   Object
	Value Object
}

// Hash hash object
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Inspect implement Object interface
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Type implement Object interface
func (h *Hash) Type() Type { return HASHOBJ }

// Hashable hashable interfece
type Hashable interface {
	HashKey() HashKey
}
