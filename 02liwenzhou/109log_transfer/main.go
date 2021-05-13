package main

import (
	"fmt"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/109log_transfer/conf"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/109log_transfer/es"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/109log_transfer/kafka"
	"gopkg.in/ini.v1"
)

// log transfer
// 将日志数据从kafka取出来发往es
func main() {
	// 0. 加载配置文件
	var cfg = new(conf.LogTransferCfg)
	err := ini.MapTo(&cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Printf("load ini failed, err:%v\n", err)
		return
	}
	fmt.Printf("cfg:%v\n", cfg)
	// 1.初始化ES
	// 1.1 初始刷一个ES连接的client
	err = es.Init(cfg.ESCfg.Address, cfg.ESCfg.ChanMaxSize, cfg.ESCfg.Nums)
	if err != nil {
		fmt.Printf("init ES client failed, err:%v\n", err)
		return
	}
	fmt.Println("init es success")
	// 2. 初始化kafka
	// 2.1 连接kafka, 创建分区的消费者
	// 2.2 每个分区的消费者分别取出数据 通过SendToES()将数据发往ES
	// fmt.Println(cfg.KafkaCfg.Address)
	// fmt.Println(cfg.KafkaCfg.Topic)
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Printf("init Kafka consumer failed, err:%v\n", err)
		return
	}
	// fmt.Println("init kafka success")
	select {}
	fmt.Println("init kafka success")
	// 1. 从kafka取日志数据
	// 2. 发往ES
	// 3.

}
