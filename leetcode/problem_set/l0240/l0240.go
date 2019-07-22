package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	// start from the top right corner
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	x := len(matrix[0]) - 1
	y := 0
	for x >= 0 && y < len(matrix) {
		if matrix[y][x] == target {
			return true
		} else if matrix[y][x] > target {
			x--
		} else {
			y++
		}
	}
	return false
}

func main() {
	mt := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Println(searchMatrix(mt, 5))
	fmt.Println(searchMatrix(mt, 20))

	var empty [][]int
	fmt.Println(searchMatrix(empty, 0))

	mt1 := [][]int{
		{1, 1},
	}
	fmt.Println(searchMatrix(mt1, 0))

}
