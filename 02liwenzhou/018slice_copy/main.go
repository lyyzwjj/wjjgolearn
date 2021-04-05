package main

import (
	"fmt"
	"sort"
)

// copy

func main() {
	a1 := []int{1, 3, 5}
	a2 := a1                   // 赋值	浅拷贝 拷贝指向的地址值
	var a3 = make([]int, 0, 3) // make创建切片
	copy(a3, a1)               // copy	将切片a1中的元素复制到切片a3
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)
	fmt.Printf("p(a1)=%p p(a2)=%p p(a3)=%p\n", a1, a2, a3) // a1 a2 相等 a3 不同

	// 将a1中的索引为1的3这元素删掉
	fmt.Println(a1[:1])              // 虽然a1[:1] 只去第一个数a1[0]	即是 [100]	但是其底层数组还是 [100, 3, 5]
	fmt.Println(a1[0], a1[1], a1[2]) // 还是能访问 a1[0] a1[1] a1[2]的
	a1 = append(a1[:1], a1[2:]...)
	fmt.Println(a1)
	fmt.Println(cap(a1))    // 还是当初创建时候的3
	x1 := [...]int{1, 3, 5} // 数组 有... 或数字的
	s1 := x1[:]             // 切片
	// 1. 切片不保存具体的值
	// 2. 切片对应一个底层数组
	// 3. 底层数组都是一块连续的内存
	fmt.Printf("%p\n", &s1[0])
	fmt.Println(s1, len(s1), cap(s1))
	// s1[:1] 其实底层数组还是[1,3,5] 只是len = 1 只能看到1位	但是cap还是3
	s1 = append(s1[:1], s1[2:]...) // 修改了底层数组!
	fmt.Printf("%p\n", &s1[0])
	fmt.Println(s1, len(s1), cap(s1))
	fmt.Println(x1) // 1 5 5  s1 其实底层指向的切片是x1
	s1[0] = 100
	fmt.Println(x1) // 100 5 5
	sliceTest()
}

func sliceTest() {
	var a = make([]int, 5, 10) // len 为 5 cap 为10
	fmt.Println(&a[3])
	fmt.Println(&a[4])
	// fmt.Println(&a[5])
	// fmt.Println(a[5]) // 报错 a[5] 为nil
	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)

	// 初始值
	// [0 0 0 0 0]
	// [0 0 0 0 0 nil nil nil nil nil]
	// [0 0 0 0 0 0 ]
	// [0 0 0 0 0 0 1]
	// [0 0 0 0 0 0 1 2]
	// [0 0 0 0 0 0 1 2 3]
	// ...
	// [0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]

	// 排序
	// 使用sort 包对数组排序
	var l = [...]int{2, 5, 7, 4}
	sort.Ints(l[:]) // 参数是切片 对切片进行排序
	fmt.Println(l)

	a1 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	s1 := a1[:]
	// 删掉索引为1的那个3
	s1 = append(s1[0:1], s1[2:]...)
	// s1[0:1] 底层 1, //后面3, 5, 7, 9, 11, 13, 15, 17都当看不见 其实实际还有 对应a 1 3, 5, 7, 9, 11, 13, 15, 17
	// 1, 5, 7, 9, 11, 13, 15, 17,17
	fmt.Println(s1)
	fmt.Println(a1)

}
