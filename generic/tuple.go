package generic

import (
	"fmt"
	"reflect"
)

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
