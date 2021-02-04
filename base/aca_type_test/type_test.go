package ca_type_test

import "testing"

type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	//  b = a //报错  不支持隐式类型转换
	b = int64(a)
	var c MyInt
	// c = b  必须显示类型转换
	c = MyInt(b)
	t.Log(a, b, c)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	t.Log(a, aPtr)
	// aPtr = aPtr + 1 //go语言不支持指针运算
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var s string // 初始化是空字符串
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" { // 判断空值

	}
	//if s == nil {
	//
	//}
}
