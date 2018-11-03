package graph

import (
	"fmt"
	"testing"

	qt "github.com/frankban/quicktest"
)

type graphTest struct {
	arcs     [][2]int
	from, to int
	want     [][2]int
}

var graphTests = []graphTest{{
	arcs: [][2]int{
		{0, 1},
		{2, 0},
		{2, 4},
		{2, 5},
		{2, 3},
		{1, 5},
		{2, 5},
	},
	from: 0,
	to:   5,
	want: [][2]int{
		{1, 5},
		{0, 1},
	},
}, {
	arcs: [][2]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{2, 3},
		{3, 4},
		{4, 2},
		{4, 5},
		{7, 0},
	},
	from: 7,
	to:   5,
	want: [][2]int{
		{4, 5},
		{3, 4},
		{0, 3},
		{7, 0},
	},
}}

func TestShortestPath(t *testing.T) {
	c := qt.New(t)

	for i, test := range graphTests {
		c.Run(fmt.Sprint("test", i), func(c *qt.C) {
			testShortestPath(c, test)
		})
	}
}

func testShortestPath(c *qt.C, test graphTest) {
	nodes := make(map[int]*TestNode)
	for _, a := range test.arcs {
		nodes[a[0]] = &TestNode{
			index: a[0],
		}
		nodes[a[1]] = &TestNode{
			index: a[1],
		}
	}
	for _, a := range test.arcs {
		from, to := nodes[a[0]], nodes[a[1]]
		edge := &TestEdge{
			from: from,
			to:   to,
		}
		from.edges = append(from.edges, edge)
		to.edges = append(to.edges, edge)
	}
	// Note: the results are in reverse order.
	path := ShortestPath__10(nodes[test.from], nodes[test.to])
	var got [][2]int
	for _, e := range path {
		got = append(got, [2]int{e.from.index, e.to.index})
	}
	c.Assert(got, qt.DeepEquals, test.want)
}
