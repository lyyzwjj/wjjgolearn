package fca_util_all_done

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/9/1 1:25 上午
 * @description 所有任务都完成
 */
func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}
func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner) // 解决办法就是变成 buffchan
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n" // 等到所有值返回了才返回
	}
	return finalRet
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 输出当前系统中的协程数
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine()) // 此处有10个线程
}
