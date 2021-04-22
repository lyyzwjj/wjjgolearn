package main

import "fmt"

// select 当有一个chan 取第一个完成的
func main() {
	ch := make(chan int, 10) // 生产和消费时3:2  生产可能性大于消费  第一次是生产 第二次是生产和消费二选一
	// ch := make(chan int, 1) // 生产和消费只能1:1 生产就不会消费 消费就不会生产
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch: // 从 chan里面取出i
			fmt.Println(x)
		case ch <- i: // 将i又放到chan里面
		}
	}

}
