// +build ignore

package graph

contract Graph(n Node, e Edge) {
	var _ []Edge = n.Edges()
	var from, to Node = e.Nodes()
	map[Node]bool{}
}

// item holds an item in the node fringe being calculated by
// ShortestPath. We might normally declare this inside ShortestPath
// itself, but that would mean we couldn't refer to it globally in the
// generated code (not a problem when doing generics directly in the
// compiler)
type item(type Node, Edge) struct {
	n Node
	dist int
	index int
	edge Edge
}

func ShortestPath(type Node, Edge Graph)(from, to Node) []Edge {
	type item = item(Node, Edge)
	h := NewHeap([]*item(Node, Edge){{
		n: from,
		dist: 0,
		index: 0,
	}}, func(i1, i2 *item) bool {
		return i1.dist < i2.dist
	}, func(it **item, i int) {
		(*it).index = i
	})
	// Note: we'd like to use map[Node] *item but we
	// can't do that since we can't provide the equality and
	// hash functions to the internal map implementation.
	nodes := make(map[interface{}] *item)
	var found *item
	for len(h.Items) > 0 {
		nearest := h.Pop()
		if nearest.n == to {
			found = nearest
			break
		}
		for _, e := range nearest.n.Edges() {
			edgeFrom, edgeTo := e.Nodes()
			if from != nearest.n {
				continue
			}
			dist := nearest.dist + 1		// Could use e.Length() instead of 1.
			toItem, ok := nodes[edgeTo]
			if !ok {
				nodes[edgeTo] = &item{
					n: edgeTo,
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
	var edges []Edge
	for {
		edges = append(edges, found.edge)
		edgeFrom, edgeTo := found.edge.Nodes()
		if edgeFrom == from {
			break
		}
		found = nodes[edgeFrom]
	}
	return edges
}

