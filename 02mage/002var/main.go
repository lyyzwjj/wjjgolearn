package main

import "fmt"

// Go语言中推荐使用驼峰式命名
// var student_name string   下划线
// var studentName string	小驼峰
// var StudentName string	大驼峰
// 声明变量
// var name string
// var age int
// var isOk bool

// 批量声明
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func foo() (int, string) {
	return 10, "Wjj"
}

// go fmt main.go  会格式化文件
func main() {
	name = "理想"
	age = 16
	isOk = true
	// Go语言中非全局变量声明必须使用,不使用就编译不过去
	fmt.Print(isOk)             // Î在终端中输出要打印的内容
	fmt.Printf("name:%s", name) // %s:占位符 使用name这个变量的值去替换占位符
	fmt.Println(age)            // 打印完指定的内容之后会在后面加一个换行符
	// 声明变量同时赋值
	var s1 string = "Wjj"
	fmt.Println(s1)
	// 类型推导 (根据值判断该变量是什么类型)
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明 := 只能在函数中声明
	s3 := "哈哈哈"
	fmt.Println(s3)
	// 匿名变量(anonymous variable) 用_表示 多重赋值时,如果想要忽略某个值(lua中也叫做哑元变量)
	x, _ := foo()
	_, y := foo()
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	// 注意事项
	// 1. 函数外的每个语句都必须使用关键字开始 (var const func等)
	// 2. :=不能用在函数外
	// 3. _多用于占位,表示忽略值
	// s1 := "Zst" 同一个作用域({})不能重复声明同名的变量
}
