package main

import (
	"fmt"
)

var (
	coins = 50000
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	var left int = 0
	for _, name := range users{
		for _, alpha := range name{
			switch alpha{
			case 'e', 'E':
				distribution[name] += 1
				left += 1
			case 'i', 'I':
				distribution[name] += 2
				left += 2
			case 'o', 'O':
				distribution[name] += 3
				left += 3
			case 'u', 'U':
				distribution[name] += 4
				left += 4
			default:
				continue
			}
		}
		// fmt.Println(left)
	}
	return coins - left
}

func main()  {
	left := dispatchCoin()
	fmt.Println("The left: ", left)
	fmt.Println(distribution)
	for k, v := range distribution{
		fmt.Printf("%s%d\n", k, v)
	}
}