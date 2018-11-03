package graph

import (
	"unsafe"
)

var _inst_ShortestPath__10 = _inst_ShortestPath__p_p{
	_func_NewHeap: (func([]*item__p_p, func(*item__p_p, *item__p_p) bool, func(**item__p_p, int)) *Heap__p)(unsafe.Pointer(NewHeap__11)),
	_operator_Node_equal: (func(_generic_p, _generic_p) bool)(unsafe.Pointer(func(n1, n2 *TestNode) bool {
		return n1 == n2
	})),
	_method_Node__Edges: (func(_generic_p) []_generic_p)(unsafe.Pointer((*TestNode).Edges)),
	_method_Edge__Nodes: (func(_generic_p) (_generic_p, _generic_p))(unsafe.Pointer((*TestEdge).Nodes)),
	_interface_Node: (func(_generic_p) interface{})(unsafe.Pointer(func(n *TestNode) interface{} {
		return n
	})),
	_method_Heap__Pop:  (func(*Heap__p) *item__p_p)(unsafe.Pointer((*Heap__11).Pop)),
	_method_Heap__Fix:  (func(*Heap__p, int))(unsafe.Pointer((*Heap__11).Fix)),
	_method_Heap__Push: (func(*Heap__p, *item__p_p))(unsafe.Pointer((*Heap__11).Push)),
}

func ShortestPath__10(from, to *TestNode) []*TestEdge {
	return (func(_inst *_inst_ShortestPath__p_p, from, to *TestNode) []*TestEdge)(unsafe.Pointer(ShortestPath__p_p))(&_inst_ShortestPath__10, from, to)
}

type _inst_ShortestPath__p_p struct {
	_func_NewHeap        func([]*item__p_p, func(*item__p_p, *item__p_p) bool, func(**item__p_p, int)) *Heap__p
	_operator_Node_equal func(_generic_p, _generic_p) bool
	_method_Heap__Pop    func(*Heap__p) *item__p_p
	_method_Heap__Fix    func(*Heap__p, int)
	_method_Heap__Push   func(*Heap__p, *item__p_p)
	_method_Node__Edges  func(_generic_p) []_generic_p
	_method_Edge__Nodes  func(_generic_p) (_generic_p, _generic_p)
	_interface_Node      func(_generic_p) interface{}
}

func ShortestPath__p_p(_inst *_inst_ShortestPath__p_p, from, to _generic_p) []_generic_p {
	type item = item__p_p
	h := _inst._func_NewHeap([]*item__p_p{{
		n:     from,
		dist:  0,
		index: 0,
	}}, func(i1, i2 *item__p_p) bool {
		return i1.dist < i2.dist
	}, func(it **item__p_p, i int) {
		(*it).index = i
	})
	// Note: we'd like to use map[Node] *item but we
	// can't do that since we can't provide the equality and
	// hash functions to the internal map implementation.
	nodes := make(map[interface{}]*item)
	var found *item
	for len(h.Items) > 0 {
		nearest := _inst._method_Heap__Pop(h)
		if _inst._operator_Node_equal(nearest.n, to) {
			found = nearest
			break
		}
		for _, e := range _inst._method_Node__Edges(nearest.n) {
			edgeFrom, edgeTo := _inst._method_Edge__Nodes(e)
			if edgeFrom != nearest.n {
				continue
			}
			dist := nearest.dist + 1 // Could use e.Length() instead of 1.
			toItem, ok := nodes[_inst._interface_Node(edgeTo)]
			if !ok {
				it := &item{
					n:    edgeTo,
					dist: dist,
					edge: e,
				}
				nodes[_inst._interface_Node(edgeTo)] = it
				_inst._method_Heap__Push(h, it)
			} else if dist < toItem.dist {
				toItem.dist = dist
				toItem.edge = e
				_inst._method_Heap__Fix(h, toItem.index)
			}
		}
	}
	var edges []_generic_p
	for {
		edges = append(edges, found.edge)
		edgeFrom, _ := _inst._method_Edge__Nodes(found.edge)
		if edgeFrom == from {
			break
		}
		found = nodes[_inst._interface_Node(edgeFrom)]
	}
	return edges
}
