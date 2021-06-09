package main

import (
	"fmt"
	"testing"
)

type A struct {
	p int
}
type B struct {
	a1, a2 *A
}
type C struct {
	a1, a2 A
}

func TestCustomizeKey(t *testing.T) {
	mb := make(map[B]string)
	b1 := B{
		a1: &A{
			p: 2,
		},
		a2: &A{
			p: 3,
		},
	}
	mb[b1] = "b1"
	b2 := B{
		a1: &A{
			p: 2,
		},
		a2: &A{
			p: 3,
		},
	}
	mb[b2] = "b2"
	for k, v := range mb {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
	fmt.Println(len(mb))

	mc := make(map[C]string)
	c1 := C{
		a1: A{
			p: 2,
		},
		a2: A{
			p: 3,
		},
	}
	mc[c1] = "c1"
	c2 := C{
		a1: A{
			p: 2,
		},
		a2: A{
			p: 3,
		},
	}
	mc[c2] = "c2"
	for k, v := range mc {
		fmt.Printf("k: %v, v: %v\n", k, v)
	}
	fmt.Println(len(mc))
}
