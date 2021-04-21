package main

import "fmt"

// 关闭通道 channel

// 程序启动之后会创建一个主goroutine去执行
func main() {
	ch1 := make(chan bool, 2)
	ch1 <- true
	ch1 <- true
	close(ch1)
	//for x := range ch1 {		//  range ch1 内部会帮忙判断返回是不是false
	//	fmt.Println(x)
	//}

	<-ch1
	<-ch1
	x, ok := <-ch1 // 对一个关闭的通道继续去获取值的时候是能取到的不报错 只是ok返回false x 则为对应类型零值
	fmt.Println(x, ok)
	x, ok = <-ch1
	fmt.Println(x, ok)
}
