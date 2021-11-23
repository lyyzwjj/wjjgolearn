package graph

import (
	"errors"
	"fmt"
	common "github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_common"
	list "github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_list"
	"strings"
)

type Graph interface {
	VerticesSize() int
	EdgesSize() int
	AddVertex(v interface{})
	RemoveVertex(v interface{})
	AddEdge(from, to interface{})
	AddEdgeWithWeight(from, to interface{}, weight Weight)
	RemoveEdge(from, to interface{})
	breadthFirstSearch(v interface{})
	depthFirstSearch(v interface{})
	topologicalSort() []interface{}
	minimumSpanningTree() map[EdgeInfo]interface{}
	prim() map[EdgeInfo]interface{}
	kruskal(vertexValueComparator common.Comparator) map[EdgeInfo]interface{}
	shortestPath(v interface{}) map[interface{}]*PathInfo
	dijkstra(v interface{}) map[interface{}]*PathInfo
	bellmanFord(v interface{}) map[interface{}]*PathInfo
	shortestPathAllVertex(vertexValueComparator common.Comparator) (paths map[interface{}]map[interface{}]*PathInfo)
	floyd(vertexValueComparator common.Comparator) (paths map[interface{}]map[interface{}]*PathInfo)
	Print()
}

type Weight interface {
	compare(o Weight) int
	add(o Weight) Weight
	zero() Weight
	getValue() interface{}
	ToString() string
}

type vertex struct {
	value    interface{}
	inEdges  map[*edge]interface{}
	outEdges map[*edge]interface{}
}

func newVertex(value interface{}) *vertex {
	vertex := &vertex{
		value:    value,
		inEdges:  make(map[*edge]interface{}),
		outEdges: make(map[*edge]interface{}),
	}
	return vertex
}

func (v *vertex) ToString() string {
	return fmt.Sprintf("%v", v.value)
}

type edge struct {
	weight Weight
	from   *vertex
	to     *vertex
}

func newEdge(from, to *vertex) *edge {
	edge := &edge{
		from: from,
		to:   to,
	}
	return edge
}

type edgeKey struct {
	from interface{}
	to   interface{}
}

func (e *edge) GetEdgeKey() edgeKey {
	edgeKey := edgeKey{
		from: e.from.value,
		to:   e.to.value,
	}
	return edgeKey
}
func (e *edge) info() EdgeInfo {
	edgeInfo := EdgeInfo{
		from:   e.from.value,
		to:     e.to.value,
		weight: e.weight,
	}
	return edgeInfo
}
func NewEdgeKey(from, to interface{}) edgeKey {
	edgeKey := edgeKey{
		from: from,
		to:   to,
	}
	return edgeKey
}

type EdgeInfo struct {
	from   interface{}
	to     interface{}
	weight Weight
}

type PathInfo struct {
	weight    Weight
	edgeInfos list.List
}

func NewPathInfo() *PathInfo {
	return &PathInfo{
		edgeInfos: list.NewLinkedList(),
	}
}
func NewPathInfoWithWeight(weight Weight) *PathInfo {
	return &PathInfo{
		weight:    weight,
		edgeInfos: list.NewLinkedList(),
	}
}

func (p *PathInfo) ToString() (str string) {
	edgeInfosStr := "["
	for _, e := range p.edgeInfos.GetAll() {
		edgeInfo := e.(EdgeInfo)
		edgeInfosStr += edgeInfo.ToString() + ", "
	}
	if strings.HasSuffix(edgeInfosStr, ", ") {
		edgeInfosStr = edgeInfosStr[0 : len(edgeInfosStr)-2]
	}
	edgeInfosStr = edgeInfosStr + "]"
	str = fmt.Sprintf("PathInfo [weight=%v, edgeInfos=%v]", p.weight.ToString(), edgeInfosStr)
	return
}

func (e *edge) ToString() (str string) {
	str = fmt.Sprintf("Edge {from=%#v, to=%#v, weight=%#v}", e.from.ToString(), e.to.ToString(), e.weight.ToString())
	return
}
func (e *EdgeInfo) ToString() (str string) {
	str = fmt.Sprintf("EdgeInfo [from=%v, to=%v, weight=%v]", e.from, e.to, e.weight.ToString())
	return
}

var weightValueErr = errors.New("weight value type must be int or float64")

type WeightImpl struct {
	Value interface{}
}

func NewWeightImpl(value interface{}) Weight {
	weightImpl := &WeightImpl{
		Value: value,
	}
	return weightImpl
}

func (w WeightImpl) getValue() interface{} {
	return w.Value
}

func (w WeightImpl) ToString() string {
	return fmt.Sprintf("%v", w.Value)
}

func (w WeightImpl) compare(o Weight) (compareResult int) {
	fw, fwOk := w.getValue().(float64)
	fo, foOk := o.getValue().(float64)
	if fwOk && foOk {
		if fw > fo {
			compareResult = 1
		} else if fw == fo {
			compareResult = 0
		} else {
			compareResult = -1
		}
		return
	}
	iw, iwOk := w.getValue().(int)
	io, ioOk := o.getValue().(int)
	if iwOk && ioOk {
		if iw > io {
			compareResult = 1
		} else if iw == io {
			compareResult = 0
		} else {
			compareResult = -1
		}
		return
	}
	panic(weightValueErr)
}

func (w WeightImpl) add(o Weight) Weight {
	var Value interface{}
	Value = nil
	fw, fwOk := w.getValue().(float64)
	fo, foOk := o.getValue().(float64)
	if fwOk && foOk {
		Value = fw + fo
	}
	iw, iwOk := w.getValue().(int)
	io, ioOk := o.getValue().(int)
	if iwOk && ioOk {
		Value = iw + io
	}
	if Value != nil {
		return &WeightImpl{
			Value: Value,
		}
	}
	panic(weightValueErr)

}

func (w WeightImpl) zero() Weight {
	var Value interface{}
	Value = nil
	_, fwOk := w.getValue().(float64)
	if fwOk {
		Value = 0.0
	}
	_, iwOk := w.getValue().(int)
	if iwOk {
		Value = 0
	}
	if Value != nil {
		return &WeightImpl{
			Value: Value,
		}
	}
	panic(weightValueErr)
}
