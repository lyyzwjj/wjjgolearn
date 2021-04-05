package calc

import "fmt"

var x = 100

const pi = 3.14

// init 方法不能被手动调用
// 在函数运行时自自动被调用执行
func init() {
	fmt.Println("import 我自动执行...")
	fmt.Println(x, pi)
}

// 包中的标识符(变量名\函数名\结构体\接口等)如果首字母是小写的,表示私有(只能在当前这个包中使用)
func Add(x, y int) int {
	return x + y
}
