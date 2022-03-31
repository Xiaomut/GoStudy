package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// var exitChan chan bool = make(chan bool, 1)

func f(ctx context.Context) {
	defer wg.Done()
	go f2(ctx)
FORLOOP:
	for {
		fmt.Println("hunt")
		time.Sleep(time.Millisecond * 300)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func f2(ctx context.Context) {
	defer wg.Done()
FORLOOP:
	for {
		fmt.Println("xiaomu")
		time.Sleep(time.Millisecond * 300)
		select {
		case <-ctx.Done():
			break FORLOOP
		default:
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go f(ctx)
	time.Sleep(time.Second * 5)
	// 如何通知goroutine退出
	cancel()
	wg.Wait()
}
