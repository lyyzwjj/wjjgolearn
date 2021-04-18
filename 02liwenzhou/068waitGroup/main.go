package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// waitGroup
// 等待所有的goroutine结束完之后去执行

func f() {
	rand.Seed(time.Now().Unix()) // 不加种子编译后每次生成的随机数是一样的 0-10 保证每次执行的时候都有点不一样
	for i := 0; i < 5; i++ {
		r1 := rand.Int()    // int64随机数
		r2 := rand.Intn(10) // 0<=x<10
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	timeInterval := rand.Intn(500)
	time.Sleep(time.Millisecond * time.Duration(timeInterval))
	fmt.Println(i, timeInterval)
}

var wg sync.WaitGroup // 计数器 类似
// 程序启动之后会创建一个主goroutine去执行
func main() {
	// f()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// wg.Add(1) 此方式 可能第九次很快执行完了  就结束了 轮不到第10次操作了
		go f1(i)
	}
	// ?如何知道这10个goroutine都结束了
	// time.Sleep(?)
	wg.Wait()
}
