package common

import (
	"errors"
	"fmt"
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
	elementNotFound = -1
	defaultCapacity = 10
)

type AbstractList struct {
	size         int
	AddWithIndex func(index int, element int)
	Remove       func(index int) int
	Clear        func()
	Get          func(index int) int
	Set          func(index int, element int) int
	IndexOf      func(element int) int
}

func (al *AbstractList) Add(element int) {
	al.AddWithIndex(al.size, element)
}

func (al *AbstractList) Size() int {
	return al.size
}

func (al *AbstractList) IsEmpty() bool {
	return al.size == 0
}
func (al *AbstractList) Contains(element int) bool {
	return al.IndexOf(element) != elementNotFound
}
func (al *AbstractList) RemoveElement(element int) {
	al.Remove(al.IndexOf(element))
}

func (al *AbstractList) rangeCheck(index int) error {
	if index < 0 || index >= al.size {
		return errors.New("Size: " + strconv.Itoa(al.size) + " Index: " + strconv.Itoa(index))
	}
	return nil
}

type ArrayList struct {
	AbstractList
	elements []int
}

func NewArrayListWithCapacity(capacity int) *ArrayList {
	arrayList := new(ArrayList)
	arrayList.elements = [10]int
	arrayList.AbstractList.AddWithIndex = arrayList.AddWithIndex
	arrayList.AbstractList.Remove = arrayList.Remove
	arrayList.AbstractList.Clear = arrayList.Clear
	arrayList.AbstractList.Get = arrayList.Get
	arrayList.AbstractList.Set = arrayList.Set
	arrayList.AbstractList.IndexOf = arrayList.IndexOf
	return arrayList
}
func NewArrayList() *ArrayList {
	return NewArrayListWithCapacity(defaultCapacity)
}
func (arrayList *ArrayList) AddWithIndex(index int, element int) {

}

func (arrayList *ArrayList) Remove(index int) int {

}
func (arrayList *ArrayList) Clear() {

}
func (arrayList *ArrayList) Get(index int) int {

}
func (arrayList *ArrayList) Set(index int, element int) int {

}
func (arrayList *ArrayList) IndexOf(element int) int {

}
