package main

import (
	"fmt"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/conf"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/etcd"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/kafka"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/taillog"
	"gopkg.in/ini.v1"
	"time"
)

var (
	cfg = new(conf.AppConf)
)

// logAgent入口程序
// ./kafka-console-consumer.sh --bootstrap-server=127.0.0.1:9092 --topic=web_log --from-beginning 手动消费

func main() {
	// 0. 加载配置文件

	err := ini.MapTo(cfg, "./conf/config.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	// 1. 初始化kafka连接
	// fmt.Println(cfg.KafkaConf.Address)
	err = kafka.Init([]string{cfg.KafkaConf.Address}, cfg.KafkaConf.ChanMaxSize)
	if err != nil {
		fmt.Printf("init Kafka failed, err:%v\n", err)
		return
	}
	fmt.Println("init kafka success")

	// 2. 初始化etcd
	err = etcd.Init(cfg.EtcdConf.Address, time.Duration(cfg.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Printf("init etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("init etcd success")

	// 2.1 从etcd中获取日志收集项的配置信息
	logEntryConf, err := etcd.GetConf(cfg.EtcdConf.Key)
	// 2.2 派一个哨兵去监视日志收集项的变化(有变化及时通知我的logAgent实现热加载配置)
	if err != nil {
		fmt.Printf("etcd GetConf failed,err:%v\n", err)
		return
	}
	// fmt.Printf("get conf from etcd success, %v\n", logEntryConf)
	for index, value := range logEntryConf {
		fmt.Printf("index:%v value:%v\n", index, value)
	}
	// 3. 收集日志发往kafka
	// 3.1 循环每一个日志收集项,创建TailObj
	taillog.Init(logEntryConf)
	time.Sleep(1000 * time.Hour)

	// 2. 打开日志文件准备日志
	//fmt.Println(cfg.TailllogConf.FileName)
	//err = taillog.Init(cfg.TailllogConf.FileName)
	//if err != nil {
	//	fmt.Printf("Init taillog failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println("init taillog success")
	//run()
}
