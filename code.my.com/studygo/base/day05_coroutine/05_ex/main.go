package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

/*
使用goroutine和channel实现一个计算int64随机数个位数和的程序
1. 开启一个goroutine循环生成int64类型的随机数，发送到jobChan
2. 开启24个goroutine从jobChan中取出随机数计算个位数和，将结果发送到resultChan
3. 主goroutine从resultChan取出结果并打印到终端输出
*/
 
type job struct {
	value int64
}

type result struct {
	job *job
	value int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func create_random(num chan<- *job){
	defer wg.Done()
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		num <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func get_random(cr <-chan *job, rc chan<- *result) {
	defer wg.Done()
	for {
		job := <-cr
		sum := int64(0)
		n := job.value
		for n > 0{
			sum += n % 10
			n /= 10
		}
		res := &result{
			job: job,
			value: sum,
		}
		rc <- res
	}
}

func main() {
	wg.Add(1)
	go create_random(jobChan)
	wg.Add(24)
	for i := 0; i < 24; i++{
		go get_random(jobChan, resultChan)
	}
	for res := range resultChan{
		fmt.Printf("value: %d sum: %d\n", res.job.value, res.value)
	}
	wg.Wait()
}