package daa_groutine

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/8/30 6:07 下午
 * @description 协程
 */

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		/* // 为什么这个方法能输出 1 -10 因为go的方法都是值传递 复制一份值进函数 每次值就是1-10 所有没有问题 输出0-9
		go func(i int) {
			fmt.Println(i) // 每一个协程里面此处的i是不同的地址
		}(i)
		*/
		go func() {
			fmt.Println(i) // 每一个协程里面此处的i是同一个地址
		}()
		// 输出变成了 很多个10  i这个值  在test协程 和其他启动的协程里面被共享了
	}
	time.Sleep(time.Millisecond * 50)
}
