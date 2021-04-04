package main

import "fmt"

// 接口
// 接口是一种类型,是一种特殊的类型,它规定了变量有哪些方法
// 在编程中会遇到一下场景:
// 我们不关心一个变量是什么类型,我只关心能调用它的什么方法
// 定义
// type 接口名 interface {
// 		方法名1(参数1,参数2,...)(返回值1,返回值2,...)
// 		方法名2(参数1,参数2,...)(返回值1,返回值2,...)
//		...
// }

// 用来给变量\参数\返回值等设置类型

// 接口实现
// 一个变量如果实现了接口中规定的所有的方法,那么这个变量就实现了这个接口,可以称为这个接口类型的变量.

// 接口保存的分为两部分:值的类型和值本身	动态类型 和动态值  这样就实现了接口变量能够存储不同的值.	最开始的都是nil 没有指定任何类型的时候

// 定义一个能叫的类型
type Speaker interface {
	speak() // 只要实现了speak方法的变量都是speaker类型	方法签名
}
type Cat struct {
}
type Dog struct {
}
type Person struct {
}

func (c Cat) speak() {
	fmt.Println("喵喵喵~")
}
func (d Dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p Person) speak() {
	fmt.Println("啊啊啊~")
}

func da(s Speaker) {
	// 接收一个参数,传进来什么,我就打什么
	s.speak()
}

func main() {
	var c1 Cat
	var d1 Dog
	var p1 Person
	da(c1)
	da(d1)
	da(p1)

	var ss Speaker // 定义一个接口类型:Speaker 的变量:ss
	ss = c1
	ss = d1
	ss = p1
	fmt.Println(ss)
}
