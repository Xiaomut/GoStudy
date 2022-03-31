package main

import (
	"github.com/hpcloud/tail"
)

// tailf的用法
func main() {
	filename := "./my.log"
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll: true,
	} 
}