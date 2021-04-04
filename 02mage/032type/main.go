package main

import "fmt"

// 自定义类型和类型别名
// type后面跟的是类型

type myInt int // 自定义类型

type yourInt = int // 类型别名

func main() {
	var n myInt
	var m yourInt
	n = 100
	m = 100
	fmt.Println(n)
	fmt.Printf("%T\n", n) // main.myInt
	fmt.Println(m)
	fmt.Printf("%T\n", m) // int

	var c rune // rune 是int32类型别名
	c = '中'
	fmt.Println(c)
	fmt.Printf("%T\n", c) // int32
}
