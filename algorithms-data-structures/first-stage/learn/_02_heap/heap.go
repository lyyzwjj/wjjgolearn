package heap

import (
	"errors"
	common "github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/first-stage/learn/_00_common"
)

type Heap interface {
	Size() int
	IsEmpty() bool
	Clear()
	Add(element interface{})
	Get() (element interface{})
	Remove() (root interface{})
	Replace(element interface{}) (oldElement interface{})
	AddAll(elements []interface{})
}

const (
	defaultCapacity = 10
)

type BinaryHeap struct {
	size       int
	elements   []interface{}
	comparator common.Comparator
}

func NewBinaryHeap(comparator common.Comparator) Heap {
	binaryHeap := &BinaryHeap{
		comparator: comparator,
		elements:   make([]interface{}, defaultCapacity),
	}
	return binaryHeap
}
func NewBinaryHeapWithElements(elements []interface{}, comparator common.Comparator) Heap {
	binaryHeap := &BinaryHeap{
		comparator: comparator,
	}
	if elements == nil || len(elements) == 0 {
		binaryHeap.elements = make([]interface{}, defaultCapacity)
	} else {
		size := len(elements)
		capacity := size
		binaryHeap.size = size
		if capacity < defaultCapacity {
			capacity = defaultCapacity
		}
		binaryHeap.elements = make([]interface{}, capacity)
		for index, element := range elements {
			binaryHeap.elements[index] = element
		}
	}
	binaryHeap.heapify()
	return binaryHeap
}

func (b *BinaryHeap) Size() int {
	return b.size
}

func (b *BinaryHeap) IsEmpty() bool {
	return b.size == 0
}
func (b *BinaryHeap) Clear() {
	b.elements = []interface{}{}
	b.size = 0
}
func (b *BinaryHeap) Add(element interface{}) {
	b.elementNotNullCheck(element)
	b.ensureCapacity(b.size + 1)
	b.elements[b.size] = element
	b.size++
	b.siftUp(b.size - 1)

}
func (b *BinaryHeap) Get() (element interface{}) {
	b.emptyCheck()
	element = b.elements[0]
	return
}
func (b *BinaryHeap) Remove() (root interface{}) {
	b.emptyCheck()
	lastIndex := b.size - 1
	b.size = lastIndex
	root = b.elements[0]
	b.elements[0] = b.elements[lastIndex]
	b.elements[lastIndex] = nil
	b.siftDown(0)
	return
}
func (b *BinaryHeap) Replace(element interface{}) (oldElement interface{}) {
	b.elementNotNullCheck(element)
	if b.size == 0 {
		b.elements[0] = element
		b.size++
	} else {
		oldElement = b.elements[0]
		b.elements[0] = element
		b.siftDown(0)
	}
	return
}
func (b *BinaryHeap) AddAll(elements []interface{}) {
	for _, element := range elements {
		b.Add(element)
	}
}

// heapify 批量建堆
func (b *BinaryHeap) heapify() {
	// 自上而下的上滤  每次向最后插入元素上滤
	/*for i := 1; i < b.size; i++ {
		b.siftUp(i)
	}*/
	// 自下而上的下滤  左右子节点都建好堆后合并
	for i := (b.size >> 1) - 1; i >= 0; i-- {
		b.siftDown(i)
	}
}
func (b *BinaryHeap) siftUp(index int) { // 从这个元素开始上滤
	element := b.elements[index]
	for index > 0 {
		parentIndex := (index - 1) >> 1
		parent := b.elements[parentIndex]
		if b.comparator(element, parent) <= 0 {
			break
		}
		b.elements[index] = parent
		index = parentIndex
	}
	b.elements[index] = element
}
func (b *BinaryHeap) siftDown(index int) { // 下滤
	element := b.elements[index]
	half := b.size >> 1
	for index < half {
		childIndex := index<<1 + 1
		child := b.elements[childIndex]
		rightIndex := childIndex + 1
		if rightIndex < b.size && b.comparator(b.elements[rightIndex], child) > 0 {
			childIndex = rightIndex
			child = b.elements[childIndex]
		}
		if b.comparator(element, child) >= 0 {
			break
		}
		b.elements[index] = child
		index = childIndex
	}
	b.elements[index] = element
}

func (b *BinaryHeap) ensureCapacity(capacity int) {
	oldCapacity := len(b.elements)
	if oldCapacity >= capacity {
		return
	}
	newCapacity := oldCapacity + oldCapacity>>1
	newElements := make([]interface{}, newCapacity)
	for i := 0; i < b.size; i++ {
		newElements[i] = b.elements[i]
	}
	b.elements = newElements
}

type BinaryTrees struct {
}

var heapIsEmptyErr = errors.New("heap is empty")
var elementIsNilErr = errors.New("element must not be null")

func (b *BinaryHeap) emptyCheck() {
	if b.size == 0 {
		panic(heapIsEmptyErr)
	}
}
func (b *BinaryHeap) elementNotNullCheck(element interface{}) {
	if element == nil {
		panic(elementIsNilErr)
	}
}

type BinaryTreeInfo interface {
	root() interface{}
	left(node interface{}) interface{}
	right(node interface{}) interface{}
	string(node interface{}) interface{}
}

func (b *BinaryHeap) root() interface{} {
	// TODO
	return nil
}
func (b *BinaryHeap) left(node interface{}) interface{} {
	// TODO
	return nil
}
func (b *BinaryHeap) right(node interface{}) interface{} {
	// TODO
	return nil
}
func (b *BinaryHeap) string(node interface{}) interface{} {
	// TODO
	return nil
}

type PrintStyle int

const (
	LEVEL_ORDER = PrintStyle(1)
	INORDER     = PrintStyle(2)
)
