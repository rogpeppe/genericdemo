package generic

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func isGenericType(t reflect.Type) bool {
	return strings.HasPrefix(t.Name(), "_generic_")
}

/*
TODO
is rt-generic t2 equivalent to instantated t1 with respect to generic type gt?

if gt is a concrete type, t1 must be identical to t2
if gt is a type parameter, t2 must be equivalent type to t2.
if gt is a parameterized type, look up generic type
and substitute its type parameters for the actual type arguments.
*/

// equivalent reports whether the instantiated type t1 is
// runtime-representation-equivalent to t2. t1 must not contain any
// _generic types.
func equivalent(t1, t2 reflect.Type) (bool, error) {
	// TODO check for cycles.

	if isGenericType(t1) {
		return false, fmt.Errorf("instantiated type holding generic type")
	}
	if isGenericType(t2) {
		return typeMapOf(t1).equal(typeMapOf(t2)), nil
	}
	if t1.Kind() != t2.Kind() {
		return false, nil
	}
	if t1.NumMethod() != t2.NumMethod() {
		return false, nil
	}
	for i := 0; i < t1.NumMethod(); i++ {
		m1, m2 := t1.Method(i), t2.Method(i)
		if m1.Name != m2.Name || m1.PkgPath != m2.PkgPath {
			return false, nil
		}
		// Note: except for interface types, which don't include the
		// receiver in their Type, we want to skip the receiver
		// when comparing the methods because otherwise
		// we'll get into an infinite loop.
		if ok, err := equivalentFunc(m1.Type, m2.Type, t1.Kind() != reflect.Interface); !ok || err != nil {
			return false, err
		}
	}
	switch t1.Kind() {
	case reflect.Ptr,
		reflect.Slice,
		reflect.Chan:
		return equivalent(t1.Elem(), t2.Elem())
	case reflect.Map:
		ok1, err1 := equivalent(t1.Key(), t2.Key())
		ok2, err2 := equivalent(t1.Key(), t2.Key())
		if err1 != nil {
			return false, err1
		}
		if err2 != nil {
			return false, err2
		}
		return ok1 && ok2, nil
	case reflect.Array:
		if t1.Len() != t2.Len() {
			return false, nil
		}
		return equivalent(t1.Elem(), t2.Elem())
	case reflect.Struct:
		if t1.NumField() != t2.NumField() {
			return false, nil
		}
		for i := 0; i < t1.NumField(); i++ {
			f1, f2 := t1.Field(i), t2.Field(i)
			ft1, ft2 := f1.Type, f2.Type
			f1.Type, f2.Type = nil, nil
			if !reflect.DeepEqual(f1, f2) {
				return false, nil
			}
			if ok, err := equivalent(ft1, ft2); !ok || err != nil {
				return false, err
			}
		}
		return true, nil
	case reflect.Func:
		return equivalentFunc(t1, t2, false)
	case reflect.Interface:
		// We've already checked the interface methods above.
		return true, nil
	default:
		return t1 == t2, nil
	}
}

func equivalentFunc(t1, t2 reflect.Type, skip1 bool) (bool, error) {
	if t1.NumIn() != t2.NumIn() ||
		t1.NumOut() != t2.NumOut() ||
		t1.IsVariadic() != t2.IsVariadic() {
		return false, nil
	}
	skip := 0
	if skip1 {
		skip = 1
	}
	for i := skip; i < t1.NumIn(); i++ {
		if ok, err := equivalent(t1.In(i), t2.In(i)); !ok || err != nil {
			return false, err
		}
	}
	for i := 0; i < t1.NumOut(); i++ {
		if ok, err := equivalent(t1.In(i), t2.In(i)); !ok || err != nil {
			return false, err
		}
	}
	return true, nil
}

// equivType returns a type that's equivalent to t
// for storage - that is, it has the same size,
// alignment and pointer map. It also returns
// a name to use for the type.
func equivType(t reflect.Type) (reflect.Type, string) {
	panic("still to do")
}

type typeMap struct {
	align      int
	fieldAlign int
	size       uintptr
	// ptrs holds the word map of where pointers
	// are located within the object. The last element
	// of ptrs is always true - any remaining bytes are
	// implied by size.
	ptrs []bool
}

func (m *typeMap) equal(m1 *typeMap) bool {
	return reflect.DeepEqual(m, m1)
}

func typeMapOf(t reflect.Type) *typeMap {
	m := &typeMap{
		align:      t.Align(),
		fieldAlign: t.FieldAlign(),
		size:       t.Size(),
	}
	m.addPtrs(0, t)
	return m
}

const ptrSize = unsafe.Sizeof(uintptr(0))

func (m *typeMap) addPtr(offset uintptr) {
	if offset%ptrSize != 0 {
		panic("pointer with non-pointer alignment")
	}
	b := int(offset / ptrSize)
	if b < len(m.ptrs) {
		panic("pointer added twice")
	}
	for i := len(m.ptrs); i < b; i++ {
		m.ptrs = append(m.ptrs, false)
	}
	m.ptrs = append(m.ptrs, true)
}

func (m *typeMap) typeName() string {
	if len(m.ptrs) == 0 {
		if m.fieldAlign != m.align {
			panic("when does this happen?")
		}
		return fmt.Sprintf("v%da%d", m.size, m.align)
	}
	if m.align != int(ptrSize) || m.fieldAlign != int(ptrSize) {
		panic("unexpected alignment for type with pointers")
	}
	var buf strings.Builder
	runStart := 0
	for i, b := range m.ptrs {
		if !b {
			continue
		}
		run := i - runStart
		if run > 0 {
			fmt.Fprintf(&buf, "v%d", uintptr(run)*ptrSize)
		}
		buf.WriteByte('p')
		runStart = i + 1
	}
	if remain := m.size - uintptr(runStart)*ptrSize; remain > 0 {
		fmt.Fprintf(&buf, "v%d", remain)
	}
	return buf.String()
}

func (m *typeMap) addPtrs(offset uintptr, t reflect.Type) {
	switch t.Kind() {
	case reflect.Ptr,
		reflect.Map,
		reflect.Func,
		reflect.Chan,
		reflect.UnsafePointer:
		m.addPtr(offset)
	case reflect.String:
		m.addPtr(offset)
	case reflect.Slice:
		m.addPtr(offset)
	case reflect.Interface:
		m.addPtr(offset)
		m.addPtr(offset + ptrSize)
	case reflect.Array:
		elem := t.Elem()
		elemSize := t.Size() / uintptr(t.Len())
		for i := uintptr(0); i < uintptr(t.Len()); i++ {
			m.addPtrs(offset+i*elemSize, elem)
		}
	case reflect.Struct:
		n := t.NumField()
		for i := 0; i < n; i++ {
			f := t.Field(i)
			m.addPtrs(offset+f.Offset, f.Type)
		}
	}
}
