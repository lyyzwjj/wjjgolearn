package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

// 专门从日志文件收集之日志的模块

var (
	tailObj *tail.Tail
)

func Init(fileName string) (err error) {
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
	tailObj, err = tail.TailFile(fileName, config) // 读取文件内容放到tails 通道中
	if err != nil {
		fmt.Println("tail fail failed, err", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
