package main

import (
	"fmt"
	"time"
	// "sync"
)

// var wg sync.WaitGroup
var notifyChan = make(chan struct{}, 5)

func worker(id int, jobs <-chan int, results chan<- int)  {
	for j := range jobs{
		time.Sleep(time.Second)
		fmt.Printf("worker: %d end job: %d\n", id, j)
		results <- j * 2
		notifyChan <- struct{}{}
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 5 tasks
	go func(){
		for j := 1; j <= 5; j++{
			jobs <- j
		}
		close(jobs)
	}()

	// open 3 goroutines
	for i := 1; i <= 3; i++{
		go worker(i, jobs, results)
	}

	go func(){
		for k := 0; k < 5; k++{
			<- notifyChan
		}
		close(results)
	}()

	// outputs
	for x := range results{
		fmt.Println(x)
	}

}