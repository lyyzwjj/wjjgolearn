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
	Add(element int)
	AddWithIndex(index int, element int)
	Remove(index int) int
	RemoveElement(element int) int
	Clear()
	Size() int
	IsEmpty() bool
	Contains(element int) bool
	Get(index int) int
	Set(index, element int) int
	IndexOf(element int) int
}
type BaseList struct {
	size         int
	AddWithIndex func(index, element int)
	Remove       func(index int) int
	IndexOf      func(element int) int
}

func (b *BaseList) Add(element int) {
	b.AddWithIndex(b.size, element)
}

func (b *BaseList) Size() int {
	return b.size
}

func (b *BaseList) IsEmpty() bool {
	return b.size == 0
}

func (b *BaseList) Contains(element int) bool {
	return b.IndexOf(element) != -1
}

func (b *BaseList) RemoveElement(element int) int {
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
