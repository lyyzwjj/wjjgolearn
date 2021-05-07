package conf

type AppConf struct {
	KafkaConf
	TailllogConf
}
type KafkaConf struct {
	Address string
	Topic   string
}
type TailllogConf struct {
	FileName string
}
