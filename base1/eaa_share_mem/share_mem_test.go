package eaa_share_mem

import (
	"sync"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/8/30 6:19 下午
 * @description 锁机制
 */

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d", counter)
}

func TestCounterThreadSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			defer func() { // 配合defer 确保每个协程结束时候能释放锁
				mut.Unlock()
			}()
			mut.Lock()
			counter++
		}()
	}
	time.Sleep(1 * time.Second) // 缺少此处代码 打印出的counter还是少于5000
	t.Logf("counter = %d", counter)
}

func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			defer func() { // 配合defer 确保每个协程结束时候能释放锁
				mut.Unlock()
			}()
			mut.Lock()
			counter++
			wg.Done()
		}()
	}
	wg.Wait() // 阻塞方法
	t.Logf("counter = %d", counter)
}
