package list

import (
	"errors"
	"strconv"
)

//type List interface {
//	Add(element interface{})
//	AddWithIndex(index int, element interface{})
//	Remove(index int) interface{}
//	RemoveElement(element interface{})
//	Clear()
//	Size() int
//	IsEmpty() bool
//	Contains(element interface{}) bool
//	Get(index int) interface{}
//	Set(index int, element interface{}) interface{}
//	IndexOf(element interface{}) int
//}

const (
	defaultCapacity = 10
	elementNotFound = -1
)

type List interface {
	Add(element interface{})
	AddWithIndex(index int, element interface{})
	Remove(index int) (element interface{})
	RemoveElement(element interface{}) int
	Clear()
	Size() int
	IsEmpty() bool
	Contains(element interface{}) bool
	Get(index int) (element interface{})
	// GetObj(index int, obj interface{})
	Set(index int, element interface{}) (oldElement interface{})
	IndexOf(element interface{}) int
}
type BaseList struct {
	size         int
	AddWithIndex func(index int, element interface{})
	Remove       func(index int) (element interface{})
	IndexOf      func(element interface{}) int
}

func (b *BaseList) Add(element interface{}) {
	b.AddWithIndex(b.size, element)
}

func (b *BaseList) Size() int {
	return b.size
}

func (b *BaseList) IsEmpty() bool {
	return b.size == 0
}

func (b *BaseList) Contains(element interface{}) bool {
	return b.IndexOf(element) != -1
}

func (b *BaseList) RemoveElement(element interface{}) int {
	b.Remove(b.IndexOf(element))
	return -1
}

func (b *BaseList) rangeCheck(index int) {
	if index < 0 || index >= b.size {
		b.outOfBound(index)
	}
}

func (b *BaseList) rangeCheckForAdd(index int) {
	if index < 0 || index > b.size {
		b.outOfBound(index)
	}
}

func (b *BaseList) outOfBound(index int) {
	panic(errors.New("Size: " + strconv.Itoa(b.size) + " Index: " + strconv.Itoa(index)))
}
