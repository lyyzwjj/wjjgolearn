package main

import (
	"encoding/json"
	"fmt"
)

// json

type Person struct {
	Name string `json:"name"` // 字段一定要大写才能被外界感知 才能被json	// 跨包
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"周林","age":26}`
	var p Person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)
	// var m interface{}
	//m = make(map[interface{}]interface{})
	m := make(map[interface{}]interface{})
	m["name"] = "周林"
	m["age"] = 26
	var i interface{}
	i = m
	marshal, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(marshal)
}
