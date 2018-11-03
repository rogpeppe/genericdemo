package graph

import "github.com/rogpeppe/genericdemo/generic"

var _inst_NewHeap__3 = _inst_NewHeap__v8a8{
	_method_Heap__down: generic.AsType((*Heap__3).down, (func(*Heap__v8a8, int, int) bool)(nil)).(func(*Heap__v8a8, int, int) bool),
}

func NewHeap__3(items []int, less func(int, int) bool, setIndex func(*int, int)) *Heap__3 {
	return generic.AsType(NewHeap__v8a8, (func(_inst *_inst_NewHeap__v8a8, items []int, less func(int, int) bool, setIndex func(*int, int)) *Heap__3)(nil)).(func(_inst *_inst_NewHeap__v8a8, items []int, less func(int, int) bool, setIndex func(*int, int)) *Heap__3)(&_inst_NewHeap__3, items, less, setIndex)
}

type _inst_NewHeap__v8a8 struct {
	_method_Heap__down func(*Heap__v8a8, int, int) bool
}

func NewHeap__v8a8(_inst *_inst_NewHeap__v8a8, items []_generic_v8a8, less func(_generic_v8a8, _generic_v8a8) bool, setIndex func(e *_generic_v8a8, i int)) *Heap__v8a8 {
	h := &Heap__v8a8{
		Items:    items,
		less:     less,
		setIndex: setIndex,
	}
	n := len(h.Items)
	for i := n/2 - 1; i >= 0; i-- {
		_inst._method_Heap__down(h, i, n)
	}
	return h
}

var _inst_Heap__3 = _inst_Heap__v8a8{}

type Heap__3 struct {
	Items    []int
	less     func(int, int) bool
	setIndex func(*int, int)
}

func (h *Heap__3) Push(x int) {
	generic.AsType(Heap__v8a8_Push, (func(*_inst_Heap__v8a8, *Heap__3, int))(nil)).(func(*_inst_Heap__v8a8, *Heap__3, int))(&_inst_Heap__3, h, x)
}

func (h *Heap__3) Pop() int {
	return generic.AsType(Heap__v8a8_Pop, (func(*_inst_Heap__v8a8, *Heap__3) int)(nil)).(func(*_inst_Heap__v8a8, *Heap__3) int)(&_inst_Heap__3, h)
}

func (h *Heap__3) Fix(i int) {
	generic.AsType(Heap__v8a8_Fix, (func(*_inst_Heap__v8a8, *Heap__3, int))(nil)).(func(*_inst_Heap__v8a8, *Heap__3, int))(&_inst_Heap__3, h, i)
}

func (h *Heap__3) Remove(i int) int {
	return generic.AsType(Heap__v8a8_Remove, (func(*_inst_Heap__v8a8, *Heap__3, int) int)(nil)).(func(*_inst_Heap__v8a8, *Heap__3, int) int)(&_inst_Heap__3, h, i)
}

func (h *Heap__3) down(i0, n int) bool {
	return generic.AsType(Heap__v8a8_down, (func(_inst *_inst_Heap__v8a8, h *Heap__3, i0, n int) bool)(nil)).(func(_inst *_inst_Heap__v8a8, h *Heap__3, i0, n int) bool)(&_inst_Heap__3, h, i0, n)
}

type _inst_Heap__v8a8 struct{}

type Heap__v8a8 struct {
	Items    []_generic_v8a8
	less     func(_generic_v8a8, _generic_v8a8) bool
	setIndex func(*_generic_v8a8, int)
}

func Heap__v8a8_Push(_inst *_inst_Heap__v8a8, h *Heap__v8a8, x _generic_v8a8) {
	h.Items = append(h.Items, x)
	Heap__v8a8_up(_inst, h, len(h.Items)-1)
}

func Heap__v8a8_Pop(_inst *_inst_Heap__v8a8, h *Heap__v8a8) _generic_v8a8 {
	n := len(h.Items) - 1
	Heap__v8a8_swap(_inst, h, 0, n)
	Heap__v8a8_down(_inst, h, 0, n)
	return Heap__v8a8_pop(_inst, h)
}

func Heap__v8a8_Fix(_inst *_inst_Heap__v8a8, h *Heap__v8a8, i int) {
	if !Heap__v8a8_down(_inst, h, i, len(h.Items)) {
		Heap__v8a8_up(_inst, h, i)
	}
}

func Heap__v8a8_swap(_inst *_inst_Heap__v8a8, h *Heap__v8a8, i, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
	if h.setIndex != nil {
		h.setIndex(&h.Items[i], i)
		h.setIndex(&h.Items[j], j)
	}
}

func Heap__v8a8_Remove(_inst *_inst_Heap__v8a8, h *Heap__v8a8, i int) _generic_v8a8 {
	n := len(h.Items) - 1
	if n != i {
		Heap__v8a8_swap(_inst, h, i, n)
		if !Heap__v8a8_down(_inst, h, i, n) {
			Heap__v8a8_up(_inst, h, i)
		}
	}
	return Heap__v8a8_pop(_inst, h)
}

func Heap__v8a8_pop(_inst *_inst_Heap__v8a8, h *Heap__v8a8) _generic_v8a8 {
	n := len(h.Items) - 1
	x := h.Items[n]
	h.Items = h.Items[0:n]
	return x
}

func Heap__v8a8_up(_inst *_inst_Heap__v8a8, h *Heap__v8a8, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(h.Items[j], h.Items[i]) {
			break
		}
		Heap__v8a8_swap(_inst, h, i, j)
		j = i
	}
}

func Heap__v8a8_down(_inst *_inst_Heap__v8a8, h *Heap__v8a8, i0, n int) bool {
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
		Heap__v8a8_swap(_inst, h, i, j)
		i = j
	}
	return i > i0
}
