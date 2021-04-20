package main

import (
	"fmt"
	"runtime"
	"sync"
)

// GOMAXPROCS		对应GMP模型里面的P的数量	GOMAXPROCS --> P -->M
// M:N 把m个goroutine分配给n个操作系统线程去执行
// goroutine初始栈的大小是2k

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("A:%d\n", i)
	}
}
func b() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		fmt.Printf("B:%d\n", i)
	}
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	// runtime.GOMAXPROCS(1) // 如果只选一个P 则两个goroutine会排队执行 如果多个P 则两个goroutine会并发执行
	runtime.GOMAXPROCS(2)         // 默认是跑满cpu 跑满物理核心	除非是日志收集等模块 必须限制
	fmt.Println(runtime.NumCPU()) // cpu核心数量
	wg.Add(2)
	go a()
	go b()
	wg.Wait()
}
