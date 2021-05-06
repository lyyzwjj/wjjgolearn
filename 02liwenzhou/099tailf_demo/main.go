package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

// tailf的用法示例

// go get github.com/hpcloud/tail

func main() {
	fileName := "./my.log"
	config := tail.Config{
		ReOpen: true, // 重新打开  日志分隔时候 重新打开文件
		Follow: true, // 是否跟随   原本是my.log  变成my20210507.log  没读完要跟着读
		Location: &tail.SeekInfo{ // 从文件的哪个地方开始读  filebeat  存储了 读到的位置 registry中
			Offset: 0,
			Whence: 2,
		},
		MustExist: false, // 文件不存在不报错
		Poll:      true,
	}
	tails, err := tail.TailFile(fileName, config) // 读取文件内容放到tails 通道中
	if err != nil {
		fmt.Println("tail fail failed, err", err)
		return
	}
	var (
		msg *tail.Line
		ok  bool
	)
	for {
		msg, ok = <-tails.Lines // 从tails通道中读取文件内容
		if !ok {
			fmt.Printf("tail file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second) // 如果没读到睡1s
			continue
		}
		fmt.Println("msg", msg.Text)
	}

}
