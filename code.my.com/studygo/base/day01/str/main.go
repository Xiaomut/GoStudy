package main

import (
	"strings"
	"fmt"
)

func main()  {
	s1 := "I'm ok"
	fmt.Println(s1)
	s2 := `X:\GoStudy`
	fmt.Println(s2)

	// 1. 分割
	s3 := strings.Split(s2, "\\")
	fmt.Println(s3)

	// 2. 包含
	fmt.Println(strings.Contains(s2, "go"))

	// 3. 前缀后缀
	fmt.Println(strings.HasPrefix(s2, "X"))
	fmt.Println(strings.HasSuffix(s2, "d"))

	// 4. 长度
	fmt.Println(len(s1))
	name := "hunt"
	age := "24"

	// 5. 相加
	ws := name + age
	fmt.Println(ws)
	ws1 := fmt.Sprintf("%s%s", name, age)
	fmt.Println(ws1)
	
	// 6. 子串
	s4 := "abcdefc"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "c"))
}