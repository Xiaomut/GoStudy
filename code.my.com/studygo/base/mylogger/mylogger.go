package mylogger

// 自定义日志

import (
	"strings"
	"runtime"
	"fmt"
	"path"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG 
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

type Logger interface{
	Debug(format string, a ...interface{})
	Info(format string, a ...interface{})
	Warning(format string, a ...interface{})
	Error(format string, a ...interface{})
	Fatal(format string, a ...interface{})
}

func getInfo(skip int) (funcname, filename string, linenum int) {
	pc, f, linenum, ok := runtime.Caller(skip)
	if !ok{
		fmt.Println("runtime Caller() failed: ", ok)
		return 
	}
	funcname = runtime.FuncForPC(pc).Name()
	funcname = strings.Split(funcname, ".")[1]
	filename = path.Base(f)

	return funcname, filename, linenum
}
