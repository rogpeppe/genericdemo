package graph

type _p_ShortestPath__t0 interface {
	Edges() []_p_ShortestPath__t1
	_comparable()
}

type _p_ShortestPath__t1 interface {
	Nodes() (from, to _p_ShortestPath__t0)
}

func ShortestPath__g(from, to _p_ShortestPath__t0) []_p_ShortestPath__t1 {
	return nil
}

type _p_item__t0 interface{}
type _p_item__t1 interface{}

type item__g struct {
	n     _p_item__t0
	dist  int
	index int
	edge  _p_item__t1
}
