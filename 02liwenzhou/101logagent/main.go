package _101logagent

import (
	"fmt"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/101logagent/kafka"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/101logagent/taillog"
	"time"
)

// logAgent入口程序

func run() {
	// 1. 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SentToKafka("web_log", line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}

func main() {
	// 1. 初始化kafka连接
	err := kafka.Init([]string{"127.0.0.1:9092"})
	if err != nil {
		fmt.Printf("init Kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2. 打开日志文件准备日志
	err = taillog.Init("./my.log")
	if err != nil {
		fmt.Printf("Init taillog failed, err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")
	run()
}
