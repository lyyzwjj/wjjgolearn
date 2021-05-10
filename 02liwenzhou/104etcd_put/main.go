package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

// etcd client put/get demo

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
	value := `[{"path":"/Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/nginx.log","topic":"web_log"},{"path":"/Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/redis.log","topic":"redis_log"},{"path":"/Users/wjj/go/src/github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/mysql.log","topic":"mysql_log"}]`
	_, err = cli.Put(ctx, "/logagent/192.168.1.11/collect_config", value)
	// _, err = cli.Put(ctx, "qimi", "dsb",clientv3.WithPrefix()) 可以设置前缀
	cancel()
	if err != nil {
		fmt.Printf("put to etcd failed, err:%v\n", err)
		return
	}

}
