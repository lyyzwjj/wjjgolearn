package aja_string

import (
	"strconv"
	"strings"
	"testing"
	"unsafe"
)

/**
 * @author  wjj
 * @date  2020/8/24 11:55 下午
 * @description
 */
func TestString(t *testing.T) {
	var s string
	t.Log(s) //  初始化默认零值""
	s = "hello"
	t.Log(len(s))
	// s[1] = '3' //string 是不可变的byte slice
	s = "\xE4\xB8\xA5" // 可以存储任何二进制数据   16进制  严
	// s = "\xE4\xBA\xB5\xFF"
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s))  // 是byte数  3个字节
	c := []rune(s) // 取出字符串里面的unicode
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人名共和国"
	for _, c := range s {
		// t.Logf("%[1]c %[1]d", c) // 两个用同一个变量格式化
		t.Logf("%[1]c %[1]x", c) // 两个用同一个变量格式化  16进制
	}
}

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Logf(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestStringConv(t *testing.T) {
	// 数字转字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)
	// 字符串转数字
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}

}
