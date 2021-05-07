package conf

// go get gopkg.in/ini.v1

type AppConf struct {
	KafkaConf    `ini:"kafka"`
	TailllogConf `ini:"taillog"`
}
type KafkaConf struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}
type TailllogConf struct {
	FileName string `ini:"filename"`
}
