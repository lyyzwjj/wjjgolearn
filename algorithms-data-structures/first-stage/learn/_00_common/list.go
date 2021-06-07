package _00_common

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
	Set(index, element int)
	IndexOf(element int) int
}

type BaseList struct {
	size         int
	AddWithIndex func(index, element int)
	Remove       func(index int) int
	IndexOf      func(element int) int
}

func (al *BaseList) Add(element int) {
	al.AddWithIndex(al.size, element)
}

func (al *BaseList) Size() int {
	return al.size
}

func (al *BaseList) IsEmpty() bool {
	return al.size == 0
}

func (al *BaseList) Contains(element int) bool {
	return al.IndexOf(element) != -1
}

func (al *BaseList) RemoveElement(element int) int {
	al.Remove(al.IndexOf(element))
	return -1
}

func (al *BaseList) rangeCheck(index int) {
	if index < 0 || index >= al.size {
		al.outOfBound(index)
	}
}

func (al *BaseList) rangeCheckForAdd(index int) {
	if index < 0 || index > al.size {
		al.outOfBound(index)
	}
}

func (al *BaseList) outOfBound(index int) {
	panic(errors.New("Size: " + strconv.Itoa(al.size) + " Index: " + strconv.Itoa(index)))
}
