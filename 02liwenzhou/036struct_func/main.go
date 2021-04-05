package main

import "fmt"

// 构造函数

type Person struct {
	name string
	age  int
}

type Dog struct {
	name string
}

// 构造函数: 约定俗称用new开头
// 返回的是结构体还是结构体指针

// struct 字段少 结构体 小 拷贝开销少

func newPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

// struct 字段4 5 个往上返回结构体指针
// 当结构体比较大的时候 尽量使用结构体指针,减少程序运行的内存开销
func newPerson1(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func newDog(name string) Dog {
	return Dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("元帅", 18)
	p2 := newPerson("周林", 900)
	fmt.Println(p1, p2)
	d1 := newDog("周林")
	fmt.Println(d1)
}
