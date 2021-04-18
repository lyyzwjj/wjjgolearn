package main

import (
	"fmt"
	"strconv"
)

// strconv
// 	数字变字符串 ret3 := strconv.Itoa(i)
// 	字符串变数字 retInt, err := strconv.Atoi(str)

func main() {
	// 从字符串中解析出对应的数据
	str := "10000"
	// ret1 := int64(str)
	ret1, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		fmt.Println("parseInt failed,err:", err)
		return
	}
	fmt.Printf("%#v %T\n", ret1, int(ret1))
	// Atoi: 字符串转换成int
	retInt, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("parseInt failed,err:", err)
		return
	}
	fmt.Printf("%#v %T\n", retInt, retInt)

	// 从字符串中解析出布尔值
	boolStr := "true"
	retBool, err := strconv.ParseBool(boolStr)
	if err != nil {
		fmt.Println("parseBool failed,err:", err)
		return
	}
	fmt.Printf("%#v %T\n", retBool, retBool)
	// 从字符串中解析出浮点数
	floatStr := "1.234"
	retFloat, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		fmt.Println("parseBool failed,err:", err)
		return
	}
	fmt.Printf("%#v %T\n", retFloat, retFloat)
	// 把数字转换成字符串类型
	i := 97
	// ret2 :=string(i) // "a"
	ret2 := fmt.Sprintf("%d", i) // "97"
	fmt.Printf("%#v\n", ret2)
	// Itoa: int转换成字符串
	ret3 := strconv.Itoa(i)
	fmt.Printf("%#v\n", ret3)

}
