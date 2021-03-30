package main

import "fmt"

// if条件判断

func main() {
	age := 19
	if age > 18 { // 如果age > 18 就执行这个{}中的代码
		fmt.Println("澳门首家线上赌场开业啦!")
	} else { // 否则就执行这个{}中的代码
		fmt.Println("该写暑假作业拉!")
	}
	// 多个条件判断
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("好好学习!")
	}
	// 作用域
	// age变量此时只在if条件判断语句中生效
	if age1 := 19; age1 > 18 { // 如果age > 18 就执行这个{}中的代码
		fmt.Println("澳门首家线上赌场开业啦!")
	} else { // 否则就执行这个{}中的代码
		fmt.Println("该写暑假作业拉!")
	}
	// fmt.Println(age1)  在这里是找不到age1的 报错  不仅仅是书写合并成了一行 连age1的作用域也窄了
}
