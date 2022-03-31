package main

import "fmt"

/*
题目：​企业发放的奖金根据利润提成。利润(I)低于或等于 10 万元时，奖金可提成 10%；利润高于 10 万元，低于 20 万元，低于 10 万元的部分按 10% 提成，高于 10 万元的部分，可提成 7.5%。
​20 万到 40 万之间时，高于 20 万元的部分，可提成 5%；40 万到 60 万之间时高于 40 万元的部分，可提成 3%；60 万到 100 万之间时，高于 60 万元的部分，可提成 1.5%，高于 100 万元时，超过 100 万元的部分按 1% 提成。
​从键盘输入当月利润 I，求应发放奖金总数？

答案：https://haicoder.net/case/golang-hundred-cases/golang-1-2.html
*/

func main() {
	// var profit float
	// fmt.Scan(&profit)
	profit := 52.
	var bonus float64

	if profit <= 10 {
		bonus = profit * 0.1
	} else if profit > 10 && profit < 20 {
		bonus = 10*0.1 + (profit-10)*0.075
	} else if profit > 20 && profit < 40 {
		bonus = 10*0.1 + 10*0.075 + (profit-20)*0.05
	} else if profit > 40 && profit < 60 {
		bonus = 10*0.1 + 10*0.075 + 20*0.05 + (profit-40)*0.03
	} else if profit > 60 && profit < 100 {
		bonus = 10*0.1 + 10*0.075 + 20*0.05 + 20*0.03 + (profit-60)*0.015
	} else {
		bonus = 10*0.1 + 10*0.075 + 20*0.05 + 20*0.03 + 40*0.015 + (profit-100)*0.01
	}

	fmt.Println("bonus: ", bonus)
}
