package _03_graph

type Graph interface {
	VerticesSize() int
	EdgesSize() int
	AddVertex(v int)
	RemoveVertex(v int)
	AddEdge(fromV, toV int)
	AddEdgeWithWeight(fromV, toV, weight int)
	RemoveEdge(fromV, toV int)
}

type vertex struct {
	value int
}

type edge struct {
	value int
}

type AbstractGraph struct {
}
