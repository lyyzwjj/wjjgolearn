package main

import (
	"fmt"
	"sync"
)

// channel 练习
// 单向队列 time.Tick
// func Tick(d Duration) <-chan Time {   限制某种通道作为函数参数 只能进行某种操作的时候
// 1. 启动一个goroutine,生成100个数发送到ch1
// 2. 启动一个goroutine,从ch1中取值,计算其平方放到ch2中
// 3. 在main中 从ch2取值打印出来

var wg sync.WaitGroup
var once sync.Once

// func f1(ch1 chan<- int) { // chan<- 定义只能存的channel  单向channel
func f1(ch1 chan int) { // chan<- 定义只能存的channel  单向channel
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	// <-ch1 从ch1中取值直接报错
	close(ch1) // 关闭了 不能写 但能读 一直到没有了 return  ok为false
}

// func f2(ch1 <-chan int, ch2 chan<- int) { // ch1 只能取值 ch2只能存值
func f2(ch1 chan int, ch2 chan int) { // ch1 只能取值 ch2只能存值
	defer wg.Done()
	//for x:=range ch1{
	//	ch2 <-x*x
	//}
	for {
		x, ok := <-ch1
		if !ok { // ch1 没数据了 则返回false
			break
		}
		ch2 <- x * x
	}
	once.Do(func() {
		close(ch2) // 确保某个操作只执行一次
	})
}

// 程序启动之后会创建一个主goroutine去执行
func main() {
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b) // 两个消费者
	// 一个goroutine 存值一个goroutine立马取值 且f1放的速度超过f2取的速度 所以是有序的
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}

}
