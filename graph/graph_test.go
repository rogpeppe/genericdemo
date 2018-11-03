package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {

	nodes := make([]*TestNode, 6)
	for i := range nodes {
		nodes[i] = &TestNode{
			name: fmt.Sprintf("node%d", i),
		}
	}
	arcs := [][2]int{
		{0, 1},
		{2, 0},
		{2, 4},
		{2, 3},
		{1, 5},
		{2, 5},
	}
	for _, a := range arcs {
		from, to := nodes[a[0]], nodes[a[1]]
		edge := &TestEdge{
			from: from,
			to:   to,
		}
		from.edges = append(from.edges, edge)
		to.edges = append(to.edges, edge)
	}
	path := ShortestPath__10(nodes[0], nodes[5])
	for _, e := range path {
		t.Logf("%v -> %v", e.from.name, e.to.name)
	}
}
