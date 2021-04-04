package main

import "fmt"

// 使用值接收者和指针接受者的区别?
// 使用值接收者实现接口, 结构体类型和结构体指针类型的变量都能存
// 指针接收者实现接口只能存结构体指针
type Animal interface {
	move()
	eat(string)
}

type Cat struct {
	name string
	feet int8
}
type Cat1 struct {
	name string
	feet int8
}

// 使用值接受者实现了接口的所有方法 方法使用值接受者
func (c Cat) move() {
	fmt.Printf("走猫步...\n")
}

func (c Cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

// 使用指针接收者
func (c *Cat1) move() {
	fmt.Printf("走猫步...\n")
}

func (c *Cat1) eat(food string) {
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

	// 指针接收者实现接口只能存结构体指针
	g1 := Cat1{"tom", 4}  //cat1
	g2 := &Cat1{"假老练", 4} // *Cat1
	a1 = &g1              // 实现animal这个接口的是Cat的指针
	fmt.Println(g1)
	a1 = g2 // c2 也能赋值给a1
	fmt.Println(g2)

}
