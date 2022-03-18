package set

import (
	"fmt"
	"testing"
)

func initSet[T comparable]() Set[T] {
	var set Set[T]
	set = NewHashSet[T]()
	return set
}

func TestSet(t *testing.T) {
	set := initSet[int]()
	for i := 0; i < 6; i++ {
		set.Add(i)
	}
	fmt.Printf("%#v\n", set)
	set.Add(1)
	set.Add(2)
	set.Add(3)
	set.Remove(5)
	fmt.Println(set.Contains(5))
	fmt.Println(set.Size())
	fmt.Printf("%#v\n", set)
}
