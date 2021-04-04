package main

import "fmt"

// 方法

// 标识符:变量名 函数名 类型名 方法名
// Go语言中如果标识符首字母是大写的,就表示对外部包不可见(暴露的,公有的).

// Dog 这是一个狗的结构体
type Dog struct {
	name string
}

type Person struct {
	name string
	age  int
}

// 构造函数
func newDog(name string) Dog {
	return Dog{
		name: name,
	}
}

// 方法是作用于特定类型的函数
// 接受者表示的是调用该方法的具体类型变量,多用类型名首字母小写表示

// 方法接收者
// func(接受者变量 接受者类型) 方法名(参数列表) (返回参数){
// 函数体
// }
func (d Dog) bark() {
	fmt.Printf("%s:汪汪汪\n", d.name)
}

// 使用值接受者
func (p Person) newYear() { // 值传递
	p.age++
}

// 指针接受者:传内存地址进去
func (p *Person) newYearReal() { // 值传递
	// (*p).age++
	p.age++
}

func (p *Person) Dream() { // 值传递
	fmt.Println("不上班也能挣钱!")
}

func newPerson(name string, age int) Person {
	return Person{
		name: name,
		age:  age,
	}
}

// 什么时候应该使用指针接收者?

// 1. 需要修改接受者中的值
// 2. 接收者是拷贝代价比较大的大对象
// 3. 保持一致性,如果有某个方法使用了指针接受者,那么其他的方法也应该使用指针接收者
func main() {
	d1 := newDog("zhoulin")
	d1.bark()
	p1 := newPerson("元帅", 18) // 18
	p1.newYear()              // 值传递 将p1 拷贝一份进方法 age++ 是副本值++
	fmt.Println(p1.age)       // 18
	p1.newYearReal()
	fmt.Println(p1.age) // 19
	p1.Dream()
}
