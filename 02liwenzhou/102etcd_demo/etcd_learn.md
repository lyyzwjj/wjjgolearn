## etcd学习

### 单机启动./etcd 连接使用设置

> 1. export ETCDCTL_API=3
> 2. ./etcdctl --endpoints=http://127.0.0.1:2379 put wjj "dsb"
> 3. ./etcdctl --endpoints=http://127.0.0.1:2379 get wjj
> 3. ./etcdctl --endpoints=http://127.0.0.1:2379 del wjj