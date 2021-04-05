package main

import "fmt"

// 模拟继承

// 动物类
type Animal struct {
	name string
}

// 狗类
type Dog struct {
	feet   uint8
	Animal // 继承了Animal的字段以及方法 Animal拥有的方法,dog此时也有了
}

// 给狗实现一个汪汪汪的方法

func (d Dog) bark() {
	fmt.Printf("%s在叫:汪汪汪~\n", d.name)
}

func (a Animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

func main() {
	d1 := Dog{
		Animal: Animal{name: "周林"},
		feet:   4,
	}
	fmt.Println(d1)
	d1.bark()
	d1.move()
}
