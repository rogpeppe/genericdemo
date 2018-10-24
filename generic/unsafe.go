package generic

import (
	"fmt"
	"reflect"
	"unsafe"
)

// AsType retypes the value in x to the type of t.
func AsType(x, t interface{}) interface{} {
	return valueAsType(reflect.ValueOf(x), reflect.TypeOf(t)).Interface()
}

// UnsafeSet sets the value pointed to by p to x,
// ignoring type considerations.
func UnsafeSet(p interface{}, x interface{}) {
	pv := reflect.ValueOf(p).Elem()
	pv.Set(valueAsType(reflect.ValueOf(x), pv.Type()))
}

// valueAsType retypes the value in x to t.
func valueAsType(x reflect.Value, t reflect.Type) reflect.Value {
	if x.Type().Size() != t.Size() {
		panic(fmt.Errorf("mismatched size %v vs %v", x.Type(), t))
	}
	// TODO also sanity check that the types are
	// pointer map, and alignment compatible.
	xcopy := reflect.New(x.Type())
	xcopy.Elem().Set(x)
	return reflect.NewAt(t, unsafe.Pointer(xcopy.Pointer())).Elem()
}
