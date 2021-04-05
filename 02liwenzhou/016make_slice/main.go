package main

import "fmt"

// make()函数创造切片

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
	// fmt.Println(s1[6]) 报错
	s2 := make([]int, 5)
	fmt.Printf("s2=%v len(s2)=%d cap(s2)=%d\n", s2, len(s2), cap(s2))
	// 切片就是一个框 框住了一块连续的内存. 只能存储相同类型的数据 属于引用类型 真正的数据都是保存在底层数组里的
	// 切片不能直接比较
	// 空切片 只定义了的切片 长度和容量都是0 一个nil值的切片并没有底层数组	但是我们不能说一个长度和容量都是0的切片一定是nil
	var r1 []int
	r2 := []int{}
	r3 := make([]int, 0)
	fmt.Println(r1, r2, r3)
	fmt.Printf("len(r1)=%d cap(r1)=%d s1==nil ? %v\n", len(r1), cap(r1), s1 == nil)

	// 数组的赋值		值类型
	b1 := [3]int{1, 3, 5}
	b2 := b1
	fmt.Println(b1, b2)
	b2[0] = 100
	fmt.Println(b1, b2)

	// 切片的赋值		引用类型
	s3 := []int{1, 3, 5}
	s4 := s3 // s3和s4都指向了同一个底层数组
	fmt.Println(s3, s4)
	s4[0] = 1000
	fmt.Println(s3, s4)

	// 切片的遍历
	// 1. 索引遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	// 2. for range循环
	for i, v := range s3 {
		fmt.Println(i, v)
	}
}
