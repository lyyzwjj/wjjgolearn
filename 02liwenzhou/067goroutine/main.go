package main

import (
	"fmt"
	"time"
)

// goroutine
// goroutine对应的函数结束了,goroutine结束了
// main函数执行完了,有main函数创建的那些goroutine都结束了

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	for i := 0; i < 10000; i++ {
		// go hello(i) // 开启一个单独的goroutine去执行hello函数(任务)
		// 匿名函数
		go func(j int) {
			fmt.Println("hello", j) // 用的是函数参数的那个i,不是外面的i
			// fmt.Println("hello", i) // 外部作用域的
		}(i)
	}
	fmt.Println("main")
	time.Sleep(time.Second)
	// main函数结束了 由main函数启动的goroutine也都结束了

}
