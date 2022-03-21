package main

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var exitChan chan bool = make(chan bool, 1)

func f() {
	defer wg.Done()
	for {
		fmt.Println("hunt")
		time.Sleep(time.Millisecond * 300)
		select {
		case <-exitChan:
			break
		default:
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	// 如何通知goroutine退出
	exitChan <- true
	wg.Wait()
}