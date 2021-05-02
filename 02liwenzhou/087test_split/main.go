package main

import (
	"fmt"
	"wjjgolearn/02liwenzhou/086split_string"
)

// Split 切割字符串
// example:
// abc, b=> [a c]

func main() {
	ret := _86split_string.Split("babcbef", "b")
	fmt.Printf("%#v\n", ret)

	ret2 := _86split_string.Split("bbb", "b")
	fmt.Printf("%#v\n", ret2)

	ret3 := _86split_string.Split("ejosada", "b")
	fmt.Printf("%#v\n", ret3)
}
