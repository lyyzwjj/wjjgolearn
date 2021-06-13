package graph

import (
	"errors"
	"fmt"
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
	Print()
}

type Weight interface {
	compare(o Weight) int
	add(o Weight) Weight
	zero() Weight
	getValue() interface{}
	ToString() string
}

//type vertexKey struct {
//	value interface{}
//}

type vertex struct {
	value    interface{}
	inEdges  map[*edge]interface{}
	outEdges map[*edge]interface{}
}

//func (v *vertex) GetVertexKey() vertexKey {
//	return vertexKey{
//		value: v.value,
//	}
//}

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
func NewEdgeKey(from, to interface{}) edgeKey {
	edgeKey := edgeKey{
		from: from,
		to:   to,
	}
	return edgeKey
}
func (e *edge) ToString() (str string) {
	str = fmt.Sprintf("Edge {from=%#v, to=%#v, weight=%#v}", e.from.ToString(), e.to.ToString(), e.weight.ToString())
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

func (w WeightImpl) compare(o Weight) (compareResult int) {
	fw, fwOk := w.getValue().(float64)
	fo, foOk := o.getValue().(float64)
	if fwOk && foOk {
		if fo > fw {
			compareResult = 1
		} else if fo == fw {
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
		} else if fo == fw {
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
func (w WeightImpl) ToString() string {
	return fmt.Sprintf("%v", w.Value)
}
