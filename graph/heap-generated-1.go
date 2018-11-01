// +build ignore

package graph

import "github.com/rogpeppe/genericdemo/generic"

var typeIds = map[generic.TypeTuple]int{
	generic.Types(new(_param_NewHeap__E)): 1,
}

type _param_NewHeap__E generic.Type

func NewHeap(items []_param_NewHeap__E, less func(_param_NewHeap__E, _param_NewHeap__E) bool) *Heap__1 {
	h := &Heap__1{
		Items: items,
		less:  less,
	}
	n := len(h.Items)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	return h
}

type _param_Heap_E generic.Type

type Heap struct {
	Items []_param_Heap_E
	less  func(_param_Heap_E, _param_Heap_E) bool
}

type _param_Heap__Push__E = _param_Heap_E

func (h *Heap) Push(x _param_Heap__Push__E) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.Items = append(h.Items, x)
	h.up(len(h.Items) - 1)
}

type _param_Heap__Pop__E = _param_Heap_E

func (h *Heap) Pop() _param_Heap__Pop__E {
	n := len(h.Items) - 1
	h.swap(0, n)
	h.down(0, n)
	return h.pop()
}

func (h *Heap) swap(i, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
}

type _param_Heap__Remove__E = _param_Heap_E

func (h *Heap) Remove(i int) _param_Heap__Remove__E {
	n := len(h.Items) - 1
	if n != i {
		h.swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	return h.pop()
}

type _param_Heap__pop__E = _param_Heap_E

func (h *Heap) pop() _param_Heap__pop__E {
	n := len(h.Items) - 1
	x := h.Items[n]
	h.Items = h.Items[0:n]
	return x
}

func (h *Heap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(h.Items[j], h.Items[i]) {
			break
		}
		h.swap(i, j)
		j = i
	}
}

func (h *Heap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.less(h.Items[j2], h.Items[j1]) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.less(h.Items[j], h.Items[i]) {
			break
		}
		h.swap(i, j)
		i = j
	}
	return i > i0
}
