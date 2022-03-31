package main

import (
	"fmt"
)

// 1. 类型断言

func main() {
	var a interface{} // 定义一个空接口

	a = 100 
	switch v := a.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("None")
	}


}