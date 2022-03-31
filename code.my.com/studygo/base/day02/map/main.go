package main

import (
	"fmt"
)

func main()  {
	// 元素为map的切片
	var s1 = make([]map[int]string, 10, 10)
	// 需要对map做初始化
	s1[0] = make(map[int]string, 1)
	s1[0][10] = "hunt"
	// [map[10:hunt] map[] map[] map[] map[] map[] map[] map[] map[] map[]]
	fmt.Println(s1)
}