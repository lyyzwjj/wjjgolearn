package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

// 1. 序列化: 把Go语言中的结构体变量 --> json格式的字符串
// 2. 反序列化: json格式的字符串  --> Go语言中能够识别的结构体变量
type Person struct {
	//name string
	//age  int	// json 中 db 中 ini 中 叫什么
	Name string `json:"name" db:"name" ini:"name"` // 字段一定要大写才能被外界感知 才能被json	// 跨包
	Age  int    `json:"age"`
}

type myfloat32 float32

func main() {
	i := myfloat32(3.88)
	marshal, err2 := json.Marshal(i)
	if err2 != nil {
		fmt.Printf("marshal failed, err:%v", err2)
		return
	}
	println(marshal)
	var i2 myfloat32
	err2 = json.Unmarshal([]byte(marshal), &i2) // 传指针是为了能在json.Unmarshal内部修改p2的值
	if err2 != nil {
		fmt.Printf("unmarshal failed, err:%v", err2)
		return
	}
	fmt.Printf("%#v\n", i2)

	p1 := Person{"周林", 26}
	// 序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b)) // 强转  字符串本身就是有字节数组组成的
	// 反序列化
	str := `{"name":"理想","age":18}`
	var p2 Person
	err = json.Unmarshal([]byte(str), &p2) // 传指针是为了能在json.Unmarshal内部修改p2的值
	if err != nil {
		fmt.Printf("unmarshal failed, err:%v", err)
		return
	}
	fmt.Printf("%#v\n", p2)
}
