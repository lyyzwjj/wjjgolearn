package main

import "fmt"

// 同一个结构体可以实现多个接口
//
type Mover interface {
	move()
}

type Eater interface {
	eat(string)
}

// 接口嵌套
type Animal interface {
	Mover
	Eater
}

type Cat struct {
	name string
	feet int8
}

// Cat实现了Mover接口
func (c Cat) move() {
	fmt.Printf("走猫步...\n")
}

// Cat实现了Eater接口
func (c Cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

func main() {
	var a1 Animal
	// 使用值接收者实现接口, 结构体类型和结构体指针类型的变量都能存
	c1 := Cat{"tom", 4}  //cat
	c2 := &Cat{"假老练", 4} // *Cat
	a1 = c1
	fmt.Println(a1)
	a1 = c2 // c2 也能赋值给a1
	fmt.Println(a1)
}
