package _00_common

type node struct {
	element int
	next    *node
	prev    *node
}

func NewNode(element int, prev, next *node) *node {
	return &node{element, prev, next}
}

type LinkedList struct {
	BaseList
	first *node
	last  *node
}

func NewLinkedList() *LinkedList {
	linkedList := &LinkedList{}
	linkedList.BaseList.AddWithIndex = linkedList.AddWithIndex
	linkedList.BaseList.Remove = linkedList.Remove
	linkedList.BaseList.IndexOf = linkedList.IndexOf
	return linkedList
}

func (l *LinkedList) Get(index int) int {
	return l.node(index).element
}

func (l *LinkedList) Set(index int, element int) int {
	node := l.node(index)
	old := node.element
	node.element = element
	return old
}

func (l *LinkedList) IndexOf(element int) int {
	node := l.first
	for i := 0; i < l.size; i++ {
		if element == node.element {
			return i
		}
		node = node.next
	}
	return elementNotFound
}

func (l *LinkedList) AddWithIndex(index int, element int) {
	l.rangeCheckForAdd(index)
	if index == l.size { // 往最后一个元素时
		oldLast := l.last
		l.last = NewNode(element, l.last, nil)
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

func (l *LinkedList) Remove(index int) int {
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

func (l *LinkedList) Clear() {
	l.size = 0
	l.first = nil
	l.last = nil
}

func (l *LinkedList) node(index int) (node *node) {
	l.rangeCheck(index)
	if index < l.size>>1 { // 从头遍历
		node = l.first
		for i := 0; i < index; i++ {
			node = node.next
		}
	} else {
		node := l.last
		for i := 0; i < index; i++ {
			node = node.prev
		}
	}
	return
}
