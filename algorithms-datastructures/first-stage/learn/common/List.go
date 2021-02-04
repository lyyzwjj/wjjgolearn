package common

type List interface {
	Add(element interface{})
	AddWithIndex(index int, element interface{})
	Remove(index int) interface{}
	RemoveElement(element interface{})
	Clear()
	Size() int
	IsEmpty() bool
	Contains(element interface{}) bool
	Get(index int) interface{}
	Set(index int, element interface{}) interface{}
	IndexOf(element interface{})
}

type AbstractList struct {
	 ElementNotFount int
	 size int
}
