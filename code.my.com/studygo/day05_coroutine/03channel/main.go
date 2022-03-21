package main

import (
	"sync"
	"fmt"
)

var a []int
var b chan int
var wg sync.WaitGroup

func noBufChannel()  {
	fmt.Println(b)
	b = make(chan int) //// 不带缓冲区通道的初始化
	wg.Add(1)
	go func(){
		defer wg.Done()
		x := <-b
		fmt.Println("let 10 send in channel", x)
	}()

	b <- 10
	fmt.Println("let 10 send in channel")
	wg.Wait()
}

func bufChannel()  {
	fmt.Println(b) // nil
	b = make(chan int, 1) //// 带缓冲区通道的初始化
	b <- 10
	fmt.Println("let 10 send in channel")
}


func main()  {
	bufChannel()
}