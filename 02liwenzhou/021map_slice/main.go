package main

import "fmt"

// map和slice组合
// slice 和 map 都要做初始化

func main() {
	// 元素类型为map的切片
	//	var s1 = make([]map[int]string, 0, 10)	// 0个元素后面会空指针
	//	s1[0][100] = "A"
	//	fmt.Println(s1)
	var s1 = make([]map[int]string, 10, 10) // map的零值 map[]
	// 没有对内部的map做初始化
	// s1[0][100] = "A"
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "沙河"
	fmt.Println(s1)

	// 值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}
	fmt.Println(m1)
}
