package main

import (
	"fmt"
	"strings"
)

// 普通函数只需要看 函数体内部的内容
// 闭包 一个函数除了可以引用内部的代码  它还可以引用它外部的的变量 adder的x 可以被adder 内部的匿名函数使用
// 闭包是一个函数,这个函数包含了他外部作用域的一个变量
// 底层的原理:
// 1. 函数可以作为返回值
// 2. 函数内部查找变量的顺序,先在自己内部找,找不到往外层找

func f1(f func()) {
	fmt.Println("this is f1")
	f()

}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)

}

// 定义一个函数对f2进行包装
func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

// 要求f1(f2)
func f3(f func(int, int), x, y int) func() {
	tmp := func() {
		// fmt.Println(x + y)
		f(x, y)
	}
	return tmp
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			name += suffix
		}
		return name
	}

}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	lixiang(100)
	lixiang1(f2, 100, 200)
	ret := lixiang2(f2, 100, 200)
	ret()
	ret1 := adder(100) // 指向adder 内返回的函数 还额外包含了 adder的入参 x
	ret2 := ret1(200)
	fmt.Println(ret2)
	ret3 := f3(f2, 1, 2) // 把原本需要传递两个int类型的参数包装成一个不需要传参的函数
	fmt.Printf("%T\n", ret3)
	ret3() // 可以直接执行
	f1(ret3)
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")
	fmt.Println(jpgFunc("test"))
	fmt.Println(jpgFunc("test.jpg"))
	fmt.Println(txtFunc("test"))

	f1, f2 := calc(10)        // base = 10	// base 函数包含了他外部作用域的一个变量  这个base在外部作用域中一直在改变
	fmt.Println(f1(1), f2(2)) // 11 base 11 9 base 9
	fmt.Println(f1(3), f2(4)) // 12 base 12 9 base 8
	fmt.Println(f1(5), f2(6)) // 13 base 13 7 base 7
}

func lixiang(x int) {
	tmp := func() {
		fmt.Println(x) // 现在自己这里找 在往上找
	}
	tmp()
}

func lixiang1(x func(int, int), m, n int) {
	tmp := func() {
		x(m, n)
	}
	tmp()
}

func lixiang2(x func(int, int), m, n int) func() {
	tmp := func() {
		x(m, n)
	}
	return tmp
}
