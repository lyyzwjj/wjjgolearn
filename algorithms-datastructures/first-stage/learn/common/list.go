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
	defaultCapacity = 10
)

type AbstractList struct {
	size         int
	AddWithIndex func(index int, element int)
	Remove       func(index int) (int, bool)
	Clear        func()
	Get          func(index int) *int
	Set          func(index int, element int) (int, bool)
	IndexOf      func(element int) *int
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
	return al.IndexOf(element) != nil
}
func (al *AbstractList) RemoveElement(element int) {
	al.Remove(*al.IndexOf(element))
}

func (al *AbstractList) rangeCheck(index int) error {
	if index < 0 || index >= al.size {
		return errors.New("Size: " + strconv.Itoa(al.size) + " Index: " + strconv.Itoa(index))
	}
	return nil
}
func (al *AbstractList) rangeCheckForAdd(index int) error {
	if index < 0 || index > al.size {
		return errors.New("Size: " + strconv.Itoa(al.size) + " Index: " + strconv.Itoa(index))
	}
	return nil
}

type ArrayList struct {
	AbstractList
	elements []interface{}
}

func NewArrayListWithCapacity(capacity int) *ArrayList {
	arrayList := new(ArrayList)
	arrayList.elements = make([]interface{}, capacity)
	// println("size: " + strconv.Itoa(len(arrayList.elements)) + "capacity: " + strconv.Itoa(cap(arrayList.elements)))
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
