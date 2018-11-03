package graph

type Int int

type indexedInt struct {
	x int
	i int
}

type TestNode struct {
	index int
	edges []*TestEdge
}

func (n *TestNode) Edges() []*TestEdge {
	return n.edges
}

type TestEdge struct {
	from, to *TestNode
}

func (n *TestEdge) Nodes() (from, to *TestNode) {
	return n.from, n.to
}
