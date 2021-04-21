package main

import (
	"fmt"
	"sync"
)

// channel

var a []int
var b chan int // 需要指定通道中元素的类型
var wg sync.WaitGroup

func noBufChannel() {
	fmt.Println(b)     // nil
	b = make(chan int) // 通道的初始化 不带缓冲区 一定要初始化才能使用
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b // 接收
		fmt.Println("后台goroutine从通道b中取到了", x)
	}()
	b <- 10 // fatal error: all goroutines are asleep - deadlock! 没有goroutine接收
	fmt.Println("10发送到通道b中了...")
	wg.Wait()
}

func bufChannel() {
	fmt.Println(b)        // nil
	b = make(chan int, 1) // 通道的初始化 不带缓冲区 一定要初始化才能使用
	b <- 10               // fatal error: all goroutines are asleep - deadlock! 没有goroutine接收		带通道和不会阻塞
	fmt.Println("10发送到通道b中了...")
	b <- 20 // 当chan只有一个 且缓存大小为1 放两个就卡住了
	fmt.Println("20发送到通道b中了...")
	x := <-b // 接收		// 有缓冲区
	fmt.Println("后台goroutine从通道b中取到了", x)
	close(b) // 手动关闭channel
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	bufChannel()
}
