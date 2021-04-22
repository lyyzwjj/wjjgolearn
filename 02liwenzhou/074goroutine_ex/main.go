package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
使用goroutine和channel实现一个计算int64随机数各位数和的程序
1. 开启一个goroutine循环生成int64类型的随机数,发送到jobChan
2. 开启24个goroutine从jobChan中取出随机数计算个位数的和,将结果发送到resultChan
3. 主goroutine从resultChan取出结果并打印到终端输出
*/

func worker(id int, jobs <-chan int64, results chan<- int) {
	for j := range jobs {
		fmt.Printf("worker:%d start job:%d\n", id, j)
		addSum := 0
		for j > 0 {
			each := j % 10
			j = j / 10
			addSum += int(each)
		}
		fmt.Printf("worker:%d end job:%d addSum:%d \n", id, j, addSum)
		results <- addSum
	}
}

func myFunc() {
	jobs := make(chan int64, 100)
	results := make(chan int, 100)
	// 开启3个goroutine
	for w := 1; w <= 24; w++ {
		go worker(w, jobs, results)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// 生成100次
	for j := 1; j <= 100; j++ {
		jobs <- r.Int63()
	}
	close(jobs)
	// 输出结果
	for a := 1; a <= 100; a++ {
		r := <-results
		fmt.Println(r)
	}
}

// job ...
type job struct {
	value int64
}

// result ...
type result struct {
	job *job
	sum int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func producer(prod chan<- *job) {
	// 循环生成int64类型的随机数,发送到jobChan
	for {
		x := rand.Int63()
		newJob := &job{x}
		prod <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func cons(prod <-chan *job, res chan<- *result) {
	defer wg.Done()
	// 从jobChan中取出随机数计算各位数的和,将结果发送到resultChan
	for {
		job := <-prod
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		// newResult := &result{job, sum}
		newResult := &result{job: job, sum: sum}
		res <- newResult
	}

}
func main() {
	// myFunc()
	wg.Add(1)
	go producer(jobChan)
	wg.Add(24)
	// 开启24个goroutine执行cons
	for i := 0; i < 24; i++ {
		go cons(jobChan, resultChan)
	}
	// 主goroutine从resultChan取出结果并打印打终端输出
	for result := range resultChan {
		fmt.Printf("value:%d sum:%d\n", result.job.value, result.sum)
	}
	wg.Wait()
}
