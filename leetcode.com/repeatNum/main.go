package main

import "fmt"

var res int

func findRepeatNumber(nums []int) int {
	dict := make(map[int]int, len(nums))

	for _, num := range nums {
		v, ok := dict[num]
		if ok {
			dict[num] = v + 1
		} else {
			dict[num] = 1
		}
	}
	fmt.Printf("%v", dict)

	for k, v := range dict {
		if v != 1 {
			res = k
			break
		}
	}
	return res
}

func main() {
	list := []int{2, 3, 1, 0, 2, 5, 3}
	res := findRepeatNumber(list)
	fmt.Println(res)
}
