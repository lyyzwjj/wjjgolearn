package conf

// go get gopkg.in/ini.v1

type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}
type KafkaConf struct {
	Address string `ini:"address"`
}
type EtcdConf struct {
	Address string `ini:"address"`
	Timeout int    `ini:"timeout"`
}

// -----used -----
type TailllogConf struct {
	FileName string `ini:"filename"`
}
