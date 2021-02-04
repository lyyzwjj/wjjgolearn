package aka_function

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

/**
 * @author  wjj
 * @date  2020/8/29 12:39 上午
 * @description  闭包
 */

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 计算机程序的构造和解释
func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)
	c, _ := returnMultiValues()
	t.Log(c)
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(5))
}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

// 延迟函数 类似try catch finally
// defer 释放资源释放锁
func TestDefer(t *testing.T) {
	defer func() {
		fmt.Println("Clear resources.")
	}()
	fmt.Println("Start")
	panic("err")
}
