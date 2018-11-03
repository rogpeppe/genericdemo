package graph

import (
	"unsafe"
)

var _inst_NewHeap__4 = _inst_NewHeap__p{
	_method_Heap__down: (func(*Heap__p, int, int) bool)(unsafe.Pointer((*Heap__4).down)),
}

func NewHeap__4(items []*int, less func(*int, *int) bool, setIndex func(**int, int)) *Heap__4 {
	return (func(_inst *_inst_NewHeap__p, items []*int, less func(*int, *int) bool, setIndex func(**int, int)) *Heap__4)(unsafe.Pointer(NewHeap__p))(&_inst_NewHeap__4, items, less, setIndex)
}

var _inst_NewHeap__5 = _inst_NewHeap__p{
	_method_Heap__down: (func(*Heap__p, int, int) bool)(unsafe.Pointer((*Heap__5).down)),
}

func NewHeap__5(items []*indexedInt, less func(*indexedInt, *indexedInt) bool, setIndex func(**indexedInt, int)) *Heap__5 {
	return (func(_inst *_inst_NewHeap__p, items []*indexedInt, less func(*indexedInt, *indexedInt) bool, setIndex func(**indexedInt, int)) *Heap__5)(unsafe.Pointer(NewHeap__p))(&_inst_NewHeap__5, items, less, setIndex)
}

var _inst_NewHeap__11 = _inst_NewHeap__p{
	_method_Heap__down: (func(*Heap__p, int, int) bool)(unsafe.Pointer((*Heap__11).down)),
}

func NewHeap__11(items []*item__10, less func(*item__10, *item__10) bool, setIndex func(**item__10, int)) *Heap__11 {
	return (func(_inst *_inst_NewHeap__p, items []*item__10, less func(*item__10, *item__10) bool, setIndex func(**item__10, int)) *Heap__11)(unsafe.Pointer(NewHeap__p))(&_inst_NewHeap__11, items, less, setIndex)
}

type item__p_p struct {
	n     _generic_p
	dist  int
	index int
	edge  _generic_p
}

type item__10 struct {
	n     *TestNode
	dist  int
	index int
	edge  *TestEdge
}

type _inst_NewHeap__p struct {
	_method_Heap__down func(*Heap__p, int, int) bool
}

func NewHeap__p(_inst *_inst_NewHeap__p, items []_generic_p, less func(_generic_p, _generic_p) bool, setIndex func(e *_generic_p, i int)) *Heap__p {
	h := &Heap__p{
		Items:    items,
		less:     less,
		setIndex: setIndex,
	}
	n := len(h.Items)
	if setIndex != nil {
		for i := 0; i < n; i++ {
			setIndex(&items[i], i)
		}
	}
	for i := n/2 - 1; i >= 0; i-- {
		_inst._method_Heap__down(h, i, n)
	}
	return h
}

var _inst_Heap__4 = _inst_Heap__p{}

type Heap__4 struct {
	Items    []*int
	less     func(*int, *int) bool
	setIndex func(*int, int)
}

func (h *Heap__4) Push(x *int) {
	(func(*_inst_Heap__p, *Heap__4, *int))(unsafe.Pointer(Heap__p_Push))(&_inst_Heap__4, h, x)
}

func (h *Heap__4) Pop() *int {
	return (func(*_inst_Heap__p, *Heap__4) *int)(unsafe.Pointer(Heap__p_Pop))(&_inst_Heap__4, h)
}

func (h *Heap__4) Fix(i int) {
	(func(*_inst_Heap__p, *Heap__4, int))(unsafe.Pointer(Heap__p_Fix))(&_inst_Heap__4, h, i)
}

func (h *Heap__4) Remove(i int) int {
	return (func(*_inst_Heap__p, *Heap__4, int) int)(unsafe.Pointer(Heap__p_Remove))(&_inst_Heap__4, h, i)
}

func (h *Heap__4) down(i0, n int) bool {
	return (func(_inst *_inst_Heap__p, h *Heap__4, i0, n int) bool)(unsafe.Pointer(Heap__p_down))(&_inst_Heap__4, h, i0, n)
}

var _inst_Heap__5 = _inst_Heap__p{}

type Heap__5 struct {
	Items    []*indexedInt
	less     func(*indexedInt, *indexedInt) bool
	setIndex func(*indexedInt, int)
}

func (h *Heap__5) Push(x *indexedInt) {
	(func(*_inst_Heap__p, *Heap__5, *indexedInt))(unsafe.Pointer(Heap__p_Push))(&_inst_Heap__5, h, x)
}

func (h *Heap__5) Pop() *indexedInt {
	return (func(*_inst_Heap__p, *Heap__5) *indexedInt)(unsafe.Pointer(Heap__p_Pop))(&_inst_Heap__5, h)
}

func (h *Heap__5) Fix(i int) {
	(func(*_inst_Heap__p, *Heap__5, int))(unsafe.Pointer(Heap__p_Fix))(&_inst_Heap__5, h, i)
}

func (h *Heap__5) Remove(i int) int {
	return (func(*_inst_Heap__p, *Heap__5, int) int)(unsafe.Pointer(Heap__p_Remove))(&_inst_Heap__5, h, i)
}

func (h *Heap__5) down(i0, n int) bool {
	return (func(_inst *_inst_Heap__p, h *Heap__5, i0, n int) bool)(unsafe.Pointer(Heap__p_down))(&_inst_Heap__5, h, i0, n)
}

var _inst_Heap__11 = _inst_Heap__p{}

type Heap__11 struct {
	Items    []*item__10
	less     func(*item__10, *item__10) bool
	setIndex func(*item__10, int)
}

func (h *Heap__11) Push(x *item__10) {
	(func(*_inst_Heap__p, *Heap__11, *item__10))(unsafe.Pointer(Heap__p_Push))(&_inst_Heap__11, h, x)
}

func (h *Heap__11) Pop() *item__10 {
	return (func(*_inst_Heap__p, *Heap__11) *item__10)(unsafe.Pointer(Heap__p_Pop))(&_inst_Heap__11, h)
}

func (h *Heap__11) Fix(i int) {
	(func(*_inst_Heap__p, *Heap__11, int))(unsafe.Pointer(Heap__p_Fix))(&_inst_Heap__11, h, i)
}

func (h *Heap__11) Remove(i int) int {
	return (func(*_inst_Heap__p, *Heap__11, int) int)(unsafe.Pointer(Heap__p_Remove))(&_inst_Heap__11, h, i)
}

func (h *Heap__11) down(i0, n int) bool {
	return (func(_inst *_inst_Heap__p, h *Heap__11, i0, n int) bool)(unsafe.Pointer(Heap__p_down))(&_inst_Heap__11, h, i0, n)
}

type _inst_Heap__p struct{}

type Heap__p struct {
	Items    []_generic_p
	less     func(_generic_p, _generic_p) bool
	setIndex func(*_generic_p, int)
}

func Heap__p_Push(_inst *_inst_Heap__p, h *Heap__p, x _generic_p) {
	h.Items = append(h.Items, x)
	Heap__p_up(_inst, h, len(h.Items)-1)
}

func Heap__p_Pop(_inst *_inst_Heap__p, h *Heap__p) _generic_p {
	n := len(h.Items) - 1
	Heap__p_swap(_inst, h, 0, n)
	Heap__p_down(_inst, h, 0, n)
	return Heap__p_pop(_inst, h)
}

func Heap__p_Fix(_inst *_inst_Heap__p, h *Heap__p, i int) {
	if !Heap__p_down(_inst, h, i, len(h.Items)) {
		Heap__p_up(_inst, h, i)
	}
}

func Heap__p_swap(_inst *_inst_Heap__p, h *Heap__p, i, j int) {
	h.Items[i], h.Items[j] = h.Items[j], h.Items[i]
	if h.setIndex != nil {
		h.setIndex(&h.Items[i], i)
		h.setIndex(&h.Items[j], j)
	}
}

func Heap__p_Remove(_inst *_inst_Heap__p, h *Heap__p, i int) _generic_p {
	n := len(h.Items) - 1
	if n != i {
		Heap__p_swap(_inst, h, i, n)
		if !Heap__p_down(_inst, h, i, n) {
			Heap__p_up(_inst, h, i)
		}
	}
	return Heap__p_pop(_inst, h)
}

func Heap__p_pop(_inst *_inst_Heap__p, h *Heap__p) _generic_p {
	n := len(h.Items) - 1
	x := h.Items[n]
	h.Items = h.Items[0:n]
	return x
}

func Heap__p_up(_inst *_inst_Heap__p, h *Heap__p, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.less(h.Items[j], h.Items[i]) {
			break
		}
		Heap__p_swap(_inst, h, i, j)
		j = i
	}
}

func Heap__p_down(_inst *_inst_Heap__p, h *Heap__p, i0, n int) bool {
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
		Heap__p_swap(_inst, h, i, j)
		i = j
	}
	return i > i0
}
