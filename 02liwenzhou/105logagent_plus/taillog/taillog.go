package taillog

import (
	"context"
	"fmt"
	"github.com/hpcloud/tail"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/kafka"
)

// 专门从日志文件收集之日志的模块

func (t *TailTask) ReadChan() <-chan *tail.Line {
	return t.instance.Lines
}

// TailTask 一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
	// 为了实现能退出t.run()
	ctx        context.Context
	cancelFunc context.CancelFunc
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	ctx, cancel := context.WithCancel(context.Background())
	tailObj = &TailTask{
		path:       path,
		topic:      topic,
		ctx:        ctx,
		cancelFunc: cancel,
	}
	tailObj.init() // 根据路劲去打开对应的日志
	return
}

func (t *TailTask) init() {
	config := tail.Config{
		ReOpen: true, // 重新打开  日志分隔时候 重新打开文件
		Follow: true, // 是否跟随   原本是my.log  变成my20210507.log  没读完要跟着读
		Location: &tail.SeekInfo{ // 从文件的哪个地方开始读  filebeat  存储了 读到的位置 registry中
			Offset: 0,
			Whence: 2,
		},
		MustExist: false, // 文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config) // 读取文件内容放到tails 通道中
	if err != nil {
		fmt.Println("tail fail failed, err", err)
	}
	// 当roroutine执行的函数退出的时候, goroutine就结束了
	go t.run() // 直接去采集日志发送到kafka
}

func (t *TailTask) run() {
	for {
		select {
		case <-t.ctx.Done():
			fmt.Printf("tail task:%s_%s 结束了...\n", t.path, t.topic)
			return
		case line := <-t.instance.Lines: // 从tailObj通道中一行一行的读取日志数据
			// 3.2 发往kafka
			// kafka.SentToKafka(t.topic, line.Text) // 函数调用函数 发送到kafka要排队发送
			// 先把日志数据发送到一个缓冲通道中
			fmt.Printf("正在发送缓冲区topic: %v, msg: %v", t.topic, line.Text)
			kafka.SendToChan(t.topic, line.Text)
			// kafka那个包中单独有goroutine去取日志数据发送到kafka
		}
	}

}
