package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	cli *clientv3.Client // 声明一个全局的etcd 客户端
)

type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// Init 初始化ETCD的函数
func Init(addr string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{addr}, // endpoints就是集群节点
		DialTimeout: timeout,
	})
	// watch操作
	// watch操作用来获取未来更改的通知.
	if err != nil {
		// handle error !
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	return
}

// GetConf 从etcd中根据key获取配置项
func GetConf(key string) (logEntryConf []*LogEntry, err error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		// fmt.Printf("%s:%s\n", ev.Key, ev.Value)
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Printf("unmarshal etcd value failed,err:%v\n", err)
			return
		}
	}
	return
}

// WatchConf etcd watch
func WatchConf(key string, newConfCh chan<- []*LogEntry) {
	// 派一个哨兵一直监视着 qimi这个key的变化(新增\修改\删除)
	ch := cli.Watch(context.Background(), key)
	// cancel() // 立即释放掉和context 相关的资源 类似gc
	// 从通道尝试取值(监视的信息)
	for wresp := range ch {
		// watchResponse 对象
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))
			// 通知taillog.tskMgr
			// 1. 先判断操作的类型
			var newConf []*LogEntry
			if evt.Type != clientv3.EventTypeDelete {
				// 如果不是删除操作
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					fmt.Printf("unmarshal failed, err:%v\n", err)
					continue
				}
			}
			fmt.Printf(" get new conf:%v\n", newConf)
			newConfCh <- newConf
		}
	}
}
