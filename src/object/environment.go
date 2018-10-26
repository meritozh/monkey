// Copyright (c) 2018 meritozh
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package object

// Environment symbol table to track identifier and value binding
type Environment struct {
	store map[string]Object
	outer *Environment
}

// NewEnclosedEnvironment create new environment used in enclosed block
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// NewEnvironment create new Environment instance
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// Get get symbol bound object and status
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set bind object to symbol
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
