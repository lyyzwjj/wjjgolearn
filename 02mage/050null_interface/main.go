package main

import "fmt"

// 类型断言
// 	str, ok := a.(string)
//  a.(type)

func assign(a interface{}) {
	// fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串:", str)
	}
}

// 类型断言
func assign2(a interface{}) {
	fmt.Printf("%T\n", a)
	switch v := a.(type) {
	case string:
		fmt.Printf("是一个字符串:%v", v)
	case int:
		fmt.Printf("是一个int:%v", v)
	case int64:
		fmt.Printf("是一个int64:%v", v)
	case bool:
		fmt.Printf("是一个bool:%v", v)
	}
}

func show(a interface{}) {
	fmt.Printf("type:%T value:%#v\n", a, a)

}
func main() {
	assign(88)
	assign("88")
	assign2(true)
	assign2("hahaha")
	assign2(int64(200))
	assign2(64)
}
