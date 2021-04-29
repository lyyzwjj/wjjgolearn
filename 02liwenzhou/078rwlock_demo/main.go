package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var (
	x    = 0
	lock sync.RWMutex
	wg   sync.WaitGroup
)

func read() {
	defer wg.Done()
	lock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	lock.RUnlock()
}
func write() {
	defer wg.Done()
	lock.Lock()
	x = x + 1
	time.Sleep(2 * time.Millisecond)
	lock.Unlock()
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	// time.Sleep(time.Second)
	for i := 0; i < 10000000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))

}
