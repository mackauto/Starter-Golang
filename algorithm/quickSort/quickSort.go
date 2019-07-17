package main

import (
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

func partition(nums []int, left, right int) int {
	middle := (left + right) / 2
	pivot := nums[middle]
	for left <= right {
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		// double check loop condition
		if left <= right {
			util.Swap(nums, left, right)
			left++
			right--
		}
	}
	return left
}

func quickSort(nums []int, left, right int) {
	index := partition(nums, left, right)
	if left < index-1 {
		quickSort(nums, left, index-1)
	}
	if index < right {
		quickSort(nums, index, right)
	}
}

func main() {
	nums := util.GetExampleNums()
	util.PrintNums(nums)
	quickSort(nums, 0, len(nums)-1)
	util.PrintNums(nums)
}
