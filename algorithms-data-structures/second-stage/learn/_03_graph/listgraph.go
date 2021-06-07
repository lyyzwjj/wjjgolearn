package graph

type ListGraph struct {
	vertices map[int]*vertex
	// verticesMap map[vertexKey]*vertex
	edges    map[edgeKey]interface{}
	edgesMap map[edgeKey]*edge
}

func (l *ListGraph) VerticesSize() int {
	return len(l.vertices)
}

func (l *ListGraph) EdgesSize() int {
	return len(l.edges)
}

func (l *ListGraph) AddVertex(v int) {
	_, ok := l.vertices[v]
	if !ok {
		l.vertices[v] = newVertex(v)
	}
}

func (l *ListGraph) RemoveVertex(v int) {
	vertex, ok := l.vertices[v]
	if ok {
		delete(l.vertices, v)
	} else {
		println(vertex)
	}
}

func (l *ListGraph) AddEdge(from, to int) {
	l.AddEdgeWithWeight(from, to, nil)
}

func (l *ListGraph) AddEdgeWithWeight(from, to int, weight *int) {
	fromVertex, ok := l.vertices[from]
	if !ok {
		fromVertex = newVertex(from)
		l.vertices[from] = fromVertex
	}
	toVertex, ok := l.vertices[to]
	if !ok {
		toVertex = newVertex(to)
		l.vertices[to] = toVertex
	}
	edge := newEdge(fromVertex, toVertex)
	edge.weight = weight
	// 删了重新来
	edgeKey := edge.GetEdgeKey()
	l.edgesMap[edgeKey] = edge
	//delete(l.edgesMap, edgeKey)
	//delete(fromVertex.outEdges, edgeKey)
	//delete(toVertex.inEdges, edgeKey)

}

func (l *ListGraph) RemoveEdge(from, to int) {

}
