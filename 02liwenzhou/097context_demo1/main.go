package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context

var (
	wg sync.WaitGroup
)

func f2(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("保德路")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): //  <-chan struct{} 只读的chan
			break FORLOOP
		default:

		}
	}
}

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx) // 传递多级 也可以同时结束 用户请求goroutine  mysql 的goroutine 连接  同时结束
FORLOOP:
	for {
		fmt.Println("周琳")
		time.Sleep(time.Millisecond * 500)
		select {
		case <-ctx.Done(): //  <-chan struct{} 只读的chan
			break FORLOOP
		default:

		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background()) // 造一个取消函数 cancel CancelFunc 取消函数 调用即停止goroutine
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 如何通知子goroutine退出
	cancel()
	wg.Wait()
}
