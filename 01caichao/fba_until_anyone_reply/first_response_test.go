package fba_until_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/9/1 1:12 上午
 * @description 任何一个完成即可
 */

func runTask(id int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("The result is from %d", id)
}

func FirstResponse() string {
	numOfRunner := 10
	// ch := make(chan string) // 阻塞chan
	ch := make(chan string, numOfRunner) // 解决办法就是变成 buff chan
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	return <-ch // 其他的协程因为没有地方获取数据 就一直卡在此处
}

func TestFirstResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine()) // 输出当前系统中的协程数
	t.Log(FirstResponse())
	time.Sleep(time.Second * 1)
	t.Log("After:", runtime.NumGoroutine())
}
