package main

import (
	"time"
	"wjjgolearn/02liwenzhou/059mylogger/mylogger"
)

// 测试我们自己写的日志库
func main() {
	log := mylogger.NewLog("INFO")
	for {
		log.Debug("这是一条debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条info日志")
		log.Error("这是一条error日志")
		log.Fatal("这是一条fatal日志")
		time.Sleep(time.Second)
	}
}
