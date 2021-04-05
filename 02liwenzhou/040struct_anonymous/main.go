package main

import "fmt"

// 匿名字段	相同类型只能写一个
// 字段比较少也比较简单的场景
// 不常用!!

type Person struct {
	string
	int
}

func main() {
	p1 := Person{
		"周林", 26,
	}
	fmt.Println(p1)
	fmt.Println(p1.string)
	fmt.Println(p1.int)
}
