package main

import "fmt"

// struct	值类型 指针类型的结构体
// go 语言中 函数传参永远传的是拷贝 副本

type Person struct {
	name, gender string
}

// 值类型传递  永远都是拷贝
func f(x Person) {
	x.gender = "女" // 修改的是副本的gender
}

func f1(x *Person) {
	// (*x).gender = "女"	// 根据内存地址找到那个原变量,修改的就是原来的变量
	// 语法糖 等价于上面 自动根据资政找到对应的变量
	x.gender = "女" // 修改的是副本的gender
}
func main() {
	var p Person
	p.name = "周林"
	p.gender = "男"
	f(p)
	fmt.Println(p)
	f1(&p) // ox1234ac23  拷贝这个地址传参
	fmt.Println(p)
	// 结构体指针1
	var p2 = new(Person)      // 开辟一个Person的内存地址 并用p2 指向   p2就是一个指针
	fmt.Println("p2= ", p2)   // &{ }			// 指针类型打印 &加具体内容{}
	fmt.Println("*p2= ", *p2) //  { }			// 指针类容取值 就取的是具体的内容{}
	fmt.Println("&p2= ", &p2) // 0xc0000ae020	// 指向p2指针的指针
	fmt.Printf("%T\n", p2)    // *main.Person		指针类型
	fmt.Printf("%T\n", &p2)   // **main.Person		 &p2 指针类型的指针 的类型
	fmt.Printf("%p\n", *p2)   // %!p(main.Person={ })	// 不是指针类型强行用%p
	fmt.Printf("%p\n", p2)    // 0xc0000a6060	// 指针类型取 取的就是他自己的内存地址
	fmt.Printf("%p\n", &p2)   // 0xc0000ae020	// 指针类型取 取的就是他自己的内存地址 就是二次指针的地址 就等于&p
	p2.name = "理想"
	fmt.Printf("%x\n", p2) // &{e79086e683b3 }
	// 结构体指针
	// 2.1 key-value初始化	可以不包括所有值  最常用的
	var p3 = &Person{
		name:   "元帅",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3) // %# 八进制数前面加o(%#o),十六进制数前面加ox(%#x)或0X(%#X),指针去掉前面的ox(%#p) 对%q(%#q) %U(%#U)会输出空格和单引号括起来的go字面值;
	// 2.2 使用值列表的形式初始化,值得顺序要和结构体定义时字段的顺序一致
	p4 := &Person{"小公举", "女"} // 一定要所有字段
	fmt.Printf("%#v\n", p4)
}
