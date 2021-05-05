package main

import (
	"fmt"
	"sync"
	"time"
)

// 为什么需要context?

var (
	wg       sync.WaitGroup
	notify   bool
	exitChan = make(chan bool, 1) // 一定要记得初始化 否则会一直卡死
)

func f() {
	defer wg.Done()
	for {
		fmt.Println("周琳")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}
func f1() {
	// defer wg.Done()
	defer wg.Done()
FORLOOP: // 冒号表示后面的代码都是属于它的 要跳出就跳出整个的
	for {
		fmt.Println("周琳")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-exitChan:
			break FORLOOP
		default:
		}
	}
}

func main1() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出
	notify = true
	wg.Wait()
}
func main2() {
	// wg.Add(1)
	go f1()
	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出
	exitChan <- true
	// wg.Wait()
}

func main() {
	wg.Add(1)
	go f1()
	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出
	exitChan <- true
	wg.Wait()
}
