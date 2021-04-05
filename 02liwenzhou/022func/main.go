package main

import "fmt"

// 函数
// 函数的定义
// 函数存在的意义
// 函数是一段代码的封装
// 把一段逻辑抽象出来封装到一个函数中,给它起个名字.每次用到它的时候直接用函数名调用就行了
// 使用函数能够让代码结构更清晰\更简洁.

// 没有返回值
func f1(x int, y int) {
	fmt.Println(x + y)
}

// 没有参数没有返回值
func f2() {
	fmt.Println("f2")
}

// 没有参数但有返回值的
func f3() int {
	ret := 3
	return ret
}

// 返回值可以命名也可以不命名
// 命名返回值
func sum(x int, y int) int {
	return x + y // 没有声明返回值 必须return x+y
}

// 命名返回值 就相当于提前命名了一个变量 相当于在函数中声明了一个变量
func sum1(x int, y int) (ret int) { // ret int 是声明过程
	ret = x + y // 直接赋值
	return      // 使用命名返回值可以return 后面省略 返回值可以不写ret
}

// 多个返回值
func f5() (int, string) {
	return 1, "沙河"
}

//	参数的类型简写:
//	当参数中连续两个参数的类型一致时,我们可以将前面那个参数类型省略
/*func f6(x int, y int) int {
	return x + y
}*/

func f6(x, y int) int {
	return x + y
}

func f61(x, y int, m, n string, z, j bool) int {
	return x + y
}

// 可变长参数  必须放在函数参数的最后
func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y) // y 是一个slice 切片 []int
}

// Go语言中函数没有默认参数这个概念

func main() {
	r := sum(1, 2)
	fmt.Println(r)
	_, n := f5()
	fmt.Println(n)
	f7("下雨了")
	f7("下雨了", 1, 2, 3, 4, 5)
}
