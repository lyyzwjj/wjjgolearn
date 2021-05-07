package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// etcd client put/get demo
// use etcd/clientV3
// go get go.etcd.io/etcd/client/v3

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"}, // endpoints就是集群节点
		DialTimeout: 5 * time.Second,
	})
	// watch操作
	// watch操作用来获取未来更改的通知.
	if err != nil {
		// handle error !
		fmt.Printf("connect to etcd failed, err:%v\n", err)
		return
	}
	fmt.Println("connect to etcd success")
	defer cli.Close()
	// watch
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 派一个哨兵一直监视着 qimi这个key的变化(新增\修改\删除)
	ch := cli.Watch(context.Background(), "qimi")
	// cancel() // 立即释放掉和context 相关的资源 类似gc
	// 从通道尝试取值(监视的信息)
	for wresp := range ch {
		// watchResponse 对象
		for _, evt := range wresp.Events {
			fmt.Printf("Type:%v key:%v value:%v\n", evt.Type, string(evt.Kv.Key), string(evt.Kv.Value))

		}
	}
}
