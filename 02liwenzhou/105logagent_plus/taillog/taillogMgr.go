package taillog

import (
	"fmt"
	"github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/etcd"
	"time"
)

var tskMgr *tailLogMgr

// taillog管理者
// tailLogMgr tailTask管理者
type tailLogMgr struct {
	logEntry    []*etcd.LogEntry
	tskMap      map[string]*TailTask
	newConfChan chan []*etcd.LogEntry
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry:    logEntryConf, // 把当前的日志收集项配置信息保存起来
		tskMap:      make(map[string]*TailTask, 16),
		newConfChan: make(chan []*etcd.LogEntry), // 故意做一个无缓冲区的通道 没有新数据会一直阻塞
	}
	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		// logEntry.Path: 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
	go tskMgr.run()
}

// 监听自己的newConfChan,有了新的配置过来之后就做对应的处理
func (t *tailLogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			// 1. 配置新增
			// 2. 配置删除
			// 3. 配置变更
			fmt.Println("新的配置来了!", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan 一个函数,向外暴露tskMgr的newConfChan 内部私有字段 向外暴露
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}

// PushNewConf 一个函数,向外暴露tskMgr的newConfChan 内部私有字段 向外暴露
func PushNewConf(newConf []*etcd.LogEntry) {
	tskMgr.newConfChan <- newConf
}
