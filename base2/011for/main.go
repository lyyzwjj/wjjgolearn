package main

import "fmt"

// for循环

func main() {
	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// 变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	// 变种2
	var j = 5
	for j < 10 {
		fmt.Println(j)
		j++
	}
	// 死循环 无限循环
	/*	for {
			fmt.Println("123")
	}*/
	// for range 键值循环
	// 遍历数组 切片 字符串 map 及 通道 channel
	// 数组 切片 字符串 返回索引和值
	// map 返回键和值
	// 通道(channel)只返回通道内的值

	s := "Hello沙河"
	for i, v := range s {
		fmt.Printf("%d %c\n", i, v)
		// 值得注意的是	沙要占3个字节
		/*
			0 H
			1 e
			2 l
			3 l
			4 o
			5 沙
			8 河
		*/
	}
	// 九九乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
	// 当i=5时, 就跳出for循环
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over")

	// 当i=5时, 就跳过此次for循环 (不执行for循环内部的打印语句)
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
}
