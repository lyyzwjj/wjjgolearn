package _86split_string

import (
	"fmt"
	"strings"
)

// Split 切割字符串
// example:
// abc, b=> [a c]

func Split(str string, sep string) []string {
	// str:"babcbef" sep="b" [a cdef]
	// var ret []string
	// 优化  直接算好最终的[]string 的长度
	var ret = make([]string, 0, strings.Count(str, sep)+1)
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[0:index]) //append会帮助初始化 此处申请分配内存 避免运行时多次分配内存
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	// 仅为测试覆盖率
	if index == -5 {
		fmt.Println("真无聊")
	}
	// time.Sleep(time.Second) // 基准测试执行1s  睡1s刚好只执行一次
	ret = append(ret, str)
	return ret
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}
