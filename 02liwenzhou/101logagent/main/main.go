package main

import (
	"fmt"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/101logagent/conf"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/101logagent/kafka"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/101logagent/taillog"
	"gopkg.in/ini.v1"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

// logAgent入口程序
// ./kafka-console-consumer.sh --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning 手动消费
func run() {
	// 1. 读取日志
	for {
		select {
		case line := <-taillog.ReadChan():
			// 2. 发送到kafka
			kafka.SentToKafka(cfg.KafkaConf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}

}

func main() {
	// 0. 加载配置文件
	//cfgRow, err := ini.Load("../conf/config.ini")
	//fmt.Println(cfgRow.Section("kafka").Key("address").String())
	//fmt.Println(cfgRow.Section("kafka").Key("topic").String())
	//fmt.Println(cfgRow.Section("taillog").Key("path").String())

	err := ini.MapTo(cfg, "../conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	// 1. 初始化kafka连接
	err = kafka.Init([]string{cfg.KafkaConf.Address})
	if err != nil {
		fmt.Printf("init Kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")
	// 2. 打开日志文件准备日志
	err = taillog.Init(cfg.TailllogConf.FileName)
	if err != nil {
		fmt.Printf("Init taillog failed, err:%v\n", err)
		return
	}
	fmt.Println("init taillog success")
	run()
}
