package main

import "fmt"

// 切片 slice
func main() {
	// 定义
	var s1 []int    //	定义一个存放int类型元素的切片
	var s2 []string // 	定义一个存放string类型元素的切片
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 初始化
	s1 = []int{1, 2, 3}
	s2 = []string{"沙河", "张江", "平山村"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)
	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // [1 3 5 7]  // 基于一个数组切割, 左包含右不包含, 左闭右开
	fmt.Println(s3)
	s4 := a1[1:6]
	fmt.Println(s4)
	s5 := a1[1:] // => [1:len(a1]
	s6 := a1[:4] // => [0:4]
	s7 := a1[:]  // => [0:len(a1]
	fmt.Println(s5, s6, s7)
	// 切片的容量是指底层数组的容量
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6)) // 4  7
	// cap 底层数组从切片的第一个元素到最后元素数量
	fmt.Printf("len(s5):%d cap(s5):%d\n", len(s5), cap(s5)) // 6  6

	// 切片指向了一个底层的数组
	// 切片的长度就是它元素的个数.
	// 切片的容量是底层数组从切片的第一个元素到最后一个元素的数量.

	// 切片在切片
	s8 := s4[3:] // s4	[3 5 7 9 11]  len 5  cap 6
	fmt.Println(s8)
	fmt.Printf("len(s8):%d cap(s8):%d\n", len(s8), cap(s8)) // 从新的s4 上面去数cap位数
	// 切片是引用类型,都指向了底层的一个数组
	fmt.Println(s4)
	a1[5] = 1300
	fmt.Println(s4)
	fmt.Println(s8)

	// 使用make()函数构造切片

	var BFS_02 = [][]interface{}{
		{0, 1}, {0, 4},
		{1, 2},
		{2, 0}, {2, 4}, {2, 5},
		{3, 1},
		{4, 6}, {4, 7},
		{5, 3}, {5, 7},
		{6, 2}, {6, 7},
	}

	for _, data := range BFS_02 {
		fmt.Println(data)
		for _, d := range data {
			fmt.Println(d)
		}
	}
}
