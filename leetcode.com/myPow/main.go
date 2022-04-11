package main

import "fmt"

func myPow(a, b int) int {
	if a == 0 {
		return 0
	}
	ans, base := 1, a
	for b != 0 {
		if b&1 == 1 {
			ans *= base
		}
		b >>= 1
		base *= base
	}
	return ans
}

func main() {
	res := myPow(2, 3)
	fmt.Println(res)
}
