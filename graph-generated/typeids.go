package graph

import (
	"github.com/rogpeppe/genericdemo/generic"

	"unsafe"
)

// typeIds isn't used, but acts as documentation for what types
// are assigned which integer ids (so that the numbers can be
// used in type names).
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
}

// _generic_v8a8 represents a generic 8 byte, 8 byte aligned
// data value.
type _generic_v8a8 struct {
	t0 int64
}

// _generic_v8a8 represents a generic pointer data value.
type _generic_p struct {
	t0 unsafe.Pointer
}
