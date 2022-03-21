package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

func hello()  {
	fmt.Println("hello")
}

func f()  {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++{
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func f1(i int) {
	defer wg.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(300)))
	fmt.Println(i)
}

var wg sync.WaitGroup

func main()  {
	// go hello()
	// fmt.Println("main")
	// time.Sleep(time.Second)
	for i:=0; i<5; i++{
		wg.Add(1)
		go f1(i)
	}
	wg.Wait()
}