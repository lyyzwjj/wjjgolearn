package stack

import list "github.com/wzzst310/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_list"

type Stack interface {
	Size() int
	IsEmpty() bool
	Push(element interface{})
	Pop() (element interface{})
	Peek() (element interface{})
	Clear()
}

type BaseStack struct {
	LinkedList *list.LinkedList
}

func NewBaseStack() Stack {
	return &BaseStack{
		LinkedList: list.NewLinkedList(),
	}
}
func (b *BaseStack) Size() int {
	return b.LinkedList.Size()
}
func (b *BaseStack) IsEmpty() bool {
	return b.LinkedList.IsEmpty()
}
func (b *BaseStack) Push(element interface{}) {
	b.LinkedList.Add(element)
}
func (b *BaseStack) Pop() (element interface{}) {
	return b.LinkedList.Remove(b.LinkedList.Size() - 1)
}
func (b *BaseStack) Peek() (element interface{}) {
	return b.LinkedList.Get(b.LinkedList.Size() - 1)
}
func (b *BaseStack) Clear() {
	b.LinkedList.Clear()
}
