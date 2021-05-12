package conf

// LogTransfer ...
type LogTransferCfg struct {
	KafkaCfg `ini:"kafka"`
	ESCfg    `ini:"es"`
}

// Kafka ...
type KafkaCfg struct {
	Address string `ini:"address"`
	Topic   string `ini:"topic"`
}

// ESCfg ...
type ESCfg struct {
	Address     string `ini:"address"`
	ChanMaxSize int    `ini:"chan_max_size"`
	Nums        int    `ini:"nums"`
}
