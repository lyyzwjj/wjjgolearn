package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
	"strings"
	"time"
)

// 初始化ES, 准备接受kafka那边发来的数据

var (
	client *elastic.Client
	ch     chan *LogDataWrap
)

type LogData struct {
	Data string `json:"data"`
}
type LogDataWrap struct {
	LogData LogData
	Topic   string
}

func Init(address string, chanSize, nums int) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	// 1. 初始化一个client客户端用来连接
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		// handle error
		return
	}
	fmt.Println("connect to es success")
	ch = make(chan *LogDataWrap, chanSize)
	for i := 0; i < nums; i++ {
		go SendToES()
	}
	return
}

// 发送数据到ES
func SendToESChan(msg *LogDataWrap) {
	ch <- msg
}
func SendToES() {
	for {
		select {
		case ld := <-ch:
			// 链式操作
			put1, err := client.
				Index().
				Index(ld.Topic).
				BodyJson(ld.LogData).
				Do(context.Background())
			if err != nil {
				// Handle error
				fmt.Println(err)
			}
			fmt.Printf("Indexed student %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}

}
