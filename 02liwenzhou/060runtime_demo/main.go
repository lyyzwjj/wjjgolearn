package main

import (
	"fmt"
	"path"
	"runtime"
)

// runtime.Caller()

func main() {
	// go运行时 自带gc
	pc, file, line, ok := runtime.Caller(0)
	// Caller  一层一层的调用 0表示当前
	// ok 如果能取到 ok为true
	// pc 执行的函数信息
	// file 调用本函数的文件 调用getInfo 能拿到mylogger.go文件
	// 拿到行号
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	fmt.Println(pc)   // 17467899
	fmt.Println(file) // /Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/060runtime_demo/main.go
	fmt.Println(line) // 12
	f1(1)
}

func f1(skip int) (funcName, fileName string, lineNo int) {
	// go运行时 自带gc
	// pc, file, line, ok := runtime.Caller(0)
	pc, file, line, ok := runtime.Caller(skip)
	// Caller  一层一层的调用 0表示当前	1进一层
	// ok 如果能取到 ok为true
	// pc 执行的函数信息
	// file 调用本函数的文件 调用getInfo 能拿到mylogger.go文件
	// 拿到行号
	if !ok {
		fmt.Printf("runtime.Caller() failed\n")
		return
	}
	funcName = runtime.FuncForPC(pc).Name()
	fmt.Println(funcName)        // 1 为 main.main  0 为 main.f1
	fmt.Println(file)            // /Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/060runtime_demo/main.go
	fmt.Println(path.Base(file)) // main.go
	fmt.Println(line)            // 25
	return funcName, fileName, line
}
