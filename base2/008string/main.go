package main

import (
	"fmt"
	"strings"
)

// 字符串

func main() {
	path := "\"D:\\Go\\src\\code.oldboyedu.com\\studygo\\day01\""
	fmt.Println(path)
	s := "I'm ok"
	fmt.Println(s)
	// 多行的字符串
	s2 := `
世情薄
				人情恶
		雨送黄昏花易落
	`
	fmt.Println(s2)
	s3 := "'D:/Go/src/code.oldboyedu.com/studygo/day01'"
	fmt.Println(s3)
	// 字符串相关操作
	fmt.Println(len(s3))

	// 字符串拼接
	name := "理想"
	world := "大帅比"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	fmt.Printf("%s%s\n", name, world)
	// 分隔
	ret := strings.Split(path, "\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss, "理性"))
	fmt.Println(strings.Contains(ss, "理想"))
	// 前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s4 := "abcdeb"
	// 索引
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "g"))
	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}
