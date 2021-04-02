package main

import "fmt"

// defer
// 释放资源 打开socket连接 打开数据库连接
// 一个函数中可以有多个defer语句
// defer 相当于栈 后进先出 越后面定义的defer 最后越先执行

func deferDemo() {
	fmt.Println("start")
	defer fmt.Println("嘿嘿嘿") // defer把他后面的语句延迟到函数即将返回的时候再执行
	defer fmt.Println("呵呵呵")
	defer fmt.Println("哈哈哈") // 倒过来执行 先执行 哈哈哈  最后执行嘿嘿嘿
	fmt.Println("end")
}

func main() {
	deferDemo()
}
