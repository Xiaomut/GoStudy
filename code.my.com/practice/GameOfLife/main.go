package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Rules:
1. 当一个村花细胞邻近的存活细胞少于2个时，该细胞死亡
2. 当一个存活细胞临近有2个或3个存活细胞时，该细胞延续到下一代
3. 当一个村花细胞邻近的存活细胞多于3个时，该细胞死亡
4. 当一个村花细胞邻近的存活细胞正好有3个时，该细胞存活

Steps:
1. 判断细胞是否存活
2. 统计邻近存活细胞数量
3. 判断细胞在下一世代存活或死亡
*/

const (
	width         = 10
	height        = 10
	alive_percent = 25
)

type Universe [][]bool

func NewUniverse() (u Universe) {
	u = make(Universe, height)
	for i := range u {
		u[i] = make([]bool, width)
	}
	return u
}

func (u Universe) Show() {
	// 观察世界
	for i := range u {
		fmt.Println(u[i])
	}
}

func (u Universe) Seed() {
	// 激活细胞，可以随机激活世界中约25%的细胞
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			randnum := rand.Intn(100)
			if randnum < alive_percent {
				u[i][j] = true
			} else {
				u[i][j] = false
			}
		}
	}
}

func (u Universe) Alive(x, y int) bool {
	// 判断存活还是死亡,我们需要判断边界情况
	if x < 0 {
		x += height
	} else if x > height-1 {
		x %= height
	}
	if y < 0 {
		y += width
	} else if y > width-1 {
		y %= width
	}
	return u[x][y]
}

func (u Universe) Neighbors(x, y int) int {
	// 统计邻近细胞的存活情况
	alive_num := 0 // 定义一个数值统计
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if (i == x) && (j == y) {
				continue
			}
			if u.Alive(i, j) {
				alive_num += 1
			}
		}
	}
	return alive_num
}

func (u Universe) Next(x, y int) bool {
	// 游戏逻辑，表示给定细胞在下一世代存活或死亡
	/*
		Rules:
		1. 当一个村花细胞邻近的存活细胞少于2个时，该细胞死亡
		2. 当一个存活细胞临近有2个或3个存活细胞时，该细胞延续到下一代
		3. 当一个村花细胞邻近的存活细胞多于3个时，该细胞死亡
		4. 当一个村花细胞邻近的存活细胞正好有3个时，该细胞存活
	*/
	alive_num := u.Neighbors(x, y)
	if u.Alive(x, y) {
		if alive_num < 2 || alive_num > 3 {
			u[x][y] = false
		} else if alive_num == 2 || alive_num == 3 {
			u[x][y] = true
		}
	} else {
		if alive_num == 3 {
			u[x][y] = true
		}
	}

	return false
}

func Step(a, b Universe) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			b[i][j] = a.Next(i, j)
		}
	}
}

func main() {
	u := NewUniverse()
	u.Seed()
	u.Show()

	for {
		u_copy := NewUniverse()
		Step(u, u_copy)
		time.Sleep(time.Duration(5) * time.Second)
		u = u_copy
		u.Show()
		u_copy.Show()
		fmt.Println("*****************************************")
	}
}
