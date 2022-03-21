package main

import (
	"fmt"
)

func main(){
	// f1 := 1.23456
	// fmt.Printf("%T\n", f1)	// 默认go语言中的小数都是float64类型

	// f2 := float32(1.23456)
	// fmt.Printf("%T\n", f2) // 显示声明float32类型

	// b1 := true
	// var b2 bool	// 默认为false
	// fmt.Printf("%T\n", b1)
	// fmt.Printf("%T value: %v\n", b2, b2)

	var n = 100

	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%o\n", n)

	var s = "hello"
	fmt.Printf("%s\n", s)
	fmt.Printf("%v\n", s)
	fmt.Printf("%#v\n", s)

}