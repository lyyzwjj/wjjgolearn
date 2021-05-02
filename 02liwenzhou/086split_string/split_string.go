package _86split_string

import (
	"fmt"
	"strings"
)

// Split 切割字符串
// example:
// abc, b=> [a c]

func Split(str string, sep string) []string {
	// str:"babcbef" sep="b" [a cdef]
	var ret []string
	index := strings.Index(str, sep)
	for index >= 0 {
		ret = append(ret, str[0:index])
		str = str[index+len(sep):]
		index = strings.Index(str, sep)
	}
	// 仅为测试覆盖率
	if index == -5 {
		fmt.Println("真无聊")
	}
	ret = append(ret, str)
	return ret
}
