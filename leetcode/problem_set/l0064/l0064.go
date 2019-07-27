package main

import "fmt"

func minPathSum(grid [][]int) int {
	rLen := len(grid)
	cLen := len(grid[0])
	for r := 0; r < rLen; r++ {
		for c := 0; c < cLen; c++ {
			if r == 0 && c != 0 {
				grid[r][c] += grid[r][c-1]
			}
			if r != 0 && c == 0 {
				grid[r][c] += grid[r-1][c]
			}
			if r != 0 && c != 0 {
				grid[r][c] += min(grid[r-1][c], grid[r][c-1])
			}
		}
	}
	return grid[rLen-1][cLen-1]
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	fmt.Println(minPathSum(grid))
}
