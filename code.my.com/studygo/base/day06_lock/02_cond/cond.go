package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	romoveFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		defer c.L.Unlock()
		queue = queue[1:]
		fmt.Println("Remove fomr queue")
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to a queue")
		queue = append(queue, struct{}{})
		go romoveFromQueue(1 * time.Second)
		c.L.Unlock()
	}
}
