package graph

import (
	"github.com/rogpeppe/genericdemo/generic"

	"unsafe"
)

var typeIds = map[generic.TypeTuple]int{
	//generic.Types(new(_param_NewHeap__E)): 1,
	generic.Types(new(Int)):     2,
	generic.Types(new(int)):     3,
	generic.Types(new(Heap__3)): 4,
}

type _generic_v8a8 struct {
	_ int64
}

type _generic_p unsafe.Pointer
