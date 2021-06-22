package _01csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

func TestReadCsv(t *testing.T) {
	csvFilePath := "resources/mobile.csv"
	// readAll(csvFilePath)
	readLineByLine(csvFilePath)
}
func readAll(csvFilePath string) {
	f, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	reader := csv.NewReader(f)
	// 可以一次性读完
	result, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(reflect.TypeOf(result))
	fmt.Println(result)
}

func readLineByLine(csvFilePath string) {
	// 也可以一行一行进行读取
	// 但是注意不要两种方式都使用
	// 运行本代码第二种方式有可能没有数据，因为读指针已经指到了最后
	f, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	reader := csv.NewReader(f)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		fmt.Println(line)
		fmt.Println(reflect.TypeOf(line))
	}
}

func TestWriteCsv(t *testing.T) {
	csvFilePath := "resources/mobile_write.csv"
	f, err := os.OpenFile(csvFilePath, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	writer := csv.NewWriter(f)
	var header = []string{"id", "name", "age"}
	writer.Write(header)
	var data = []string{"3", "John", "23"}
	writer.Write(data)
	// 也可以一次性写入多条
	var d = [][]string{{"1", "Edgar", "20"}, {"2", "Tom", "18"}}
	writer.WriteAll(d)
	// 将缓存中的内容写入到文件里
	writer.Flush()
	if err = writer.Error(); err != nil {
		fmt.Println(err)
	}
}
