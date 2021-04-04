package main

import "fmt"

// struct 结构体占用一块连续的内存空间

type x struct {
	a int8 // 8bit => 1byte
	b int8
	c string
}

func main() {
	m := x{
		a: int8(10),
		b: int8(20),
		c: "嘿嘿",
	}
	fmt.Printf("%p\n", &(m.a))
	fmt.Printf("%p\n", &(m.b))
	fmt.Printf("%p\n", &(m.c))

	// 0xc00000c040	内存地址的结尾
	// 0xc00000c041	一个内存地址 +1  就是8bit 2*8 = 16 16进制1
	// 0xc00000c048 41-48结束 7个字节 6*8bit   TODO ??
	// 在 Go 中恰到好处的内存对齐
	// https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
}
