package main

import (
	"fmt"
	"math"
)

func middleMaxSum(nums []int, i, m, j int) int {
	leftMax := math.MinInt32
	leftSum := 0
	for k := m; k >= i; k-- {
		leftSum += nums[k]
		if leftSum > leftMax {
			leftMax = leftSum
		}
	}
	rightMax := math.MinInt32
	rightSum := 0
	for k := m + 1; k <= j; k++ {
		rightSum += nums[k]
		if rightSum > rightMax {
			rightMax = rightSum
		}
	}
	return leftSum + rightSum
}

func subArrayMaxSumRecursiveMethod(nums []int, i, j int) int {
	// base condition
	if i == j {
		return nums[i]
	}
	m := (i + j) / 2
	leftMax := subArrayMaxSumRecursiveMethod(nums, i, m)
	rightMax := subArrayMaxSumRecursiveMethod(nums, m+1, j)
	middleMax := middleMaxSum(nums, i, m, j)
	if leftMax >= middleMax && leftMax >= rightMax {
		return leftMax
	}
	if rightMax >= middleMax && rightMax >= leftMax {
		return rightMax
	}
	return middleMax
}

func subArrayMaxSum(nums []int) int {
	temp := 0
	max := math.MinInt32
	for i := 0; i < len(nums); i++ {
		temp += nums[i]
		if temp > max {
			max = temp
		}
		if temp < 0 {
			temp = 0
		}
	}
	return max
}

func main() {
	nums := []int{2, 4, -7, 5, 2, -1, 2, -4, 3}
	fmt.Println(subArrayMaxSum(nums)) // 8 {5, 2, -1, 2}
	// todo recursive method did not do right
	fmt.Println(subArrayMaxSumRecursiveMethod(nums, 0, len(nums)-1))
}
