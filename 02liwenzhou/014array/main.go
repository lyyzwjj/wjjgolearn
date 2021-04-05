package main

import "fmt"

// 数组
// 存放元素的容器
// 必须指定存放的元素的类型和容量(长度)
// **数组的长度是数组类型的一部分**
func main() {
	// 定义
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]
	// a1 == a2  不能比较 完全不同的类型
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组的初始化
	// 如果不初始化: 默认元素都是零值(布尔值:false, 整型和浮点型都是0, 字符串: "")
	fmt.Println(a1, a2)
	// 1. 初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2. 初始化方式2
	// a10 := [8]int{0, 1, 2, 3, 4, 5, 6, 7}
	// ... 根据初始值自动推断数组长度是多少
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	fmt.Println(a10)
	// 3. 初始化方式3
	a3 := [5]int{1, 2} // 其他位数为默认值
	fmt.Println(a3)
	// 4. 初始化方式4 根据索引来初始化
	a4 := [5]int{0: 1, 4: 2}
	fmt.Println(a4)

	// 数组的遍历
	cities := [...]string{"北京", "上海", "广州", "深圳"}
	// 1. 根据索引遍历
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	// 2. for range遍历
	for i, v := range cities {
		fmt.Println(i, v)
	}

	// 多维数组
	// [[1 2] [3 4] [5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(a11)

	// 多维数组遍历
	/*	for i := 0; i < len(a11); i++ {
		for j := 0; j < len(a11[i]); j++ {
			fmt.Println(a11[i][j])
		}
	}*/
	for _, v1 := range a11 {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

	// 数组是值类型
	b1 := [3]int{1, 2, 3} // [1 2 3]
	b2 := b1              // [1 2 3]	Ctrl+C Ctrl+V 把word文档从文件夹A拷贝到文件夹B
	b2[0] = 100           // b2: [100 2 3]
	fmt.Println(b1, b2)   // b1 ?
	// 数组支持 == 和 != 操作符 因为内存总是被初始化过的
	// [n]*T表示指针数组 *[n]T 表示数组指针
}
