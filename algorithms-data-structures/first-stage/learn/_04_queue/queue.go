package queue

import list "github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_list"

type Queue[T comparable] interface {
	Size() int
	EnQueue(element T)
	DeQueue() (element T)
	Front() (element T)
	IsEmpty() bool
	Clear()
}

type BaseQueue[T comparable] struct {
	LinkedList list.List[T]
}

func NewBaseQueue[T comparable]() Queue[T] {
	return &BaseQueue[T]{
		LinkedList: list.NewLinkedList[T](),
	}
}
func (b *BaseQueue[T]) Size() int {
	return b.LinkedList.Size()
}
func (b *BaseQueue[T]) EnQueue(element T) {
	b.LinkedList.Add(element)
}
func (b *BaseQueue[T]) DeQueue() (element T) {
	return b.LinkedList.Remove(0)
}

func (b *BaseQueue[T]) Front() (element T) {
	return b.LinkedList.Get(0)
}
func (b *BaseQueue[T]) IsEmpty() bool {
	return b.LinkedList.IsEmpty()
}
func (b *BaseQueue[T]) Clear() {
	b.LinkedList.Clear()
}
