package main

import "fmt"

func findNumberIn2DArray(matrix [][]int, target int) bool {
	row, col := len(matrix)-1, 0

	for row >= 0 && col < len(matrix[0]) {
		if target < matrix[row][col] {
			row -= 1
		} else if target > matrix[row][col] {
			col += 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}}
	res := findNumberIn2DArray(m, 3)
	fmt.Println(res)
}
