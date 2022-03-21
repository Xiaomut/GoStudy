package main

import (
	"log"
	"fmt"
	"os"
)

func main()  {
	f, err := os.OpenFile("./my.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		fmt.Println("open failed: ", err)
	}
	log.SetOutput(f)
	for {
		log.Println("This is a test log.")
	}
}