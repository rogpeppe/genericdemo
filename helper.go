package main

import (
	"fmt"
	"reflect"
)

type GenericFuncs struct {
	funcs      map[string]*genericFunc
	completers []func()
}

func (gf *GenericFuncs) Add(name string, types TypeTuple, inst interface{}) {
	if gf.funcs == nil {
		gf.funcs = make(map[string]*genericFunc)
	}
	f := gf.funcs[name]
	if f == nil {
		f = new(genericFunc)
		gf.funcs[name] = f
	}
	f.add(types, inst)
}

// AddCompleter registers a function to populate
// the generic tables after the first initialization
// has taken place. The function may assume
// that entry points for a generic function have been
// populated.
func (gf *GenericFuncs) AddCompleter(f func()) {
	gf.completers = append(gf.completers, f)
}

func (gf *GenericFuncs) Start() {
	for _, f := range gf.completers {
		f()
	}
}

// Get gets an instance of the function with the given name and
// type parameters.
func (gf *GenericFuncs) Get(name string, types TypeTuple) interface{} {
	f := gf.funcs[name]
	if f == nil {
		panic("no function found")
	}
	return f.get(types)
}

type typeIndexes []TypeTuple

func (ti typeIndexes) index(ts TypeTuple) int {
	i := ti.find(ts)
	if i == -1 {
		panic(fmt.Errorf("types %v not found", ts))
	}
	return i
}

func (ti typeIndexes) find(want TypeTuple) int {
	for i, ts := range ti {
		if ts.equal(want) {
			return i
		}
	}
	return -1
}

func (ti typeIndexes) has(ts TypeTuple) bool {
	return ti.find(ts) >= 0
}

// Types returns the type tuple of the types
// of all the values pointed to by each member of vs.
func Types(vs ...interface{}) TypeTuple {
	ts := make(TypeTuple, len(vs))
	for i, v := range vs {
		ts[i] = reflect.TypeOf(v).Elem()
	}
	return ts
}

// genericFunc represents a generic function with one or
// more instances. Each entry in Types represents an
// instance of the function for that set of type parameters,
// and the corresponding entry in Instances holds the
// function for that instance.
type genericFunc struct {
	types     typeIndexes
	instances []interface{}
}

// get gets the instance for the given type parameters.
func (f *genericFunc) get(ts TypeTuple) interface{} {
	return f.instances[f.types.index(ts)]
}

// set sets the instance for the given type parameters.
// A slot for the instance must have been allocated already
// by calling add.
func (f *genericFunc) set(ts TypeTuple, val interface{}) {
	f.instances[f.types.index(ts)] = val
}

// add adds an instance to the generic function with the given
// type parameters.
func (f *genericFunc) add(t TypeTuple, inst interface{}) {
	if f.types.has(t) {
		panic("types registered twice")
	}
	f.types = append(f.types, t)
	f.instances = append(f.instances, inst)
}

// TypeTuple represents a set of type parameters.
type TypeTuple []reflect.Type

func (t1 TypeTuple) equal(t2 TypeTuple) bool {
	if len(t1) != len(t1) {
		return false
	}
	for i, t := range t1 {
		if t2[i] != t {
			return false
		}
	}
	return true
}

func (ts TypeTuple) String() string {
	if len(ts) == 0 {
		return "()"
	}
	var s []byte
	s = append(s, '(')
	s = append(s, ts[0].String()...)
	for _, t := range ts[1:] {
		s = append(s, ", "...)
		s = append(s, t.String()...)
	}
	s = append(s, ')')
	return string(s)
}
