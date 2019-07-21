package main

import "fmt"

func square(num int) int {
	return num * num
}

func sortedSquares(A []int) []int {
	B := make([]int, len(A))
	i := 0
	j := len(A) - 1
	k := j
	for i <= j {
		i2 := square(A[i])
		j2 := square(A[j])
		if i2 > j2 {
			B[k] = i2
			i++
		} else {
			B[k] = j2
			j--
		}
		k--
	}
	return B
}

func main() {
	nums := []int{-7, -3, 2, 3, 11}
	fmt.Println(sortedSquares(nums))
}
