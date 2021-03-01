package eab_csp

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/8/30 10:16 下午
 * @description
	Golang实现了 CSP 并发模型做为并发基础，底层使用goroutine做为并发实体 M内核线程  P调度协调 G goroutine
	https://www.cnblogs.com/sunsky303/p/9115530.html
	Go实现了两种并发形式。第一种是大家普遍认知的：多线程共享内存。其实就是Java或者C++等语言中的多线程开发。另外一种是Go语言特有的，也是Go语言推荐的：CSP（communicating sequential processes）并发模型。
	CSP并发模型是在1970年左右提出的概念，属于比较新的概念，不同于传统的多线程通过共享内存来通信，CSP讲究的是“以通信的方式来共享内存”。
	请记住下面这句话：
		Do not communicate by sharing memory; instead, share memory by communicating.
		“不要以共享内存的方式来通信，相反，要通过通信来共享内存。”
	普通的线程并发模型，就是像Java、C++、或者Python，他们线程间通信都是通过共享内存的方式来进行的。
	非常典型的方式就是，在访问共享数据（例如数组、Map、或者某个结构体或对象）的时候，通过锁来访问，
	因此，在很多时候，衍生出一种方便操作的数据结构，叫做“线程安全的数据结构”。
	例如Java提供的包”java.util.concurrent”中的数据结构。Go中也实现了传统的线程并发模型。
*/

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}
func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}
func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retCh := make(chan string) // 普通的chan 里面的类型
	// retCh := make(chan string, 1) // buffered chan
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret                   // 把结果放回到 retChannel里面
		fmt.Println("service exited.") // 如果是普通chan要等取了值之后才会执行  buffer的chan不会阻塞立即释放
	}()
	return retCh
}
func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) // 从chan取数据
	time.Sleep(time.Second * 1)
}
