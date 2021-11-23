package queue

import list "github.com/lyyzwjj/wjjgolearn/algorithms-data-structures/first-stage/learn/_01_list"

type Queue interface {
	Size() int
	EnQueue(element interface{})
	DeQueue() (element interface{})
	//	deQueueObj(obj interface{})
	Front() (element interface{})
	IsEmpty() bool
	Clear()
}

type BaseQueue struct {
	LinkedList *list.LinkedList
}

func NewBaseQueue() Queue {
	return &BaseQueue{
		LinkedList: list.NewLinkedList(),
	}
}
func (b *BaseQueue) Size() int {
	return b.LinkedList.Size()
}
func (b *BaseQueue) EnQueue(element interface{}) {
	b.LinkedList.Add(element)
}
func (b *BaseQueue) DeQueue() (element interface{}) {
	return b.LinkedList.Remove(0)
}

//func (b *BaseQueue) DeQueueObj(obj interface{}) {
//	obj = b.LinkedList.Remove(0)
//	return
//}
func (b *BaseQueue) Front() (element interface{}) {
	return b.LinkedList.Get(0)
}
func (b *BaseQueue) IsEmpty() bool {
	return b.LinkedList.IsEmpty()
}
func (b *BaseQueue) Clear() {
	b.LinkedList.Clear()
}
