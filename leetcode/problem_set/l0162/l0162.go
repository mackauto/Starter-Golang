package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return len(nums) - 1
}

func findPeakElementIterationBinarySearchMethod(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid] > nums[mid+1] {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}

func main() {
	numsL := [][]int{
		{1, 2, 3, 1},
		{1, 2, 1, 3, 5, 6, 4},
		{4, 5, 6, 7, 0, 1, 2},
	}
	for _, nums := range numsL {
		fmt.Println(findPeakElement(nums))
		fmt.Println(findPeakElementIterationBinarySearchMethod(nums))
	}
}
