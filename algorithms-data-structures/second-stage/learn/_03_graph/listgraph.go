package graph

import (
	"errors"
	"fmt"
	common "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_common"
	list "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_list"
	heap "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_02_heap"
	set "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_02_set"
	queue "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_04_queue"
	stack "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_05_stack"
	union "github.com/wzzst310/wjjgolearn/algorithms-data-structures/second-stage/learn/_02_union"
)

type ListGraph struct {
	vertices map[interface{}]*vertex
	edges    set.Set
	edgesMap map[edgeKey]*edge
}

func NewListGraph() Graph {
	listGraph := &ListGraph{
		vertices: make(map[interface{}]*vertex),
		edges:    set.NewHashSet(),
		edgesMap: make(map[edgeKey]*edge),
	}
	return listGraph
}

func (l *ListGraph) VerticesSize() int {
	return len(l.vertices)
}

func (l *ListGraph) EdgesSize() int {
	return l.edges.Size()
}

func (l *ListGraph) AddVertex(v interface{}) {
	_, ok := l.vertices[v]
	if !ok {
		l.vertices[v] = newVertex(v)
	}
}

func (l *ListGraph) RemoveVertex(v interface{}) {
	if vertex, ok := l.vertices[v]; ok {
		delete(l.vertices, v)
		for edge := range vertex.outEdges {
			edgeKey := edge.GetEdgeKey()
			edge := l.edgesMap[edgeKey]
			l.edges.Remove(edge)
			delete(l.edgesMap, edgeKey)
			delete(edge.to.inEdges, edge)
		}
		for edge := range vertex.inEdges {
			edgeKey := edge.GetEdgeKey()
			edge := l.edgesMap[edgeKey]
			l.edges.Remove(edge)
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
	l.edges.Add(edgeKey)
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
	visitedVertices := set.NewHashSet()
	baseQueue := NewVertexBaseQueue()
	baseQueue.EnQueue(beginVertex)
	visitedVertices.Add(beginVertex)
	for !baseQueue.IsEmpty() {
		vertex := baseQueue.DeQueueVertex()
		fmt.Println(vertex.value)
		for edge := range vertex.outEdges {
			if !visitedVertices.Contains(edge.to) {
				baseQueue.EnQueue(edge.to)
				visitedVertices.Add(edge.to)
			}
		}
	}
}
func (l *ListGraph) depthFirstSearch(v interface{}) {
	beginVertex, ok := l.vertices[v]
	if !ok {
		return
	}
	// visitedVertices := make(map[*vertex]interface{})
	visitedVertices := set.NewHashSet()
	baseStack := NewVertexBaseStack()
	baseStack.Push(beginVertex)
	visitedVertices.Add(beginVertex)
	fmt.Println(beginVertex.value)
	for !baseStack.IsEmpty() {
		vertex := baseStack.PopVertex()
		for edge := range vertex.outEdges {
			if !visitedVertices.Contains(edge.to) {
				baseStack.Push(edge.from)
				baseStack.Push(edge.to)
				visitedVertices.Add(edge.to)
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

func (l *ListGraph) minimumSpanningTree() map[EdgeInfo]interface{} {
	return l.prim()
}

func edgeComparator(a, b interface{}) int {
	ia := a.(*edge)
	ib := b.(*edge)
	// fmt.Printf("ia: %v, ib: %v, result: %v\n", ia.weight.ToString(), ib.weight.ToString(), ib.weight.compare(ia.weight))
	return ia.weight.compare(ib.weight)
}

func GetVertexComparator(vertexValueComparator common.Comparator) common.Comparator {
	return func(a, b interface{}) int {
		ia := a.(*vertex)
		ib := b.(*vertex)
		return vertexValueComparator(ib.value, ia.value)
	}
}

func (l *ListGraph) prim() map[EdgeInfo]interface{} {
	edgeSize := len(l.vertices) - 1
	if edgeSize == -1 {
		return nil
	}
	edgeInfos := map[EdgeInfo]interface{}{}
	addedVertices := set.NewHashSet()
	var binaryHeap heap.Heap
	for _, vertex := range l.vertices {
		addedVertices.Add(vertex.value)
		fmt.Println(vertex.value)
		var outEdges []interface{}
		for edge := range vertex.outEdges {
			outEdges = append(outEdges, edge)
		}
		binaryHeap = heap.NewBinaryHeapWithElements(outEdges, edgeComparator)
		break
	}
	for binaryHeap != nil && !binaryHeap.IsEmpty() && len(edgeInfos) < edgeSize {
		edge := binaryHeap.Remove().(*edge)
		fmt.Println(edge.weight.ToString())
		if !addedVertices.Contains(edge.to.value) {
			edgeInfos[edge.info()] = nil
			addedVertices.Add(edge.to.value)
			var outEdges []interface{}
			for v := range edge.to.outEdges {
				outEdges = append(outEdges, v)
			}
			binaryHeap.AddAll(outEdges)
		}
	}
	return edgeInfos
}

func (l *ListGraph) kruskal(vertexValueComparator common.Comparator) map[EdgeInfo]interface{} {
	edgeSize := len(l.vertices) - 1
	if edgeSize == -1 {
		return nil
	}
	edgeInfos := map[EdgeInfo]interface{}{}
	var edges []interface{}
	for _, edge := range l.edgesMap {
		edges = append(edges, edge)
	}
	binaryHeap := heap.NewBinaryHeapWithElements(edges, edgeComparator)
	uf := union.NewBaseUnionFind(GetVertexComparator(vertexValueComparator))
	for _, vertex := range l.vertices {
		uf.MakeSet(vertex)
	}
	for !binaryHeap.IsEmpty() && len(edgeInfos) < edgeSize {
		edge := binaryHeap.Remove().(*edge)
		if uf.IsSame(edge.from, edge.to) {
			continue
		}
		edgeInfos[edge.info()] = nil
		uf.Union(edge.from, edge.to)
	}
	return edgeInfos
}

func (l *ListGraph) shortestPath(v interface{}) map[interface{}]*PathInfo {
	return l.dijkstra(v)
}
func (l *ListGraph) dijkstra(v interface{}) map[interface{}]*PathInfo {
	beginVertex, ok := l.vertices[v]
	if !ok || beginVertex == nil {
		return nil
	}
	selectedPaths := make(map[interface{}]*PathInfo)
	paths := make(map[*vertex]*PathInfo)
	for edge := range beginVertex.inEdges {
		paths[beginVertex] = NewPathInfoWithWeight(edge.weight.zero())
		break
	}
	for len(paths) != 0 {
		minVertex, minPathInfo := getMinPath(paths)
		selectedPaths[minVertex.value] = minPathInfo
		delete(paths, minVertex)
		for edge := range minVertex.outEdges {
			if _, ok := selectedPaths[edge.to.value]; ok {
				continue
			}
			dijkstraRelax(edge, minPathInfo, paths)
		}
	}
	delete(selectedPaths, v)
	return selectedPaths
}

func (l *ListGraph) bellmanFord(v interface{}) map[interface{}]*PathInfo {
	beginVertex, ok := l.vertices[v]
	if !ok || beginVertex == nil {
		return nil
	}
	selectedPaths := make(map[interface{}]*PathInfo)
	for edge := range beginVertex.inEdges {
		selectedPaths[v] = NewPathInfoWithWeight(edge.weight.zero())
		break
	}
	count := len(l.vertices) - 1
	for i := 0; i < count; i++ {
		for _, edge := range l.edgesMap {
			pathInfo := selectedPaths[edge.from.value]
			if pathInfo == nil {
				continue
			}
			bellmanFordRelax(edge, pathInfo, selectedPaths)
		}
	}
	// n-1次对所有边松弛后再松弛一遍
	for _, edge := range l.edgesMap {
		pathInfo := selectedPaths[edge.from.value]
		if pathInfo == nil {
			continue
		}
		if bellmanFordRelax(edge, pathInfo, selectedPaths) {
			fmt.Println("有负权环,找不到最短路径")
			return nil
		}
	}
	delete(selectedPaths, v)
	return selectedPaths
}

func getMinPath(paths map[*vertex]*PathInfo) (minVertex *vertex, minPathInfo *PathInfo) {
	for v, p := range paths {
		if minPathInfo == nil || p.weight.compare(minPathInfo.weight) < 0 {
			minVertex = v
			minPathInfo = p
		}
	}
	return
}

func dijkstraRelax(edge *edge, fromPath *PathInfo, paths map[*vertex]*PathInfo) {
	newWeight := fromPath.weight.add(edge.weight)
	oldPath := paths[edge.to]
	if oldPath != nil {
		if newWeight.compare(oldPath.weight) >= 0 {
			return
		}
		oldPath.edgeInfos.Clear()
	} else {
		oldPath = NewPathInfo()
		paths[edge.to] = oldPath
	}
	oldPath.weight = newWeight
	oldPath.edgeInfos.AddAll(fromPath.edgeInfos.GetAll())
	oldPath.edgeInfos.Add(edge.info())
}

func bellmanFordRelax(edge *edge, fromPath *PathInfo, paths map[interface{}]*PathInfo) bool {
	newWeight := fromPath.weight.add(edge.weight)
	oldPath := paths[edge.to.value]
	if oldPath != nil {
		if newWeight.compare(oldPath.weight) >= 0 {
			return false
		}
		oldPath.edgeInfos.Clear()
	} else {
		oldPath = NewPathInfo()
		paths[edge.to.value] = oldPath
	}
	oldPath.weight = newWeight
	oldPath.edgeInfos.AddAll(fromPath.edgeInfos.GetAll())
	oldPath.edgeInfos.Add(edge.info())
	return true
}

func (l *ListGraph) shortestPathAllVertex(vertexValueComparator common.Comparator) (paths map[interface{}]map[interface{}]*PathInfo) {
	return l.floyd(vertexValueComparator)
}

func (l *ListGraph) floyd(vertexValueComparator common.Comparator) map[interface{}]map[interface{}]*PathInfo {
	paths := make(map[interface{}]map[interface{}]*PathInfo)
	for _, edge := range l.edgesMap {
		m := paths[edge.from.value]
		if m == nil {
			m = make(map[interface{}]*PathInfo)
			paths[edge.from.value] = m
		}
		pathInfo := NewPathInfoWithWeight(edge.weight)
		pathInfo.edgeInfos.Add(edge.info())
		m[edge.to.value] = pathInfo
	}
	for v2 := range l.vertices {
		for v1 := range l.vertices {
			for v3 := range l.vertices {
				if vertexValueComparator(v1, v2) == 0 || vertexValueComparator(v2, v3) == 0 || vertexValueComparator(v1, v3) == 0 {
					continue
				}
				path12 := getPathInfo(v1, v2, paths)
				if path12 == nil {
					continue
				}
				path23 := getPathInfo(v2, v3, paths)
				if path23 == nil {
					continue
				}
				path13 := getPathInfo(v1, v3, paths)
				newPathWeight := path12.weight.add(path23.weight)
				if path13 != nil {
					if newPathWeight.compare(path13.weight) >= 0 {
						continue
					}
					path13.edgeInfos.Clear()
				} else {
					path13 = NewPathInfo()
					paths[v1][v3] = path13
				}
				path13.weight = newPathWeight
				path13.edgeInfos.AddAll(path12.edgeInfos.GetAll())
				path13.edgeInfos.AddAll(path23.edgeInfos.GetAll())
			}
		}
	}
	return paths
}
func getPathInfo(from, to interface{}, paths map[interface{}]map[interface{}]*PathInfo) *PathInfo {
	m := paths[from]
	if m == nil {
		return nil
	}
	return m[to]
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
	if vertex, ok := v.DeQueue().(*vertex); !ok {
		panic(queueValueErr)
	} else {
		return vertex
	}
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
	if vertex, ok := v.Pop().(*vertex); !ok {
		panic(stackValueErr)
	} else {
		return vertex
	}
}
