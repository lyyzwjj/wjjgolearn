package main

import "fmt"

// goto
// 使用switch语句可方便地对大量的值进行条件判断
// 一个变量和具体的值作比较

// goto + label实现跳出多层for循环
func main() {
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	return
	// 标签
breakTag:
	fmt.Println("结束for循环")
	unGoto()
}

// 跳出多层for循环
func unGoto() {
	var flag = false
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	breakDemo()
}

func breakDemo() {
BREAKDEMO:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break BREAKDEMO
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	fmt.Println("...")
	continueDemo()
}

func continueDemo() {
forloop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			// forloop2:
			if i == 2 && j == 2 {
				continue forloop1
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
}
