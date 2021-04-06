package mylogger

import (
	"fmt"
	"time"
)

// 往终端写日志相关的内容

// ConsoleLogger 日志结构体
type ConsoleLogger struct {
	Level LogLevel
}

// NewConsoleLogger 构造函数
func NewConsoleLogger(levelStr string) ConsoleLogger {
	level, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLogger{Level: level}
}

func (c ConsoleLogger) log(lv LogLevel, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	now := time.Now()
	funcName, fileName, lineNo := getInfo(3)
	fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), formatLogLevel(lv), fileName, funcName, lineNo, msg)
}

func (c ConsoleLogger) enable(logLevel LogLevel) bool {
	return logLevel >= c.Level
}

func (c ConsoleLogger) Debug(msg string, a ...interface{}) {
	if c.enable(DEBUG) {
		c.log(DEBUG, msg, a...)
	}
}

func (c ConsoleLogger) Info(msg string, a ...interface{}) {
	if c.enable(INFO) {
		c.log(INFO, msg, a...)
	}
}

func (c ConsoleLogger) Warning(msg string, a ...interface{}) {
	if c.enable(WARNING) {
		c.log(WARNING, msg, a...)
	}
}

func (c ConsoleLogger) Error(msg string, a ...interface{}) {
	if c.enable(ERROR) {
		c.log(ERROR, msg, a...)
	}
}

func (c ConsoleLogger) Fatal(msg string, a ...interface{}) {
	if c.enable(FATAL) {
		c.log(FATAL, msg, a...)
	}
}
