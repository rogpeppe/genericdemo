package main

import (
	"fmt"
	"reflect"
)

type GenericFuncs struct {
	funcs map[string]map[TypeTuple]interface{}
}

func (gf *GenericFuncs) Add(name string, types TypeTuple, inst interface{}) {
	if gf.funcs == nil {
		gf.funcs = make(map[string]map[TypeTuple]interface{})
	}
	f := gf.funcs[name]
	if f == nil {
		f = make(map[TypeTuple]interface{})
		gf.funcs[name] = f
	}
	if _, ok := f[types]; ok {
		panic(fmt.Errorf("function %s%s registered twice", name, types))
	}
	f[types] = inst
}

// Get gets an instance of the function with the given name and
// type parameters.
func (gf *GenericFuncs) Get(name string, types TypeTuple) interface{} {
	f := gf.funcs[name]
	if f == nil {
		panic("no function found")
	}
	inst := f[types]
	if inst == nil {
		panic(fmt.Errorf("no implementation found for %s%s", name, types))
	}
	return inst
}

// TypeTuple represents a set of type parameters.
// It may be used as a map key.
type TypeTuple struct {
	t reflect.Type
}

// Types returns the type tuple of the types
// of all the values pointed to by each member of vs.
func Types(vs ...interface{}) TypeTuple {
	fields := make([]reflect.StructField, len(vs))
	for i, v := range vs {
		fields[i] = reflect.StructField{
			Name: fmt.Sprintf("T%d", i),
			Type: reflect.TypeOf(v).Elem(),
		}
	}
	return TypeTuple{
		t: reflect.StructOf(fields),
	}
}

// At returns the i'th element of the type tuple.
func (ts TypeTuple) At(i int) reflect.Type {
	return ts.t.Field(i).Type
}

func (ts TypeTuple) String() string {
	n := ts.t.NumField()
	var s []byte
	s = append(s, '(')
	for i := 0; i < n; i++ {
		if i > 0 {
			s = append(s, ", "...)
		}
		s = append(s, ts.t.Field(i).Type.String()...)
	}
	s = append(s, ')')
	return string(s)
}

func copyVal(v reflect.Value) reflect.Value {
	v1 := reflect.New(v.Type()).Elem()
	v1.Set(v)
	return v1
}
