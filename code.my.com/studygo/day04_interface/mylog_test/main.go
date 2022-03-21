package main

import (
	// "time"
	"code.my.com/studygo/mylogger"
)

var log mylogger.Logger

func main(){
	// log = mylogger.NewConsoleLog("debug")
	log = mylogger.NewFileLogger("info", "./", "app.log", 10*1024)
	for {
		id := 100
		log.Debug("This is a debug log - %d", id)
		log.Info("This is an info log")
		log.Error("This is an error log")
		// time.Sleep(time.Second)
	}
}