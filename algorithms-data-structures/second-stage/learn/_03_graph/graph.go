package graph

import (
	"fmt"
	"strconv"
)

type Graph interface {
	VerticesSize() int
	EdgesSize() int
	AddVertex(v int)
	RemoveVertex(v int)
	AddEdge(from, to int)
	AddEdgeWithWeight(from, to int, weight *int)
	RemoveEdge(from, to int)
}

type vertexKey struct {
	value int
}

type vertex struct {
	vertexKey
	inEdges  map[edgeKey]interface{}
	outEdges map[edgeKey]interface{}
}

func (v *vertex) ToString() string {
	return strconv.Itoa(v.value)
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
		inEdges:  make(map[edgeKey]interface{}),
		outEdges: make(map[edgeKey]interface{}),
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
func (e *edge) ToString() (str string) {
	weight := "nil"
	if e.weight != nil {
		weight = strconv.Itoa(*e.weight)
	}
	str = fmt.Sprintf("Edge {from=%#v, to=%#v, weight=%#v}", e.from.ToString(), e.to.ToString(), weight)
	return
}

func newEdge(from, to *vertex) *edge {
	edge := &edge{
		from: from,
		to:   to,
	}
	return edge
}
