package main

import "fmt"

func main() {
	// 切片的定义
	var s1 []int
	var s2 []string

	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // true
	fmt.Println(s2 == nil) // true

	s1 = []int{1, 2, 3}
	s2 = []string{"SmallMu", "LittleMu", "hunt"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false
	fmt.Println(s2 == nil) // false

	// 长度和容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	// 2. 由数组得到切片
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4]
	fmt.Println(s3)
}
