package fea_obj_cache

import (
	"fmt"
	"sync"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/9/1 11:29 下午
 * @description sync.pool  缓存不止一个对象  如果没有缓存对象 则从按默认构造方法创建
 */

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object.")
			return 100
		},
	}
	v := pool.Get().(int) // 做格言
	fmt.Println(v)
	pool.Put(3)
	// runtime.GC() // GC 会清除sync.pool中的缓存的对象
	v1, _ := pool.Get().(int)
	fmt.Println(v1)
}
func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new Object.")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
