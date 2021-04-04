package main

import "fmt"

// 结构体嵌套

type Address struct {
	province string
	city     string
}

type WorkPlace struct {
	province string
	city     string
}
type Person struct {
	name string
	age  int
	//addr Address	// 此方式只能p1.addr.city 取到
	Address // 匿名嵌套 此方式p1.city 直接取到
	WorkPlace
}

type Company struct {
	name string
	addr Address
}

func main() {
	p1 := Person{
		name: "王吉吉",
		age:  26,
		/*		addr: Address{
				province: "湖南",
				city:     "长沙",
			},*/
		Address: Address{
			province: "湖南",
			city:     "长沙",
		},
	}
	fmt.Println(p1)
	//	fmt.Println(p1.name, p1.addr.city)
	// fmt.Println(p1.city) // 先在自己结构体找这个字段,找不到就去匿名嵌套的结构体中查找该字段
	// 如果匿名结构体中 不同的匿名结构体中有相同的字段 就不能直接取了
	fmt.Println(p1.Address.city)
	fmt.Println(p1.WorkPlace.city)
	fmt.Printf("p1 = %#v\n", p1)
}
