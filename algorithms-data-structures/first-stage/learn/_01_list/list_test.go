package list

import (
	"fmt"
	"testing"
)

func initList() List {
	var list List
	//  TestArrayList
	//	list = NewArrayList()
	//  TestLinkedList
	list = NewLinkedList()
	return list
}
func TestList(t *testing.T) {
	list := initList()
	for i := 0; i <= 10; i++ {
		list.Add(i)
	}
	fmt.Printf("%#v\n", list)
	fmt.Println(list.Get(2))
	fmt.Println(list.Contains(2))
	fmt.Println(list.Remove(2))
	fmt.Println(list.Contains(2))
	fmt.Println(list.Size())
	list.Clear()
	fmt.Printf("%#v\n", list)
}
