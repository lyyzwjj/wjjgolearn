package hba_unsafe

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

/**
 * @author  wjj
 * @date  2020/9/6 5:22 上午
 * @description unsafe类
 */

func TestUnsafe(t *testing.T) {
	i := 10
	// 强制将int类型转成 float64
	f := *(*float64)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	t.Log(f)
}

// The cases is suitable for unsafe
type MyInt int

//合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := (*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

//原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	writeDataFn := func() {
		//data := []int{}
		var data []int
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		// atomic.StorePointer(*unsafe.Pointer())
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}
	var wg sync.WaitGroup
	writeDataFn()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
}
