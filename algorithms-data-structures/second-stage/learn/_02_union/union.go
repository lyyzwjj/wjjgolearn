package union

import common "github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_common"

type UnionFind interface {
	Union(v1, v2 interface{})
	Find(v interface{}) (root interface{})
	IsSame(v1, v2 interface{}) bool
	MakeSet(v interface{})
}

type BaseUnionFind struct {
	nodes      map[interface{}]*node
	comparator common.Comparator
}

func NewBaseUnionFind(comparator common.Comparator) UnionFind {
	return &BaseUnionFind{
		nodes:      make(map[interface{}]*node),
		comparator: comparator,
	}
}
func (b *BaseUnionFind) MakeSet(v interface{}) {
	_, ok := b.nodes[v]
	if ok {
		return
	}
	b.nodes[v] = NewNode(v)
}
func (b *BaseUnionFind) Union(v1, v2 interface{}) {
	n1 := b.FindNode(v1)
	n2 := b.FindNode(v2)
	if n1 == nil || n2 == nil {
		return
	}
	if b.comparator(n1.Value, n2.Value) == 0 {
		return
	}
	// 树矮  的挂到 树高的下面
	if n1.Rank < n2.Rank {
		n1.Parent = n2
	} else if n1.Rank > n2.Rank {
		n2.Parent = n1
	} else {
		n1.Parent = n2
		n2.Rank++
	}

}
func (b *BaseUnionFind) Find(v interface{}) (root interface{}) {
	n := b.FindNode(v)
	if n != nil {
		root = n.Value
	}
	return
}

func (b *BaseUnionFind) IsSame(v1, v2 interface{}) bool {
	return b.comparator(b.Find(v1), b.Find(v2)) == 0
}

func (b *BaseUnionFind) FindNode(v interface{}) *node {
	n, ok := b.nodes[v]
	if !ok || n == nil {
		return nil
	}
	for b.comparator(n.Value, n.Parent.Value) != 0 {
		n.Parent = n.Parent.Parent
		n = n.Parent
	}
	return n
}

type node struct {
	Value  interface{}
	Parent *node
	Rank   int
}

func NewNode(value interface{}) *node {
	n := &node{
		Value: value,
		Rank:  1,
	}
	n.Parent = n
	return n
}
