package _00_common

type LinkedList struct {
	BaseList
}

func (ll *LinkedList) Get(index int) int {
	ll.rangeCheck(index)
	return -1
}

func (ll *LinkedList) Set(index int, element int) {
	ll.rangeCheck(index)
}

func (ll *LinkedList) IndexOf(element int) int {
	return -1
}

func (ll *LinkedList) AddWithIndex(index int, element int) {
}

func (ll *LinkedList) Remove(index int) int {
	return -1
}

func (ll *LinkedList) Clear() {
}
