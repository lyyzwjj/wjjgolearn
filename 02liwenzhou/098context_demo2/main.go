package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// context

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

// withCancel 外部调用cancel 方法 内部 ctx.Done就能获取值执行
func withCancel() {
	fmt.Println("==================withCancel==================")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

// withDeadline 直接指定具体结束时间 到了时间结束
func withDeadline() {
	fmt.Println("==================withDeadline==================")
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d) // 50毫秒后超时

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done(): // 50ms 超时 进入此channel
		fmt.Println(ctx.Err())
	}
}

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

// withTimeout 指定具体的延迟时间
func withTimeout() {
	fmt.Println("==================withTimeout==================")
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

type TraceCode string

func worker1(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	traceCode, ok := ctx.Value(key).(string) // 在子goroutine中获取trace code  .()表示类型判断  反射会用到
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func withValue() {
	fmt.Println("==================withValue==================")
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	// 在系统的入口中设置trace code传递给后续启动的goroutine实现日志数据聚合
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "12512312234")
	wg.Add(1)
	go worker1(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}

func main() {
	withCancel()
	withDeadline()
	withTimeout()
	withValue()
}
