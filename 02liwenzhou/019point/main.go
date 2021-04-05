package main

import "fmt"

// point 指针
// Go语言中不存在指针操作,只需要记住两个符号:
// 1. &:取地址
// 2. *:根据地址取值

func main() {
	// 1. &: 取地址
	// 2. *: 根据地址取值
	n := 18
	p := &n // 取地址
	fmt.Println(p)
	// fmt.Printf("%p\n", n)  // 取不出来
	fmt.Printf("%p\n", p) // 可以取出来	内存地址本身就是16进制数字
	fmt.Println(&p)       // 指向 p指针本身的指针
	fmt.Printf("%T\n", p) // 取不出来	*int int类型的指针

	m := *p // 取值
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	fmt.Printf("%p\n", &m) // 新的地址值
	// 对变量进行取地址(&)操作,可以获得这个变量的指针变量
	// 指针变量的值是指针地址
	// 对指针变量进行取值(*)操作,可以获得指针变量指向的原本变量的值.

	// var a *int // 定义一个int类型的指针   本质就是int类型的内存地址   默认值nil
	// fmt.Println(a)	// nil
	// 空的内存地址nil  *a  找不到 还要强行赋值 报空指针
	// *a = 100 // invalid memory address or nil pointer dereference
	//  fmt.Println(*a)

	var a1 *int     // 定义一个int类型的指针   本质就是int类型的内存地址   默认值nil
	fmt.Println(a1) // nil

	// new函数申请一个内存地址  用来给 string\int分配内存 *int  *string
	var a2 = new(int) // 开辟一个内存地址
	fmt.Println(a2)   //0xc00001e0c8 新地址
	fmt.Println(*a2)
	*a2 = 100
	fmt.Println(*a2)

	// make 也是用于内存分配的 区别于new 它只用于slice map 以及channel的内存创建 而且它返回的类型就是这三个类型本身 而不是他们的指针类型
	// 因为这三种类型就是引用类型,所以没必要返回他们的指针了.
	// var b map[string]int
	// b["沙河娜扎"] = 100		// 未初始化就调用assignment to entry in nil map
	// fmt.Println(b)
	var b1 = make(map[string]int)
	b1["沙河娜扎"] = 100 // 未初始化就调用assignment to entry in nil map
	fmt.Println(b1)

}
