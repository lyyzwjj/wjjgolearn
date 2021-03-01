package eac_select

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/8/30 11:55 下午
 * @description select 多路选择机制 超时控制
 */

func service() string {
	// time.Sleep(time.Millisecond * 50)
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string {
	retCh := make(chan string, 1) // buffered chan
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret                   // 把结果放回到 retChannel里面
		fmt.Println("service exited.") // 如果是普通chan要等取了值之后才会执行  buffer的chan不会阻塞立即释放
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
