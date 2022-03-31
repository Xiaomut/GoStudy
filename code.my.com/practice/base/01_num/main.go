package main

import "fmt"

/*
题目：有 1、2、3、4 这四个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？

答案：https://haicoder.net/case/golang-hundred-cases/golang-1-1.html
*/

func main() {
	count := 0

	for i := 1; i < 5; i++ {
		for j := 1; j < 5; j++ {
			if j == i {
				continue
			}
			for k := 1; k < 5; k++ {
				if j == k || k == i {
					continue
				}
				count += 1
				num := 100*i + 10*j + k
				fmt.Printf("The %d num is: %d\n", count, num)
			}
		}
	}
}
