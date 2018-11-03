package graph

type _p_NewHeap__t0 interface{}

func NewHeap__g(items []_p_NewHeap__t0, less func(_p_NewHeap__t0, _p_NewHeap__t0) bool, setIndex func(e *_p_NewHeap__t0, i int)) (_ *Heap__16) {
	return
}

type Heap__16 interface{} // placeholder

type _p_Heap__t0 interface{}

type Heap__g struct {
	Items    []_p_Heap__t0
	less     func(_p_Heap__t0, _p_Heap__t0) bool
	setIndex func(*_p_Heap__t0, int)
}

func (h *Heap__g) Push(x _p_Heap__t0) {
}

func (h *Heap__g) Pop() (_ _p_NewHeap__t0) {
	return
}

func (h *Heap__g) Fix(i int) {
}

func (h *Heap__g) Remove(i int) (_ _p_NewHeap__t0) {
	return
}
