package main

import (
	"fmt"
)

type List[T comparable] interface {
	Add(element T)
}
type BaseList[T comparable] struct {
}

func NewBaseList[T comparable]() *BaseList[T] {
	return &BaseList[T]{}
}
func (b *BaseList[T]) Add(element T) {
	fmt.Println(element)
}
func main() {
	var linkedList List[int]
	linkedList = NewBaseList[int]()
	linkedList.Add(123)
	fmt.Println(linkedList)
}
