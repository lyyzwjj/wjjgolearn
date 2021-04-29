package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 原子操作

var x int64
var wg sync.WaitGroup

func add() {
	// x++
	atomic.AddInt64(&x, 1)
	wg.Done()
}
func main() {
	//wg.Add(100000000)
	//for i := 0; i < 100000000; i++ {
	//	go add()
	//}
	//wg.Wait()
	//fmt.Println(x)
	// 比较并交换
	// old 值 和x比较 比较相等则设置新的值
	// ok := atomic.CompareAndSwapInt64(&x, 10, 200)
	ok := atomic.CompareAndSwapInt64(&x, 0, 200)
	fmt.Println(ok, x)
}
