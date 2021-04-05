package main

import "fmt"

// struct	值类型

type Person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	// 声明一个Person类型的变量p
	var p Person
	// 通过字段赋值
	p.name = "王吉吉"
	p.age = 26
	p.gender = "周林"
	p.name = "M"
	p.hobby = []string{"IT", "japanese", "movie"}
	fmt.Println(p)

	var p2 Person
	// 通过字段赋值
	p2.name = "理想"
	p2.age = 18
	fmt.Printf("type:%T value:%v\n", p2, p2)

	// 匿名结构体: 多用于一些临时场景
	var s struct { // 直接声明一个 struct类型的s变量
		x string
		y int
	}
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n", s, s)
}
