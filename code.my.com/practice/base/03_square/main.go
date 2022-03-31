package main

import (
	"fmt"
	"sort"
)

func main() {
	// i := 0
	// for {
	// 	x := int(math.Sqrt(float64(i + 100)))
	// 	y := int(math.Sqrt(float64(i + 268)))
	// 	if x*x == i+100 && y*y == i+268 {
	// 		fmt.Println("this num is: ", i)
	// 		break
	// 	}
	// 	i += 1
	// }
	var x, y, z int

	fmt.Scanf("%d %d %d", &x, &y, &z)

	slice := []int{x, y, z}
	sort.Ints(slice)

	for _, v := range slice {
		fmt.Println(v)
	}
}
