package mylogger

import (
	"fmt"
	"os"
	"path"
	"time"
)

// 往文件里面写日至相关的代码

type FileLogger struct {
	Level       LogLevel
	filePath    string // 日志文件保存的路径
	fileName    string // 日志文件保存的文件名
	fileObj     *os.File
	errFileObj  *os.File
	maxFileSize int64 // 日志文件大小
}

// 构造函数

func NewFileLogger(levelStr, fp, fn string, maxSize int64) *FileLogger {
	logLevel, err := parseLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	fl := &FileLogger{
		Level:       logLevel,
		filePath:    fp,
		fileName:    fn,
		maxFileSize: maxSize,
	}
	err = fl.initFile() //按照文件路径和文件名将文件打开
	if err != nil {
		panic(err)
	}
	return fl
}

func (f *FileLogger) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName+".log")    // 屏蔽操作系统拼接路径
	errFileName := path.Join(f.filePath, f.fileName+".err.log") // 屏蔽操作系统拼接路径
	fileObj, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open log file failed,err:%v\n", err)
		return err
	}
	errFileObj, err := os.OpenFile(errFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open err log file failed,err:%v\n", err)
		return err
	}
	// 日志文件都已经打开了
	f.fileObj = fileObj
	f.errFileObj = errFileObj
	return nil
}

func (f *FileLogger) Close() {
	f.fileObj.Close()
	f.errFileObj.Close()
}
func (f *FileLogger) enable(logLevel LogLevel) bool {
	return logLevel >= f.Level
}

func (f *FileLogger) log(lv LogLevel, format string, a ...interface{}) {
	if f.enable(lv) {
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcName, fileName, lineNo := getInfo(3)
		// fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), formatLogLevel(lv), fileName, funcName, lineNo, msg)
		fmt.Fprintf(f.fileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), formatLogLevel(lv), fileName, funcName, lineNo, msg)
		if lv >= ERROR {
			// 如果要记录的日志大于等于ERROR级别,我还要在err日中文件中再记录一遍
			fmt.Fprintf(f.errFileObj, "[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), formatLogLevel(lv), fileName, funcName, lineNo, msg)
		}
	}
}

func (f *FileLogger) Debug(msg string, a ...interface{}) {
	f.log(DEBUG, msg, a...)
}

func (f *FileLogger) Info(msg string, a ...interface{}) {
	f.log(INFO, msg, a...)
}

func (f *FileLogger) Warning(msg string, a ...interface{}) {
	f.log(WARNING, msg, a...)
}

func (f *FileLogger) Error(msg string, a ...interface{}) {
	f.log(ERROR, msg, a...)
}

func (f *FileLogger) Fatal(msg string, a ...interface{}) {
	f.log(FATAL, msg, a...)
}
