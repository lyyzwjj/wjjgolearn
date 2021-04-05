package main

import "fmt"

// Go语言中函数的return 不是原子操作,在底层是分为两部来执行
// 第一步: 返回值赋值
// defer
// 第二步: 真正的RET返回

// 函数中如果存在defer,那么defer执行的时机是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

// 匿名返回值 = x = 5
// x++
// 所以返回值还是5

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

// 返回值 x = 5
// x++
// 所以返回值还是6

func f3() (y int) {
	x := 5
	defer func() {
		x++ // x里面找不到就找外面的x
	}()
	return x
}

// 返回值 y = x
// x++
// 所以返回值还是5  简单类型  是拷贝

func f4() (x int) {
	defer func(x int) {
		x++ // 改变的是函数中的副本
	}(x)
	return 5
}

func f5() (x int) {
	defer func(x int) int {
		x++ // 改变的是函数中的副本
		return x
	}(x)
	return 5
}

// 返回值 y = x
// x++
// 所以返回值还是5  简单类型  是拷贝

func f6() (x int) {
	defer func(x *int) { // *int 入参是int类型的指针
		*x++ // *x 取出x的值
	}(&x)
	return 5
}

// 返回值 x = 5
// *&x ++
// 所以返回值是6

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
	fmt.Println(f5())
	fmt.Println(f6())
	a := 1
	b := 2
	// defer 只计一层函数
	defer calc("1", a, calc("10", a, b)) // ("1",1,3)
	a = 0
	defer calc("2", a, calc("20", a, b)) // ("2",0,2)
	b = 1
}

// 10 1 2 3
// 20 0 2 2
// 2 0 2 2
// 1 1 3 4
