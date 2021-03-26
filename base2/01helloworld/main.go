package main

import "fmt"

// go 本包下 go build
// -o 指定build后文件名字
// go build -o hello github.com/wzzst310/wjjgolearn/base2/day01/helloworld   一定是src后面的目录到main包下
// 如果出现了cannot find module providing package github.com/wzzst310/wjjgolearn/base2/day01/helloworld: working directory is not part of a module
// export GO111MODULE=auto 关闭go module

// go run 像执行脚本文件一样执行Go代码
// go install
// 1. 先编译得到一个可执行文件
// 2. 将可执行文件拷贝到`GOPATH/bin`下
// 跨平台编译 只需要指定目标操作系统的平台和处理器架构即可
// Windows下编译Linux可执行文件
// SET CGO_ENABLED=0 // 禁用CGO
// SET GOOS=linux //平台是linux
// SET GOARCH=amd64 // 目标处理器架构时amd64
// Windows下编译Mac可执行文件
// SET CGO_ENABLED=0 // 禁用CGO
// SET GOOS=darwin //平台是linux
// SET GOARCH=amd64 // 目标处理器架构时amd64
// go build
// MAC下编译Linux和Windows平台64位可执行程序
// CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
// Windows下编译Mac平台64位可执行程序
// SET CGO_ENABLED=0
// SET GOOS=darwin
// SET GOARCH=amd64
// go build
// Linux下编译Mac和Windows平台64位可执行程序
// CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
// CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
func main() {
	fmt.Println("Hello World")
}
