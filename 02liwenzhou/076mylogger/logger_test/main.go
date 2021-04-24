package main

import (
	"wjjgolearn/02liwenzhou/076mylogger/mylogger"
)

var log mylogger.Logger // 声明一个全局的全局变量

// 测试我们自己写的日志库
func main() {
	// log = mylogger.NewConsoleLogger("INFO")                         //  终端日志实例
	log = mylogger.NewFileLogger("INFO", "/Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/076mylogger/logger_test/", "customer", 10*1024) //  文件日志实例
	// log = mylogger.NewLogger("f")                                   //根据入参返回不同的Logger
	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条info日志")
		id := 10010
		name := "理想"
		log.Error("这是一条error日志,id:%d, name:%s", id, name)
		log.Fatal("这是一条fatal日志")
		// time.Sleep(2 * time.Second)
	}
}
