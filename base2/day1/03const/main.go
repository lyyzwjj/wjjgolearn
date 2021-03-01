package main

import "fmt"

// 常量
// 定义了常量之后不能修改
// 在程序运行期间不会改变的量
const pi = 3.1415926

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时.如果某一行声明后没有赋值,默认和上一行一致
const (
	n1 = 200
	n2
	n3
)

// iota 是go语言的常量计数器,只能在常量的表达式中使用
// iota 在const关键字出现时将被重置为0.  在不同的const() 中才会被重置
// iota const中每新增一行常量声明将是iota计数一次(iota可理解为const语句块中的行索引)
// 使用iota能简化定义,在定义枚举时很有用

// iota:类似枚举
const (
	a1 = iota // 0
	a2        // 1
	a3        // 2
	a4        // 3
)

// 几个常见的iota例子
const (
	b1 = iota // 0
	b2        // 1
	_         // 2
	b3        // 3
)

// 插队
const (
	c1 = iota // 0
	c2 = 100  // 100
	c3        // 100
	c4 = iota // 3   // const中每新增一行常量声明将是iota计数一次
	c5        // 5
)

// 多个常量声明在同一行
const (
	d1, d2 = iota + 1, iota + 2 // 1	2    // const中每新增一行常量声明将是iota计数一次
	d3, d4 = iota + 1, iota + 2 // 2	3
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota) // 1<< 10
	MB = 1 << (10 * iota) // 1<< 20
	GB = 1 << (10 * iota) // 1<< 30
	TB = 1 << (10 * iota) // 1<< 40
	PB = 1 << (10 * iota) // 1<< 50
)

func main() {
	// pi = 123  不能再次赋值
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)

	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("a4:", a4)

	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)

	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)
	fmt.Println("c5:", c5)

	fmt.Println("d1:", d1)
	fmt.Println("d2:", d2)
	fmt.Println("d3:", d3)
	fmt.Println("d4:", d4)

	fmt.Println("KB:", KB)
	fmt.Println("MB:", MB)
	fmt.Println("GB:", GB)
	fmt.Println("TB:", TB)
	fmt.Println("PB:", PB)
}
