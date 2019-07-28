package main

import (
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

func partition(nums []int, left, right int) int {
	middle := (left + right) / 2
	pivot := nums[middle]
	for left <= right {
		// find swap position
		for nums[left] < pivot {
			left++
		}
		for nums[right] > pivot {
			right--
		}
		// double check loop condition before swap
		if left <= right {
			util.Swap(nums, left, right)
			left++
			right--
		}
	}
	return left
}

func quickSort(nums []int, left, right int) {
	if left >= right {
		return
	}
	// main sort process in partition, which is the unit of work
	index := partition(nums, left, right)
	// recursion in left and right
	quickSort(nums, left, index-1)
	quickSort(nums, index, right)
}

func main() {
	nums := util.GetExampleNums()
	util.PrintNums(nums)
	quickSort(nums, 0, len(nums)-1)
	util.PrintNums(nums)
}
