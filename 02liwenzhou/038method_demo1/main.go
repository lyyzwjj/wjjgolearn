package main

import "fmt"

// 给自定义类型加方法
// 给内置类型添加方法
// 不能给别的包里面的类型添加方法,只能给自己包里的类型添加方法
// 或者只用 type 自定义类型

type myInt int

func (i myInt) hello() {
	fmt.Println("我是一个int")
}

func main() {
	m := myInt(100) // int8(100)  强制转换
	m.hello()

	var x int32 = 10 // int32
	fmt.Println(x)
	var x1 = 10 // 默认的int32
	fmt.Println(x1)
	var x2 = int32(10) // 强制类型转换
	fmt.Println(x2)
	x3 := int32(10)
	fmt.Println(x3)

	// 声明一个myInt类型的变量m,它的值是100
	var c myInt
	c = 100
	fmt.Println(c)
	var c11 myInt = 100
	fmt.Println(c11)
	var c1 = 100
	fmt.Println(c1)
	var c2 = myInt(100) // 强制类型转换
	fmt.Println(c2)
	c3 := myInt(10)
	fmt.Println(c3)

	// 方法1
	var p Person // 声明一个Person类型的变量p
	p.name = "元帅"
	p.age = 18
	fmt.Println(p)

	var p1 Person
	p1.name = "周林"
	p1.age = 9000
	fmt.Println(p1)
	// 方法2
	s1 := []int{1, 2, 3, 4}
	m1 := map[string]int{
		"stu1": 100,
		"stu2": 99,
		"stu3": 0,
	}
	fmt.Println(s1, m1)
	// 键值初始化  声明同时初始化
	var p2 = Person{
		name: "冠华",
		age:  15,
	}
	fmt.Println(p2)

	// 方法3
	// 值列表初始化
	var p3 = Person{
		"理想", 100,
	}
	fmt.Println(p3)

	// 构造函数
	person := newPerson("老李", 25)
	fmt.Println(person)
}

// 结构体初始化
type Person struct {
	name string
	age  int
}

// 问: 为什么要有构造函数
func newPerson(name string, age int) *Person {
	// 别人调用我,我能给她一个Person类型的变量p
	return &Person{
		name: name,
		age:  age,
	}

}
