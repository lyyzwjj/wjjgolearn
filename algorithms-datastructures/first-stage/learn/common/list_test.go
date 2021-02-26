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
}
