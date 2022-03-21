package main

import (
	"fmt"
)

func main() {
	array := [...]int{1, 3, 5, 7, 9}
	sum := 0

	for _, i := range array {
		sum += i
	}
	fmt.Println(sum)
}
