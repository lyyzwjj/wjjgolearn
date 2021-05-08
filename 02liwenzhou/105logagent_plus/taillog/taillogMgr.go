package taillog

import (
	"github.com/wzzst310/wjjgolearn/02liwenzhou/105logagent_plus/etcd"
)

var tskMgr *tailLogMgr

// taillog管理者

type tailLogMgr struct {
	logEntry []*etcd.LogEntry
	// tskMap map[string]*TailTask
}

func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &tailLogMgr{
		logEntry: logEntryConf, // 把当前的日志收集项配置信息保存起来
	}
	for _, logEntry := range logEntryConf {
		// conf: *etcd.LogEntry
		// logEntry.Path: 要收集的日志文件的路径
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
