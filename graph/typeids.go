package graph

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/rogpeppe/genericdemo/generic"

	"unsafe"
)

var typeIds = map[generic.TypeTuple]int{generic.Types(new(Int)): 2,
	generic.Types(new(int)):                       3,
	generic.Types(new(Heap__3)):                   4,
	generic.Types(new(*indexedInt)):               5,
	generic.Types(new(Heap__5)):                   6,
	generic.Types(new(*indexedInt)):               7,
	generic.Types(new(*TestNode)):                 8,
	generic.Types(new(*TestEdge)):                 9,
	generic.Types(new(*TestNode), new(*TestEdge)): 10,
	generic.Types(new(*item__10)):                 11,
	generic.Types(new(_p_NewHeap__t0)):            16,
	generic.Types(new(Heap__16)):                  17,
}

func expand(t reflect.Type) string {
	if t.Name() != "" {
		return expandTypeName(t)
	}
	switch t.Kind() {
	case reflect.Func:
		inParts := make([]string, t.NumIn())
		for i := 0; i < t.NumIn(); i++ {
			inParts[i] = expand(t.In(i))
		}
		outParts := make([]string, t.NumOut())
		for i := 0; i < t.NumOut(); i++ {
			outParts[i] = expand(t.Out(i))
		}
		// TODO variadic
		s := fmt.Sprintf("func(%s)", strings.Join(inParts, ", "))
		switch len(outParts) {
		case 0:
		case 1:
			s += " " + outParts[0]
		default:
			s += " (" + strings.Join(outParts, ", ") + ")"
		}
		return s
	case reflect.Ptr:
		return "*" + expand(t.Elem())
	case reflect.Slice:
		return "[]" + expand(t.Elem())
	}
	panic("not yet " + t.Kind().String())
}

func expandTypeName(t reflect.Type) string {
	parts := strings.Split(t.Name(), "__")
	if len(parts) > 2 {
		panic("bad type name " + t.Name())
	}
	if len(parts) == 1 {
		return t.String()
	}
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		panic("bad index in type name " + t.Name())
	}
	for tuple, index := range typeIds {
		if index == id {
			args := make([]string, tuple.Len())
			for i := 0; i < tuple.Len(); i++ {
				args[i] = expand(tuple.At(i))
			}
			return fmt.Sprintf("%s(%s)", parts[0], strings.Join(args, ", "))
		}
	}
	panic("cannot find type param for " + t.Name())
}

type _generic_v8a8 struct {
	t0 int64
}

type _generic_p struct {
	t0 unsafe.Pointer
}
