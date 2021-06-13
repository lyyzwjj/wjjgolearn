package queue

import (
	"fmt"
	"testing"
)

func TestBaseQueue(t *testing.T) {
	baseQueue := NewBaseQueue()
	baseQueue.EnQueue(11)
	baseQueue.EnQueue(22)
	baseQueue.EnQueue(33)
	baseQueue.EnQueue(44)
	fmt.Println(baseQueue.DeQueue())
	fmt.Println(baseQueue.DeQueue())
}
