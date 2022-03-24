package main

import (
	"fmt"
)

func main() {
	// 返回的是一个只读channel，resultStream被隐式的转换为只读消费者
	chanOwner := func() <-chan int {
		// 实例化一个缓冲channel
		resultStream := make(chan int, 5)
		// 启用匿名的 goroutine
		go func() {
			// 确保执行完成后通道关闭
			defer close(resultStream)
			for i := 0; i < 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for result := range resultStream {
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done Receiving!")
}

