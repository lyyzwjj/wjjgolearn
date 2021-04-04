package main

import "fmt"

// 递归:函数自己调用自己
// 永远不要高估自己
// 斐波那契
// 阶乘
// 遍历文件夹
// 斐波那契
func Fibonacci(count uint64) uint64 {
	if count == 0 {
		return 0
	} else if count == 1 {
		return 1
	} else {
		return Fibonacci(count-1) + Fibonacci(count-2)
	}
}

// 走台阶steps
func UpSteps(step uint64) uint64 {
	if step == 1 {
		return 1 // 剩一个台阶 有1中走法
	} else if step == 2 {
		return 2 // 剩两个台阶 有2种走法
	} else {
		// 剩三个台阶 就等于 剩1个1阶的 和1个2阶楼梯 两种可能
		// 剩四个台阶 走一阶 剩3个台阶  或走两阶 剩2阶  把剩3阶的走法 可能的所有走法加起来  加上 把剩2阶的走法 可能的所有走法加起来
		// 每一阶都是有两组 子台阶 可能的走法组成
		return UpSteps(step-1) + UpSteps(step-2)
	}
}

// 阶乘
func Factorial(num uint64) uint64 {
	if num > 1 {
		return num * Factorial(num-1)
	} else {
		return num
	}
}
func main() {
	fmt.Println(Fibonacci(50))
	fmt.Println(Factorial(40))
}
