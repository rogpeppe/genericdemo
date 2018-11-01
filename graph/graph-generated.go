// +build ignore

package graph

import "reflect"

type _inst_ShortestPath__p_p struct {
	_func_NewHeap       func([]_generic_p, func(_generic_p, _generic_p) bool, func(*_generic_p, int)) *Heap__p
	_method_Edge__Edges func(_generic_p) []_generic_p
	_method_Node__Nodes func(_generic_p) (_generic_p, _generic_p)
	_type_Node          reflect.Type
}

// _generic_p represents any pointer type.
type _generic_p struct {
	_ int64
}

type item__p_p struct {
	n     _generic_p
	dist  int
	index int
	edge  _generic_p
}

func ShortestPath__p_p(_inst _inst_ShortestPath__p_p, from, to _generic_p) []_generic_p {
	type item = item__p_p
	h := _inst._func_NewHeap([]*item_p_p{
		n:     from,
		dist:  0,
		index: 0,
	}, func(i1, i2 *item) bool {
		return i1.dist < i2.dist
	}, func(it **item, i int) {
		(*it).index = i
	})
	nodes := make(map[Node]*item)
	var found *item
	for len(h.Items) > 0 {
		nearest := h.Pop()
		if nearest.n == to {
			found = nearest
			break
		}
		for _, e := range _inst._method_Edge__Edges(nearest.n) {
			edgeFrom, edgeTo := _inst._method_Node__Nodes(e)
			if from != nearest.n {
				continue
			}
			dist := nearest.dist + 1 // Could use e.Length() instead of 1.
			toItem, ok := nodes[edgeTo]
			if !ok {
				nodes[edgeTo] = &item{
					n:    edgeTo,
					dist: dist,
					edge: e,
				}
			} else if dist < toItem.dist {
				toItem.dist = dist
				toItem.edge = e
				h.Fix(toItem.index)
			}
		}
	}
	var edges []_generic_p
	for {
		edges = append(edges, found.edge)
		edgeFrom, edgeTo := _int._method_Edge__Nodes(found.edge)
		if edgeFrom == from {
			break
		}
		found = nodes[edgeFrom]
	}
	return edges
}
