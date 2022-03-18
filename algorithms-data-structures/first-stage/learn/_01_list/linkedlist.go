package list

type nodeObj[T Comparable[T]] struct {
	element T
	prev    *nodeObj[T]
	next    *nodeObj[T]
}

func NewNodeObj[T Comparable[T]](element T, prev, next *nodeObj[T]) *nodeObj[T] {
	return &nodeObj[T]{
		element: element,
		prev:    prev,
		next:    next,
	}
}

type LinkedListObj[T Comparable[T]] struct {
	BaseListObj[T]
	first *nodeObj[T]
	last  *nodeObj[T]
}

func NewLinkedListObj[T Comparable[T]]() *LinkedListObj[T] {
	linkedList := &LinkedListObj[T]{}
	linkedList.BaseListObj.AddWithIndex = linkedList.AddWithIndex
	linkedList.BaseListObj.Remove = linkedList.Remove
	linkedList.BaseListObj.IndexOf = linkedList.IndexOf
	return linkedList
}

func (l *LinkedListObj[T]) Get(index int) (element T) {
	return l.node(index).element
}

//func (l *LinkedListObj) GetObj(index int, obj interface{}) {
//	obj = l.nodeObj(index).element
//}

func (l *LinkedListObj[T]) Set(index int, element T) (oldElement T) {
	node := l.node(index)
	old := node.element
	node.element = element
	return old
}

func (l *LinkedListObj[T]) IndexOf(element T) int {
	node := l.first
	for i := 0; i < l.size; i++ {
		if element.Equals(node.element) {
			return i
		}
		node = node.next
	}
	return elementNotFound
}

func (l *LinkedListObj[T]) AddWithIndex(index int, element T) {
	l.rangeCheckForAdd(index)
	if index == l.size { // 往最后一个元素时
		oldLast := l.last
		l.last = NewNodeObj[T](element, l.last, nil)
		if oldLast == nil { //链表为空时加的第一个元素
			l.first = l.last
		} else {
			oldLast.next = l.last
		}
	} else {
		next := l.node(index)
		prev := next.prev // 链表只有一个元素的时候 第一个节点的前节点为空
		node := NewNodeObj(element, prev, next)
		next.prev = node
		if prev == nil { // 在第一个元素插入
			l.first.next = node // 闭环
			// l.first = nodeObj
		} else {
			prev.next = node
		}
	}
	l.size++
}

func (l *LinkedListObj[T]) Remove(index int) (element T) {
	l.rangeCheck(index)
	node := l.node(index)
	prev := node.prev
	next := node.next
	if prev == nil { //删首节点处理
		l.first = next
	} else {
		prev.next = next
	}
	if next == nil { //删尾节点处理
		l.last = prev
	} else {
		next.prev = prev
	}
	l.size--
	return node.element
}

func (l *LinkedListObj[T]) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}
func (l *LinkedListObj[T]) GetAll() (lists []T) {
	node := l.first
	for node != nil {
		lists = append(lists, node.element)
		node = node.next
	}
	return
}
func (l *LinkedListObj[T]) node(index int) (node *nodeObj[T]) {
	l.rangeCheck(index)
	if index < l.size>>1 { // 从头遍历
		node = l.first
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		node = l.last
		for i := l.size - 1; i > index; i-- {
			node = node.prev
		}
	}
	return
}

type node[T comparable] struct {
	element T
	prev    *node[T]
	next    *node[T]
}

func NewNode[T comparable](element T, prev, next *node[T]) *node[T] {
	return &node[T]{
		element: element,
		prev:    prev,
		next:    next,
	}
}

type LinkedList[T comparable] struct {
	BaseList[T]
	first *node[T]
	last  *node[T]
}

func NewLinkedList[T comparable]() List[T] {
	linkedList := &LinkedList[T]{}
	linkedList.BaseList.AddWithIndex = linkedList.AddWithIndex
	linkedList.BaseList.Remove = linkedList.Remove
	linkedList.BaseList.IndexOf = linkedList.IndexOf
	return linkedList
}

func (l *LinkedList[T]) Get(index int) (element T) {
	return l.node(index).element
}

//func (l *LinkedList) GetObj(index int, obj interface{}) {
//	obj = l.node(index).element
//}

func (l *LinkedList[T]) Set(index int, element T) (oldElement T) {
	node := l.node(index)
	old := node.element
	node.element = element
	return old
}

func (l *LinkedList[T]) IndexOf(element T) int {
	node := l.first
	for i := 0; i < l.size; i++ {
		if element == node.element {
			return i
		}
		node = node.next
	}
	return elementNotFound
}

func (l *LinkedList[T]) AddWithIndex(index int, element T) {
	l.rangeCheckForAdd(index)
	if index == l.size { // 往最后一个元素时
		oldLast := l.last
		l.last = NewNode[T](element, l.last, nil)
		if oldLast == nil { //链表为空时加的第一个元素
			l.first = l.last
		} else {
			oldLast.next = l.last
		}
	} else {
		next := l.node(index)
		prev := next.prev // 链表只有一个元素的时候 第一个节点的前节点为空
		node := NewNode(element, prev, next)
		next.prev = node
		if prev == nil { // 在第一个元素插入
			l.first.next = node // 闭环
			// l.first = node
		} else {
			prev.next = node
		}
	}
	l.size++
}

func (l *LinkedList[T]) Remove(index int) (element T) {
	l.rangeCheck(index)
	node := l.node(index)
	prev := node.prev
	next := node.next
	if prev == nil { //删首节点处理
		l.first = next
	} else {
		prev.next = next
	}
	if next == nil { //删尾节点处理
		l.last = prev
	} else {
		next.prev = prev
	}
	l.size--
	return node.element
}

func (l *LinkedList[T]) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}
func (l *LinkedList[T]) GetAll() (lists []T) {
	node := l.first
	for node != nil {
		lists = append(lists, node.element)
		node = node.next
	}
	return
}
func (l *LinkedList[T]) node(index int) (node *node[T]) {
	l.rangeCheck(index)
	if index < l.size>>1 { // 从头遍历
		node = l.first
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		node = l.last
		for i := l.size - 1; i > index; i-- {
			node = node.prev
		}
	}
	return
}
