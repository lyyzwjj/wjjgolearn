package stack

import list "github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_list"

type Stack[T comparable] interface {
	Size() int
	IsEmpty() bool
	Push(element T)
	Pop() (element T)
	Peek() (element T)
	Clear()
}

type BaseStack[T comparable] struct {
	LinkedList list.List[T]
}

func NewBaseStack[T comparable]() Stack[T] {
	return &BaseStack[T]{
		LinkedList: list.NewLinkedList[T](),
	}
}
func (b *BaseStack[T]) Size() int {
	return b.LinkedList.Size()
}
func (b *BaseStack[T]) IsEmpty() bool {
	return b.LinkedList.IsEmpty()
}
func (b *BaseStack[T]) Push(element T) {
	b.LinkedList.Add(element)
}
func (b *BaseStack[T]) Pop() (element T) {
	return b.LinkedList.Remove(b.LinkedList.Size() - 1)
}
func (b *BaseStack[T]) Peek() (element T) {
	return b.LinkedList.Get(b.LinkedList.Size() - 1)
}
func (b *BaseStack[T]) Clear() {
	b.LinkedList.Clear()
}
