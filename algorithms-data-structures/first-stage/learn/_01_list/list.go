package list

import (
	"errors"
	"strconv"
)

//type ListObj interface {
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

type Equal[T any] interface {
	Equals(t T) bool
}

type Comparable[T any] interface {
	Equal[T]
	CompareTo(t T) int
}

type ListObj[T Comparable[T]] interface {
	Add(element T)
	AddAll(elements []T)
	AddWithIndex(index int, element T)
	Remove(index int) (element T)
	RemoveElement(element T) int
	Clear()
	Size() int
	IsEmpty() bool
	Contains(element T) bool
	Get(index int) (element T)
	Set(index int, element T) (oldElement T)
	IndexOf(element T) int
	GetAll() (lists []T)
}

type BaseListObj[T Comparable[T]] struct {
	size         int
	AddWithIndex func(index int, element T)
	Remove       func(index int) (element T)
	IndexOf      func(element T) int
}

func (b *BaseListObj[T]) Add(element T) {
	b.AddWithIndex(b.size, element)
}

func (b *BaseListObj[T]) AddAll(elements []T) {
	for _, element := range elements {
		b.Add(element)
	}
}

func (b *BaseListObj[T]) Size() int {
	return b.size
}

func (b *BaseListObj[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *BaseListObj[T]) Contains(element T) bool {
	return b.IndexOf(element) != -1
}

func (b *BaseListObj[T]) RemoveElement(element T) int {
	b.Remove(b.IndexOf(element))
	return -1
}

func (b *BaseListObj[T]) rangeCheck(index int) {
	if index < 0 || index >= b.size {
		b.outOfBound(index)
	}
}

func (b *BaseListObj[T]) rangeCheckForAdd(index int) {
	if index < 0 || index > b.size {
		b.outOfBound(index)
	}
}

func (b *BaseListObj[T]) outOfBound(index int) {
	panic(errors.New("Size: " + strconv.Itoa(b.size) + " Index: " + strconv.Itoa(index)))
}

type List[T comparable] interface {
	Add(element T)
	AddAll(elements []T)
	AddWithIndex(index int, element T)
	Remove(index int) (element T)
	RemoveElement(element T) int
	Clear()
	Size() int
	IsEmpty() bool
	Contains(element T) bool
	Get(index int) (element T)
	Set(index int, element T) (oldElement T)
	IndexOf(element T) int
	GetAll() (lists []T)
}

type BaseList[T comparable] struct {
	size         int
	AddWithIndex func(index int, element T)
	Remove       func(index int) (element T)
	IndexOf      func(element T) int
}

func (b *BaseList[T]) Add(element T) {
	b.AddWithIndex(b.size, element)
}

func (b *BaseList[T]) AddAll(elements []T) {
	for _, element := range elements {
		b.Add(element)
	}
}

func (b *BaseList[T]) Size() int {
	return b.size
}

func (b *BaseList[T]) IsEmpty() bool {
	return b.size == 0
}

func (b *BaseList[T]) Contains(element T) bool {
	return b.IndexOf(element) != -1
}

func (b *BaseList[T]) RemoveElement(element T) int {
	b.Remove(b.IndexOf(element))
	return -1
}

func (b *BaseList[T]) rangeCheck(index int) {
	if index < 0 || index >= b.size {
		b.outOfBound(index)
	}
}

func (b *BaseList[T]) rangeCheckForAdd(index int) {
	if index < 0 || index > b.size {
		b.outOfBound(index)
	}
}

func (b *BaseList[T]) outOfBound(index int) {
	panic(errors.New("Size: " + strconv.Itoa(b.size) + " Index: " + strconv.Itoa(index)))
}
