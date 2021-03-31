package main

import "fmt"

// 浮点数
func main() {
	// math.MaxFloat32 // float32最大值
	f1 := 1.23456
	fmt.Printf("%T\n", f1) // 默认Go语言中的小数都是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显示声明类型float32类型
	// f1 = f2 					// float32类型的值不能直接赋值给float64类型的变量
	var c1 complex64 // complex64 实部和虚部为32位		64位
	c1 = 1 + 2i
	var c2 complex128 // complex128 实部和虚部为64位	128位
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}
