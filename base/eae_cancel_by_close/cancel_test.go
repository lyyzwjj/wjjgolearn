package eae_cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/8/31 1:27 上午
 * @description 任务关闭  取消协程任务
 */
func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}
func cancel_1(cancelChan chan struct{}) {
	// struct {}是一个无元素的结构体类型
	// struct {} {}是一个复合字面量，它构造了一个struct {}类型的值，该值也是空。
	// struct{}  {}
	// type      empty struct    // 类似于匿名类
	cancelChan <- struct{}{}
}
func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}
func TestChannel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelCh) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Cancelled") //所有的协程都收到了 类似于广播机制
		}(i, cancelChan)
	}
	// cancel_1(cancelChan)
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
