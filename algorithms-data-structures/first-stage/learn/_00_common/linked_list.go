package _00_common

type LinkedList struct {
	BaseList
}

func (l *LinkedList) Get(index int) int {
	l.rangeCheck(index)
	return -1
}

func (l *LinkedList) Set(index int, element int) {
	l.rangeCheck(index)
}

func (l *LinkedList) IndexOf(element int) int {
	return -1
}

func (l *LinkedList) AddWithIndex(index int, element int) {
}

func (l *LinkedList) Remove(index int) int {
	return -1
}

func (l *LinkedList) Clear() {
}
