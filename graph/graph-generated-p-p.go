package graph

import (
	"github.com/rogpeppe/genericdemo/generic"
)

// To change code back to unsafe type conversions:
// gofmt -r 'generic.AsType(y, (x)(nil)).(x) -> (x)(unsafe.Pointer(y))'

var _inst_ShortestPath__10 = _inst_ShortestPath__p_p{
	_func_NewHeap: generic.AsType(NewHeap__11, (func([]*item__p_p, func(*item__p_p, *item__p_p) bool, func(**item__p_p, int)) *Heap__p)(nil)).(func([]*item__p_p, func(*item__p_p, *item__p_p) bool, func(**item__p_p, int)) *Heap__p),
	_operator_Node_equal: generic.AsType(func(n1, n2 *TestNode) bool {
		return n1 == n2
	}, (func(_generic_p, _generic_p) bool)(nil)).(func(_generic_p, _generic_p) bool),

	_method_Node__Edges: generic.AsType((*TestNode).Edges, (func(_generic_p) []_generic_p)(nil)).(func(_generic_p) []_generic_p),
	_method_Edge__Nodes: generic.AsType((*TestEdge).Nodes, (func(_generic_p) (_generic_p, _generic_p))(nil)).(func(_generic_p) (_generic_p, _generic_p)),
	_interface_Node: generic.AsType(func(n *TestNode) interface{} {
		return n
	}, (func(_generic_p) interface{})(nil)).(func(_generic_p) interface{}),

	_method_Heap__Pop:  generic.AsType((*Heap__11).Pop, (func(*Heap__p) *item__p_p)(nil)).(func(*Heap__p) *item__p_p),
	_method_Heap__Fix:  generic.AsType((*Heap__11).Fix, (func(*Heap__p, int))(nil)).(func(*Heap__p, int)),
	_method_Heap__Push: generic.AsType((*Heap__11).Push, (func(*Heap__p, *item__p_p))(nil)).(func(*Heap__p, *item__p_p)),
}

func ShortestPath__10(from, to *TestNode) []*TestEdge {
	return generic.AsType(ShortestPath__p_p, (func(_inst *_inst_ShortestPath__p_p, from, to *TestNode) []*TestEdge)(nil)).(func(_inst *_inst_ShortestPath__p_p, from, to *TestNode) []*TestEdge)(&_inst_ShortestPath__10, from, to)
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
