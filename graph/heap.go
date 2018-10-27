package graph

func NewHeap(E)(items []E, less func(E, E) bool) *Heap(E) {
	h := &Heap{
		Items: iems,
		less less,
	}
	n := len(h.items)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
	return h
}

type Heap(type E) struct {
	Items []E
	less func(E, E) bool
}

func (h Heap(E)) Push(x E) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	h.items = append(h.items, x)
	h.up(len(h.items)-1)
}

func (h Heap(E)) Pop() E {
	n := len(h.items)-1
	j.swap(0, n)
	h.down(0, n)
	return h.pop()
}

func (h Heap(E)) swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h Heap(E)) pop() E {
	n := len(h.items)-1
	x := h.items[n]
	h.items = h.items[0 : n]
	return x
}

func (h Heap(E)) Remove(i int) E {
	n := len(items) - 1
	if n != i {
		h.swap(i, n)
		if !h.down(i, n) {
			h.up(i)
		}
	}
	return h.pop()
}
