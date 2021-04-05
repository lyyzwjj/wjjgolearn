package main

import (
	"fmt"
	myCalc "wjjgolearn/02liwenzhou/052package/10calc" // 包名字不合法时候  取别名
	"wjjgolearn/02liwenzhou/052package/calc"
	// "github.com/wzzst310/wjjgolearn/02liwenzhou/052package/calc"
)

func main() {
	ret := calc.Add(1, 2)
	fmt.Println(ret)
	ret = myCalc.Sub(2, 1)
	fmt.Println(ret)
}
