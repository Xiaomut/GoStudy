package main

import "fmt"

/*
题目：古典问题：有一对兔子，从出生后第 3 个月起每个月都生一对兔子，小兔子长到第三个月后每个月又生一对兔子，假如兔子都不死，问每个月的兔子总数为多少？

答案：https://haicoder.net/case/golang-hundred-cases/golang-1-11.html
*/

func main() {

	var month int
	fmt.Scan(&month)

	fib := []int{1, 1}
	i := 2
	for {
		if month < (i + 1) {
			break
		}
		fib = append(fib, fib[i-2]+fib[i-1])
		i += 1
	}
	fmt.Printf("The mount is: %v", fib)
}
