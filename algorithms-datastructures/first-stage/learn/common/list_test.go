package common

import (
	"fmt"
	"testing"
)

func TestArrayList(t *testing.T) {
	list := NewArrayList()
	for i := 0; i <= 10; i++ {
		list.Add(i)
	}
	fmt.Println(list)
	fmt.Println(*list.Get(2))
	fmt.Println(list.Contains(2))
	fmt.Println(list.Remove(2))
	fmt.Println(list.Contains(2))
	list.Clear()
	fmt.Println(list)
}
