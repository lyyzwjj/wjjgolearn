package graph

import (
	"errors"
	"fmt"
	list "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_list"
	queue "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_04_queue"
	stack "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_05_queue"
)

type ListGraph struct {
	vertices map[interface{}]*vertex
	edges    map[edgeKey]interface{}
	edgesMap map[edgeKey]*edge
}

func NewListGraph() Graph {
	listGraph := &ListGraph{
		vertices: make(map[interface{}]*vertex),
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

func (l *ListGraph) AddVertex(v interface{}) {
	_, ok := l.vertices[v]
	if !ok {
		l.vertices[v] = newVertex(v)
	}
}

func (l *ListGraph) RemoveVertex(v interface{}) {
	vertex, ok := l.vertices[v]
	if ok {
		delete(l.vertices, v)
		for edge := range vertex.outEdges {
			edgeKey := edge.GetEdgeKey()
			edge := l.edgesMap[edgeKey]
			delete(l.edges, edgeKey)
			delete(l.edgesMap, edgeKey)
			delete(edge.to.inEdges, edge)
		}
		for edge := range vertex.inEdges {
			edgeKey := edge.GetEdgeKey()
			edge := l.edgesMap[edgeKey]
			delete(l.edges, edgeKey)
			delete(l.edgesMap, edgeKey)
			delete(edge.from.outEdges, edge)
		}
	} else {
		println(vertex)
	}
}

func (l *ListGraph) AddEdge(from, to interface{}) {
	l.AddEdgeWithWeight(from, to, nil)
}

func (l *ListGraph) AddEdgeWithWeight(from, to interface{}, weight Weight) {
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
	fromVertex.outEdges[edge] = nil
	toVertex.inEdges[edge] = nil

	//delete(l.edgesMap, edgeKey)
	//delete(fromVertex.outEdges, edgeKey)
	//delete(toVertex.inEdges, edgeKey)

}

func (l *ListGraph) RemoveEdge(from, to interface{}) {
	fromVertex, ok := l.vertices[from]
	if !ok {
		return
	}
	toVertex, ok := l.vertices[to]
	if !ok {
		return
	}
	// edge := newEdge(fromVertex, toVertex)
	// edgeKey := edge.GetEdgeKey()
	edgeKey := NewEdgeKey(from, to)
	edge, ok := l.edgesMap[edgeKey]
	if !ok {
		return
	}
	delete(fromVertex.outEdges, edge)
	delete(toVertex.inEdges, edge)
	delete(l.edgesMap, edgeKey)
}

func (l *ListGraph) breadthFirstSearch(v interface{}) {
	beginVertex, ok := l.vertices[v]
	if !ok {
		return
	}
	visitedVertices := make(map[*vertex]interface{})
	baseQueue := NewVertexBaseQueue()
	baseQueue.EnQueue(beginVertex)
	visitedVertices[beginVertex] = nil
	for !baseQueue.IsEmpty() {
		vertex := baseQueue.DeQueueVertex()
		fmt.Println(vertex.value)
		for edge := range vertex.outEdges {
			_, ok := visitedVertices[edge.to]
			if !ok {
				baseQueue.EnQueue(edge.to)
				visitedVertices[edge.to] = nil
			}
		}
	}
}
func (l *ListGraph) depthFirstSearch(v interface{}) {
	beginVertex, ok := l.vertices[v]
	if !ok {
		return
	}
	visitedVertices := make(map[*vertex]interface{})
	baseStack := NewVertexBaseStack()
	baseStack.Push(beginVertex)
	visitedVertices[beginVertex] = nil
	fmt.Println(beginVertex.value)
	for !baseStack.IsEmpty() {
		vertex := baseStack.PopVertex()
		for edge := range vertex.outEdges {
			_, ok := visitedVertices[edge.to]
			if !ok {
				baseStack.Push(edge.from)
				baseStack.Push(edge.to)
				visitedVertices[edge.to] = nil
				fmt.Println(edge.to.value)
			}
		}
	}
}
func (l *ListGraph) topologicalSort() []interface{} {
	var valueList []interface{}
	// 入度为0的vertex容器
	baseQueue := NewVertexBaseQueue()
	// 入度不为0的连带着入度存入一个map中
	ins := make(map[*vertex]int)
	for _, vertex := range l.vertices {
		in := len(vertex.inEdges)
		if in == 0 {
			baseQueue.EnQueue(vertex)
		} else {
			ins[vertex] = in
		}
	}
	for !baseQueue.IsEmpty() {
		queueVertex := baseQueue.DeQueueVertex()
		valueList = append(valueList, queueVertex.value)
		for edge := range queueVertex.outEdges {
			toIn := ins[edge.to] - 1
			if toIn == 0 {
				baseQueue.EnQueue(edge.to)
			} else {
				ins[edge.to] = toIn
			}
		}
	}
	return valueList
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

func (l *ListGraph) PrintEdges(edges map[*edge]interface{}) string {
	str := "["
	if len(edges) > 0 {
		for edge := range edges {
			str = str + edge.ToString() + ","
		}
		if str[len(str)-1:] == "," {
			str = str[:len(str)-1]
		}
	}
	str = str + "]"
	return str
}

type VertexBaseQueue struct {
	queue.BaseQueue
}

func NewVertexBaseQueue() *VertexBaseQueue {
	return &VertexBaseQueue{
		BaseQueue: queue.BaseQueue{
			LinkedList: list.NewLinkedList(),
		},
	}
}

var queueValueErr = errors.New("queue value type must be graph.vertex")

func (v *VertexBaseQueue) DeQueueVertex() *vertex {
	vertex, ok := v.DeQueue().(*vertex)
	if !ok {
		panic(queueValueErr)
	}
	return vertex
}

type VertexBaseStack struct {
	stack.BaseStack
}

func NewVertexBaseStack() *VertexBaseStack {
	return &VertexBaseStack{
		BaseStack: stack.BaseStack{
			LinkedList: list.NewLinkedList(),
		},
	}
}

var stackValueErr = errors.New("stack value type must be graph.vertex")

func (v *VertexBaseStack) PopVertex() *vertex {
	vertex, ok := v.Pop().(*vertex)
	if !ok {
		panic(stackValueErr)
	}
	return vertex
}
