package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/lyyzwjj/wjjgolearn/02liwenzhou/109log_transfer/es"
)

// 初始化kafka消费者 从kafka取数据发往ES

// Init 初始化
func Init(addrs []string, topic string) error {
	fmt.Println("addrs", addrs)
	fmt.Println("topic", topic)
	consumer, err := sarama.NewConsumer(addrs, nil)
	if err != nil {
		fmt.Printf("fail to start consumer, err:%v\n", err)
		return err
	}
	partitionList, err := consumer.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return err
	}
	fmt.Println("分区列表", partitionList)
	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			fmt.Printf("failed to start consumer for partition %d,err:%v\n", partition, err)
			return err
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n ", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
				// 直接发给ES
				ld := es.LogData{string(msg.Value)}
				msg := es.LogDataWrap{ld, topic}
				es.SendToESChan(&msg)
			}
		}(pc)
	}
	// time.Sleep(time.Second * 10)
	select {}
	return err
}
