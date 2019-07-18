package main

import (
	"github.com/JPMike/Starter-Golang/leetcode/util"
)

// implement base on heap first index = 0
func heapify(nums []int, last int) {
	for i := (last - 1) / 2; i >= 0; i-- {
		maxPos := i
		for {
			if left := i*2 + 1; left <= last && nums[maxPos] < nums[left] {
				maxPos = left
			}
			if right := i*2 + 2; right <= last && nums[maxPos] < nums[right] {
				maxPos = right
			}
			if maxPos == i {
				break
			}
			util.Swap(nums, maxPos, i)
			i = maxPos
		}
	}
}

func heapSort(nums []int) {
	for last := len(nums) - 1; last > 0; last-- {
		heapify(nums, last)
		util.Swap(nums, 0, last)
	}
}

func main() {
	nums := util.GetExampleNums()
	util.PrintNums(nums)
	heapSort(nums)
	util.PrintNums(nums)
}
