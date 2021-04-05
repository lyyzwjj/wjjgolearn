package main

import "fmt"

// 接口示例2
// 不管是什么牌子的车,都能跑
type Animal interface { // 要实现了所有接口中的方法才能任务是接口的类型
	move()
	eat(string)
}

type Cat struct {
	name string
	feet int8
}

func (c Cat) move() {
	fmt.Printf("走猫步...\n")
}

/*func (c Cat) eat() {
	fmt.Printf("猫吃鱼...")
}*/
func (c Cat) eat(food string) {
	fmt.Printf("猫吃%s...\n", food)
}

type Chicken struct {
	feet int8
}

func (c Chicken) move() {
	fmt.Printf("鸡动!\n")
}
func (c Chicken) eat() {
	fmt.Printf("吃鸡饲料!\n")
}

/*func (c Chicken) eat(food string) {
	fmt.Printf("吃鸡饲料!")
}*/

func main() {
	var a1 Animal
	fmt.Printf("%T\n", a1) // nil
	bc := Cat{
		name: "淘气",
		feet: 4,
	}
	a1 = bc
	fmt.Println(a1)
	fmt.Printf("%T\n", a1) // main.Cat
	a1.eat("小黄鱼")
	c := Chicken{
		feet: 8,
	}
	fmt.Print(c)
	// a1 = c  // 不能讲c赋值给a1 应为c 没有实现a1所有的方法

}
