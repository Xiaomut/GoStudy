package main

import (
	"sync/atomic"
	"fmt"
	"sync"
)

const (
	thread_num = 100000
)

var (
	x int64
	wg sync.WaitGroup
	lock sync.Mutex
)

func add(){
	// lock.Lock()
	// x ++
	// lock.Unlock()
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	wg.Add(thread_num)
	for i := 0; i < thread_num; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
}