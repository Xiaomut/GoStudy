package main

import (
	"log"
	"os"
)

var (
	newFile *os.File
	err     error
)

func main() {
	newFile, err = os.Create("X:/GoStudy/src/code.my.com/studygo/fileop/case01_newfile/test.txt")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
	newFile.Close()
}
