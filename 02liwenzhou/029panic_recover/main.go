package main

import "fmt"

// panic 和 recover

func funcA() {
	fmt.Println("a")
}

// recover()必须搭配defer使用
// defer一定要在可能引发panic的语句之前定义

func funcB() {
	// 刚刚打开数据库连接
	defer func() {
		err := recover() // 尽量不要使用
		fmt.Println(err)
		fmt.Println("释放数据库连接...") // panic之前执行
	}()
	panic("出现了严重的错误!!!") // 程序崩溃退出
	fmt.Println("b")
}

func funcC() {
	fmt.Println("c")
}
func main() {
	funcA()
	funcB()
	funcC()
}
