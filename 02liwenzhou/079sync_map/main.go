package main

import (
	"fmt"
	"strconv"
	"sync"
)

// Go内置的map不是并发安全的  sync.map  开心即用 不用make

var m = make(map[string]int)

// var lock = sync.Mutex{}

func get(key string) int {
	return m[key]
}
func set(key string, value int) {
	m[key] = value
}
func main1() {
	wg := sync.WaitGroup{}
	// fatal error: concurrent map writes 并发的对一个map去写  一般go编译器支持20一下的 本电脑3个以上就不行了
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			// lock.Lock() 加锁可以解决锁的问题
			set(key, n)
			// lock.Unlock()
			fmt.Printf("key:=%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}

var m2 = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			m2.Store(key, n)         // 必须使用sync.Map内置的Store方法设置键值对
			value, _ := m2.Load(key) // 必须使用sync.Map提供的Load方法根据key取值
			fmt.Printf("key:=%v,v:=%v\n", key, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
