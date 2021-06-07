package graph

type Graph interface {
	VerticesSize() int
	EdgesSize() int
	AddVertex(v int)
	RemoveVertex(v int)
	AddEdge(from, to int)
	AddEdgeWithWeight(from, to int, weight *int)
	RemoveEdge(fromV, toV int)
}

type vertexKey struct {
	value int
}

type vertex struct {
	vertexKey
	inEdges  map[edgeKey]interface{}
	outEdges map[edgeKey]interface{}
}

func (v *vertex) GetVertexKey() vertexKey {
	vertexKey := vertexKey{
		value: v.value,
	}
	return vertexKey
}

type GetVertexKey struct {
	vertexKey
	inEdges  map[edgeKey]interface{}
	outEdges map[edgeKey]interface{}
}

func newVertex(value int) *vertex {
	vertex := &vertex{
		vertexKey: vertexKey{
			value: value,
		},
	}
	return vertex
}

type edgeKey struct {
	from vertexKey
	to   vertexKey
}

type edge struct {
	weight *int
	from   *vertex
	to     *vertex
}

func (e *edge) GetEdgeKey() edgeKey {
	edgeKey := edgeKey{
		from: vertexKey{
			value: e.from.value,
		},
		to: vertexKey{
			value: e.to.value,
		},
	}
	return edgeKey
}

func newEdge(from, to *vertex) *edge {
	edge := &edge{
		from: from,
		to:   to,
	}
	return edge
}
