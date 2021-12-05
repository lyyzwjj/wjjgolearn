package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// 打开文件
// EOF end of file error

func main() {
	// readFromFile1()
	// readFromFileByBufio()
	readFromFileByIoutil()
}
func FileName() {
	csvFilePath := "resources/raw/大家的日语第二版初级2_26.csv"
	fileNameAll := path.Base(csvFilePath)
	fileSuffix := path.Ext(csvFilePath)
	// fileprefix := filenameall[0:len(filenameall) - len(filesuffix)]
	filePrefix := strings.TrimSuffix(fileNameAll, fileSuffix)
	fmt.Println("file name:", fileNameAll)
	fmt.Println("file prefix:", filePrefix)
	fmt.Println("file suffix:", fileSuffix)
}

func readFromFileByIoutil() {
	ret, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Printf("read file err:%v\n", err)
	}
	fmt.Println(string(ret))
}

// 利用bufio这个包读文件
func readFromFileByBufio() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:%v", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 创建一个用来从文件中读内容的对象
	reader := bufio.NewReader(fileObj)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("read line failed,err:%v", err)
			return
		}
		fmt.Print(line)
	}
}

func readFromFile1() {
	fileObj, err := os.Open("./main.go")
	if err != nil {
		fmt.Printf("open file failed, err:%v\n", err)
		return
	}
	// 记得关闭文件
	defer fileObj.Close()
	// 读文件
	// var tmp = make([]byte, 128) //指定读的长度
	var tmp [128]byte
	for {
		n, err := fileObj.Read(tmp[:])
		// 其他的处理
		/*		if err == io.EOF {
				fmt.Println("文件读完了")
				return
			}*/
		if err != nil {
			fmt.Printf("read from file failed,err:%v\n", err)
		}
		fmt.Printf("读了%d个字节\n", n)
		fmt.Println(string(tmp[:n]))
		if n < 128 {
			return
		}
	}
}
