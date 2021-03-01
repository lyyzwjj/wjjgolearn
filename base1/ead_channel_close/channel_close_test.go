package ead_channel_close

import (
	"fmt"
	"sync"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/31 12:13 上午
 * @description 生产者消费者模型 线程安全锁机制 go 通过csp 通过channel完成数据交互  通道关闭
 * 关闭channel
 */

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // 通道关闭
		// ch <- 11  // 如果往关闭的chan里面继续发消息会报错
		wg.Done()
	}()
}
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			if data, ok := <-ch; ok { // 正常返回 ok 为true  通道关闭 ok返回为false
				fmt.Println(data)
			} else {
				break
			}
		}
		/*for i := 0; i < 11; i++ {  // 第一次接收也没问题  data,ok 其中接收的data是默认值0
			data := <-ch
			fmt.Println(data)
		}*/
		wg.Done()
	}()
}
func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()
}
