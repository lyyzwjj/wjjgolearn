package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 打开文件写内容
// 111	三个标志位
// 第一个二进制位表示打开是否创建
// 第二个二进制位表示追加还是清空
// 第三个二进制位表示额外其他的
// os.O_RDONLY 0x0
// os.O_APPEND 0x8
// os.O_CREATE 0x200
func main() {
	// writeDemo2()
	writeDemo3()
}

func writeDemo3() {
	str := "hello 沙雕"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
func writeDemo2() {
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("read file err:%v\n", err)
	}
	defer fileObj.Close()
	// 创建一个写的对象
	wr := bufio.NewWriter(fileObj)
	wr.WriteString("hello沙河\n") // 写到缓存中
	wr.Flush()                  // 将缓存中的内容写入文件	刷盘  将内存中的写到文件中
	// 同步刷盘 内存及时写到文件中 效率低 但是数据不会丢失
	// 异步刷盘 内存中的内容定时写到文件中  吞吐量增加 但是发生断电等情况 内存中 缓冲区的内容数据会丢失
}

func writeDemo1() {
	// 												两位有一位为1 即为1		追加
	// fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	// O_TRUNC 每次都是删除旧的来新的
	fileObj, err := os.OpenFile("./xx.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Printf("read file err:%v\n", err)
	}
	fileObj.Write([]byte("zhoulin mengbi le!\n"))
	fileObj.WriteString("周林解释不了! ")
	fileObj.Close()
}
