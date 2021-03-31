package main

import "fmt"

// 整形
func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // 把十进制数转化成二进制
	fmt.Printf("%o\n", i1) // 把十进制数转化成八进制
	fmt.Printf("%x\n", i1) // 把十进制数转化成十六进制
	// 八进制	涉及到给文件设置权限的时候
	i2 := 077
	fmt.Printf("%d\n", i2)
	// 十六进制	涉及到内存地址的时候
	i3 := 0x1234567
	fmt.Printf("%x\n", i3)
	// 查看变量的类型
	fmt.Printf("%T\n", i3)
	// 声明int8类型的变量
	i4 := int8(9) // 明确指定int8类型，否则就是默认为int类型
	fmt.Printf("%T\n", i4)
}
