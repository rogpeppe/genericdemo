// +build ignore

package graph

contract Graph(n Node, e Edge) {
	var _ []Edge = n.Edges()
	var from, to Node = e.Nodes()
	map[Node]bool{}
}

func ShortestPath(Node, Edge Graph)(from, to Node) []Edge {
	q := map[Node]bool{
		from: true,
	}
	dist := make(map[Node]int)
	
	
}


