package common

import (
	"fmt"
	"strconv"
)

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
	IndexOf(element interface{}) int
}

type AbstractList struct {
	ElementNotFount int
	size            int
}

func (al *AbstractList) Size() int {
	return al.size
}
func (al *AbstractList) IsEmpty() bool {
	return al.size == 0
}

func (al *AbstractList) outOfBound(index int) {
	fmt.Println("Size: " + strconv.Itoa(al.size) + " Index: " + strconv.Itoa(index))
}
func (al *AbstractList) rangeCheck(index int) {
	if index < 0 || index >= al.size {
		al.outOfBound(index)
	}
}
func (al *AbstractList) rangeCheckForAdd(index int) bool {
	if index < 0 || index > al.size {
		al.outOfBound(index)
	} else {
		return true
	}
}
