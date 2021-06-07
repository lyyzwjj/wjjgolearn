package graph

import "fmt"

type ListGraph struct {
	vertices map[int]*vertex
	edges    map[edgeKey]interface{}
	edgesMap map[edgeKey]*edge
}

func NewListGraph() *ListGraph {
	listGraph := &ListGraph{
		vertices: make(map[int]*vertex),
		edges:    make(map[edgeKey]interface{}),
		edgesMap: make(map[edgeKey]*edge),
	}
	return listGraph
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
		for edgeKey := range vertex.outEdges {
			edge := l.edgesMap[edgeKey]
			delete(edge.to.inEdges, edgeKey)
			delete(l.edges, edgeKey)
			delete(l.edgesMap, edgeKey)
		}
		for edgeKey := range vertex.inEdges {
			edge := l.edgesMap[edgeKey]
			delete(l.edges, edgeKey)
			delete(edge.from.outEdges, edgeKey)
			delete(l.edgesMap, edgeKey)
		}
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
	l.edges[edgeKey] = edge
	fromVertex.outEdges[edgeKey] = nil
	toVertex.inEdges[edgeKey] = nil

	//delete(l.edgesMap, edgeKey)
	//delete(fromVertex.outEdges, edgeKey)
	//delete(toVertex.inEdges, edgeKey)

}

func (l *ListGraph) RemoveEdge(from, to int) {
	fromVertex, ok := l.vertices[from]
	if !ok {
		return
	}
	toVertex, ok := l.vertices[to]
	if !ok {
		return
	}
	edge := newEdge(fromVertex, toVertex)
	edgeKey := edge.GetEdgeKey()
	delete(fromVertex.outEdges, edgeKey)
	delete(toVertex.inEdges, edgeKey)
	delete(l.edgesMap, edgeKey)
}

func (l *ListGraph) Print() {
	fmt.Println("vertices==================================================")
	for v, vertex := range l.vertices {
		fmt.Println(v)
		fmt.Println("in---------------------")
		fmt.Println(l.PrintEdges(vertex.inEdges))
		fmt.Println("out---------------------")
		fmt.Println(l.PrintEdges(vertex.outEdges))
	}
	fmt.Println("edges==================================================")
	for _, v := range l.edgesMap {
		fmt.Println(v.ToString())
	}
}

func (l *ListGraph) PrintEdges(edges map[edgeKey]interface{}) string {
	str := "["
	if len(edges) > 0 {
		for edgeKey := range edges {
			edge := l.edgesMap[edgeKey]
			str = str + edge.ToString() + ","
		}
		if str[len(str)-1:] == "," {
			str = str[:len(str)-1]
		}
	}
	str = str + "]"
	return str
}
