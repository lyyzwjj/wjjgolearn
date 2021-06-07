package _00_common

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
	Remove       func(index int) *int
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
func (al *BaseList) RemoveElement(element int) {
	al.Remove(al.IndexOf(element))
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

type LinkedList struct {
	BaseList
}

func (ll *LinkedList) Get(index int) *int {
	ll.rangeCheck(index)
	return nil
}
func (ll *LinkedList) Set(index, element int) {
	ll.rangeCheck(index)
}
func (arrayList *LinkedList) IndexOf(element int) int {
	return -1
}

type ArrayList struct {
	BaseList
	elements []interface{}
}

func NewArrayListWithCapacity(capacity int) *ArrayList {
	arrayList := new(ArrayList)
	arrayList.elements = make([]interface{}, capacity)
	// println("size: " + strconv.Itoa(len(arrayList.elements)) + "capacity: " + strconv.Itoa(cap(arrayList.elements)))
	arrayList.BaseList.AddWithIndex = arrayList.AddWithIndex
	arrayList.BaseList.Remove = arrayList.Remove
	arrayList.BaseList.IndexOf = arrayList.IndexOf
	return arrayList
}
func NewArrayList() *ArrayList {
	return NewArrayListWithCapacity(defaultCapacity)
}
func (arrayList *ArrayList) ensureCapacity(capacity int) {
	oldCapacity := cap(arrayList.elements)
	if oldCapacity >= capacity {
		return
	}
	newCapacity := oldCapacity + (oldCapacity >> 1)
	newElements := make([]interface{}, newCapacity)
	for i := 0; i < arrayList.size; i++ {
		newElements[i] = arrayList.elements[i]
	}
	arrayList.elements = newElements
}
func (arrayList *ArrayList) AddWithIndex(index int, element int) {
	if err := arrayList.rangeCheckForAdd(index); err != nil {
		fmt.Println(err)
		return
	}
	arrayList.ensureCapacity(arrayList.size + 1)
	for i := arrayList.size; i > index; i-- {
		arrayList.elements[i] = arrayList.elements[i-1]
	}
	arrayList.elements[index] = element
	arrayList.size++
}

func (arrayList *ArrayList) Remove(index int) (int, bool) {
	if err := arrayList.rangeCheck(index); err != nil {
		fmt.Println(err)
		return 0, false
	}
	element := arrayList.elements[index]
	for i := index; i < arrayList.size-1; i++ {
		arrayList.elements[i] = arrayList.elements[i+1]
	}
	value, _ := element.(int)
	return value, true
}
func (arrayList *ArrayList) Clear() {
	for i := 0; i < arrayList.size; i++ {
		arrayList.elements[i] = nil
	}
	arrayList.size = 0
}
func (arrayList *ArrayList) Get(index int) *int {
	if err := arrayList.rangeCheck(index); err != nil {
		fmt.Println(err)
		return nil
	}
	value, _ := arrayList.elements[index].(int)
	return &value
}
func (arrayList *ArrayList) Set(index int, element int) (int, bool) {
	if err := arrayList.rangeCheck(index); err != nil {
		fmt.Println(err)
		return 0, false
	}
	oldValue, _ := arrayList.elements[index].(int)
	arrayList.elements[index] = element
	return oldValue, true
}
func (arrayList *ArrayList) IndexOf(element int) *int {
	for i := 0; i < arrayList.size; i++ {
		if arrayList.elements[i] == element {
			return &i
		}
	}
	return nil
}
