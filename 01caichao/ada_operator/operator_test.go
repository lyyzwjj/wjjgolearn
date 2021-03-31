package ada_operator

import (
	"fmt"
	"testing"
)

/**
 * @author  wjj
 * @date  2020/8/21 1:13 上午
 * @description   比较运算符
 */
const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// t.Log(a == c)   //长度不同不能直接比较
	t.Log(a == d)
}

// 清零操作符
func TestBitClear(t *testing.T) {
	// 1、如果右侧是0，则左侧数保持不变
	// 2、如果右侧是1，则左侧数一定清零
	fmt.Println(0 &^ 0) // 0
	fmt.Println(0 &^ 1) // 0
	fmt.Println(1 &^ 0) // 1
	fmt.Println(1 &^ 1) // 0
	//左 0111
	//右 0100
	//   0011   3
	fmt.Println(7 &^ 4) // 3
	//左 1101
	//右 0111
	//   1000 8
	fmt.Println(13 &^ 7) // 8
	//左 0111
	//右 1101
	//   0010 2
	fmt.Println(7 &^ 13) // 2
	// 把右侧为1的位数全给干掉

	a := 7 //0111
	// a := 1 //0001
	fmt.Println(Readable)
	// Readable 0001
	//          0110
	// Writable 0010
	a = a &^ Readable
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
