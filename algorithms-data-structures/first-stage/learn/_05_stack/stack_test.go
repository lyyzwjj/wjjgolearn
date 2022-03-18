package stack

import (
	"fmt"
	"testing"
)

func TestBaseStack(t *testing.T) {
	BaseStack := NewBaseStack[int]()
	BaseStack.Push(11)
	BaseStack.Push(22)
	BaseStack.Push(33)
	BaseStack.Push(44)
	fmt.Println(BaseStack.Pop())
	fmt.Println(BaseStack.Pop())
}
