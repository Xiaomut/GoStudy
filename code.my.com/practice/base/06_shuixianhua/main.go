package main

import (
	"fmt"
)

/*
题目：打印出所有的 “水仙花数”，所谓 “水仙花数” 是指一个三位数，其各位数字立方和等于该数本身。例如：153 是一个 “水仙花数”，因为 153=1 的三次方＋5 的三次方＋3 的三次方。

答案：https://haicoder.net/case/golang-hundred-cases/golang-1-13.html
*/

func main() {

	for i := 1; i < 9; i++ {
		for j := 0; j < 9; j++ {
			for k := 0; k < 9; k++ {
				num := i*i*i + j*j*j + k*k*k
				if num == (100*i + 10*j + k) {
					fmt.Println(num)
				}
			}
		}
	}
}
