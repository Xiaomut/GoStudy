package mylogger

import (
	"errors"
	"strings"
	"fmt"
	"time"
)

// 终端部分写日志

// ConsoleLogger 结构体
type ConsoleLogger struct{
	Level LogLevel
}

func parse_log_level(s string) (LogLevel, error) {
	s = strings.ToLower(s)
	switch s{
	case "debug":
		return DEBUG, nil
	case "trace":
		return TRACE, nil
	case "info":
		return INFO, nil
	case "warning":
		return WARNING, nil
	case "error":
		return ERROR, nil
	case "fatal":
		return FATAL, nil
	default:
		err := errors.New("has not implemented")
		return UNKNOWN, err
	} 
}

// 构造函数
func NewConsoleLog(level_str string) ConsoleLogger  {
	level, err := parse_log_level(level_str)
	if err != nil{
		panic(err)
	}
	return ConsoleLogger{
		Level: level,
	}
}

func (c ConsoleLogger) enable(loglevel LogLevel) bool  {
	return loglevel >= c.Level
}

func (c ConsoleLogger) get_log(lv LogLevel, format string, a ...interface{}){
	if c.enable(lv){
		msg := fmt.Sprintf(format, a...)
		now := time.Now()
		funcname, filename, linenum := getInfo(3)
		var status string
		switch lv{
		case DEBUG:
			status = "DEBUG"
		case TRACE:
			status = "TRACE"
		case INFO:
			status = "INFO"
		case WARNING:
			status = "WARNING"
		case ERROR:
			status = "ERROR"
		case FATAL:
			status = "FATAL"
		default:
			status = "UNKNOWN"
		}
		fmt.Printf("[%s] [%s] [%s:%s:%d] %s\n", now.Format("2006-01-02 15:04:05"), status, funcname, filename, linenum, msg)
	}
}


func (c ConsoleLogger) Debug(format string, a ...interface{}){
	c.get_log(DEBUG, format, a...)
}

func (c ConsoleLogger) Info(format string, a ...interface{}){
	c.get_log(INFO, format, a...)
}

func (c ConsoleLogger) Warning(format string, a ...interface{}){
	c.get_log(WARNING, format, a...)
}

func (c ConsoleLogger) Error(format string, a ...interface{})  {
	c.get_log(ERROR, format, a...)
}

func (c ConsoleLogger) Fatal(format string, a ...interface{})  {
	c.get_log(FATAL, format, a...)
}