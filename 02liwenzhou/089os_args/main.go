package main

import (
	"fmt"
	"os"
)

// OS.Args 获取命令行参数

func main() {
	fmt.Printf("%#v\n", os.Args) // []string{"./main"}
	fmt.Println(os.Args[0])
	fmt.Printf("%T\n", os.Args) // []string
}
