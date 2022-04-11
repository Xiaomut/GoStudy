package main

import "fmt"

func add(a int, b int) int {
	step := 0
	for b != 0 { //相当于进位不等于0
		step = (a & b) << 1 //算出进位
		a ^= b              //算出不带进位的和
		b = step            //更新进位
	}
	return a

}

func main() {
	res := add(1, 1)
	fmt.Println(res)
}
