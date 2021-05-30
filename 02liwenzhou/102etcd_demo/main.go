package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
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
	// put
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, "qimi", "dsb")
	// _, err = cli.Put(ctx, "qimi", "dsb",clientv3.WithPrefix()) 可以设置前缀
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}
	// get
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "qimi")
	cancel()
	if err != nil {
		fmt.Printf("get from etcd failed, err:%v\n", err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s:%s\n", ev.Key, ev.Value)
	}
	// 续期
	// etcd 续期 5秒内客户端要返回心跳
	r, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}
	_, err = cli.Put(context.TODO(), "root", "admin", clientv3.WithLease(r.ID))
	if err != nil {
		log.Fatal(err)
	}
	// 自动续期   			// 续期对象ID
	ch, err := cli.KeepAlive(context.TODO(), r.ID)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c := <-ch
		fmt.Println("c:", c)
	}

}
