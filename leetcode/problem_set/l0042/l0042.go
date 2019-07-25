package main

import (
	"fmt"
)

func trapBruteForceMethod(height []int) int {
	ans := 0
	for i := 0; i < len(height); i++ {
		maxLeft := 0
		maxRight := 0
		for j := i; j >= 0; j-- {
			maxLeft = max(height[j], maxLeft)
		}
		for k := i; k < len(height); k++ {
			maxRight = max(height[k], maxRight)
		}
		ans += min(maxRight, maxLeft) - height[i]
	}
	return ans
}

func trapDynamicProgrammingMethod(height []int) int {
	size := len(height)
	if size <= 2 {
		return 0
	}
	var ans int
	leftMax := make([]int, size)
	rightMax := make([]int, size)
	leftMax[0] = height[0]
	for i := 1; i < size; i++ {
		leftMax[i] = max(height[i], leftMax[i-1])
	}
	rightMax[size-1] = height[size-1]
	for i := size - 2; i >= 0; i-- {
		rightMax[i] = max(height[i], rightMax[i+1])
	}
	for i := 1; i < size-1; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ans
}

func trapStackMethod(height []int) int {
	var ans, current int
	var stack []int
	for current < len(height) {
		for len(stack) != 0 && height[current] > height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			distance := current - stack[len(stack)-1] - 1
			boundedHeight := min(height[current], height[stack[len(stack)-1]]) - height[top]
			ans += distance * boundedHeight
		}
		stack = append(stack, current)
		current++
	}
	return ans
}

func trapTwoPointerMethod(height []int) int {
	left := 0
	right := len(height) - 1
	var ans, leftMax, rightMax int
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trapBruteForceMethod(height))
	fmt.Println(trapDynamicProgrammingMethod(height))
	fmt.Println(trapStackMethod(height))
	fmt.Println(trapTwoPointerMethod(height))
}
