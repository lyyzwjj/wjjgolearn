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
}
